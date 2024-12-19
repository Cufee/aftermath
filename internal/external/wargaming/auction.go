package wargaming

import (
	"context"
	"fmt"
	"image"
	"net/http"
	"sync"
	"time"

	"github.com/cufee/aftermath/internal/json"
	"github.com/cufee/am-wg-proxy-next/v2/types"
	"github.com/pkg/errors"
)

type AuctionMemoryCache struct {
	mx   *sync.Mutex
	data *Cache
}

type Cache struct {
	LastUpdate time.Time
	Vehicles   []AuctionVehicle
}

func (cache *AuctionMemoryCache) Current() (*Cache, bool) {
	if cache.data == nil {
		return nil, false
	}
	return cache.data, true
}

func (cache *AuctionMemoryCache) Update(ctx context.Context) error {
	cache.mx.Lock()
	defer cache.mx.Unlock()

	vehicles, err := cache.currentAuction(ctx)
	if err != nil {
		return errors.Wrap(err, "update failed")
	}

	cache.data = &Cache{
		Vehicles:   vehicles,
		LastUpdate: time.Now(),
	}

	return nil
}

func NewAuctionCache() (*AuctionMemoryCache, error) {
	cache := &AuctionMemoryCache{
		mx: &sync.Mutex{},
	}
	return cache, nil
}

var auctionHttpClient = http.DefaultClient

type AuctionVehicle struct {
	ID    string      `json:"id"`
	Image image.Image `json:"image"`

	Available int `json:"claimed"`
	Total     int `json:"total"`

	IsPremium     bool `json:"premium"`
	IsCollectible bool `json:"collectible"`

	Price         AuctionPrice `json:"price"`
	NextPriceDrop time.Time    `json:"nextPriceDrop"`
}

type AuctionPrice struct {
	Current  Price `json:"current"`
	Upcoming Price `json:"upcoming"`
}

type priceType string

const (
	PriceGold    = priceType("gold")
	PriceSilver  = priceType("silver")
	PriceUnknown = priceType("unknown")
)

type Price struct {
	Value int       `json:"value"`
	Type  priceType `json:"type"`
}

type currency struct {
	Name     string `json:"name"`
	Count    int    `json:"count"`
	Title    string `json:"title"`
	ImageURL string `json:"image_url"`
	Type     string `json:"type"`
}

type price struct {
	Currency currency `json:"currency"`
	Value    int      `json:"value"`
}

func (p price) toPrice() Price {
	price := Price{Value: p.Value}
	switch p.Currency.Name {
	default:
		price.Type = PriceUnknown
	case "gold":
		price.Type = PriceGold
	case "silver":
		price.Type = PriceSilver
	}
	return price
}

type auctionResponse struct {
	Count   int    `json:"count"`
	HasNext bool   `json:"has_next"`
	Results []item `json:"results"`
}
type entity struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	TypeSlug        string `json:"type_slug"`
	Level           int    `json:"level"`
	RomanLevel      string `json:"roman_level"`
	UserString      string `json:"user_string"`
	ImageURL        string `json:"image_url"`
	PreviewImageURL string `json:"preview_image_url"`
	IsPremium       bool   `json:"is_premium"`
	IsCollectible   bool   `json:"is_collectible"`
}
type item struct {
	ID                 int    `json:"id"`
	Type               string `json:"type"`
	Countable          bool   `json:"countable"`
	Entity             entity `json:"entity"`
	InitialCount       int    `json:"initial_count"`
	CurrentCount       int    `json:"current_count"`
	Saleable           bool   `json:"saleable"`
	AvailableFrom      string `json:"available_from"`
	AvailableBefore    string `json:"available_before"`
	Price              price  `json:"price"`
	NextPrice          price  `json:"next_price"`
	Available          bool   `json:"available"`
	Display            bool   `json:"display"`
	NextPriceDatetime  string `json:"next_price_datetime"`
	NextPriceTimestamp int    `json:"next_price_timestamp"`
}

func (cache AuctionMemoryCache) currentAuction(ctx context.Context) ([]AuctionVehicle, error) {

	domain, ok := types.RealmNorthAmerica.DomainBlitz()
	if !ok {
		return nil, ErrRealmNotSupported
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://%s/en/api/events/items/auction/?page_size=100&type[]=vehicle&saleable=true", domain), nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	res, err := auctionHttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var data auctionResponse
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	var vehicles = map[int]AuctionVehicle{}
	for _, item := range data.Results {
		if !item.Available || item.Type != "vehicle" || item.CurrentCount == 0 {
			continue
		}
		vehicles[item.Entity.ID] = AuctionVehicle{
			ID:            fmt.Sprint(item.Entity.ID),
			Total:         item.InitialCount,
			Available:     item.CurrentCount,
			IsPremium:     item.Entity.IsPremium,
			IsCollectible: item.Entity.IsCollectible,
			NextPriceDrop: time.Unix(int64(item.NextPriceTimestamp), 0),
			Price:         AuctionPrice{item.Price.toPrice(), item.NextPrice.toPrice()},
		}
	}

	var result []AuctionVehicle
	for _, v := range vehicles {
		result = append(result, v)
	}
	return result, nil
}
