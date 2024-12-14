package wargaming

import (
	"context"
	"fmt"
	"image"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/cufee/aftermath/internal/constants"
	"github.com/cufee/aftermath/internal/external/images"
	"github.com/cufee/aftermath/internal/json"
	"github.com/cufee/aftermath/internal/log"
	"github.com/cufee/aftermath/internal/render/assets"
	"github.com/cufee/am-wg-proxy-next/v2/types"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

type AuctionMemoryCache struct {
	mx     *sync.Mutex
	data   map[types.Realm]RealmCache
	images map[string]image.Image

	onUpdateCallback func(current, previous RealmCache)
}

type RealmCache struct {
	Realm         types.Realm
	LastUpdate    time.Time
	NextPriceDrop time.Time
	Vehicles      []AuctionVehicle
}

func (cache *AuctionMemoryCache) Current(realm types.Realm) (RealmCache, bool) {
	data, ok := cache.data[realm]
	if !ok {
		return RealmCache{}, false
	}
	return data, true
}

func (cache *AuctionMemoryCache) Update(ctx context.Context, realm types.Realm) error {
	cache.mx.Lock()
	defer cache.mx.Unlock()

	vehicles, err := cache.currentAuction(ctx, realm)
	if err != nil {
		return errors.Wrap(err, "update failed on "+realm.String())
	}

	last := cache.data[realm]
	current := RealmCache{
		Realm:      realm,
		Vehicles:   vehicles,
		LastUpdate: time.Now(),
	}
	cache.data[realm] = current
	go cache.onUpdateCallback(current, last)

	return nil
}

func NewAuctionCache(updateCallback func(current, previous RealmCache)) (*AuctionMemoryCache, error) {
	fallbackImage, ok := assets.GetLoadedImage("secret-vehicle")
	if !ok {
		return nil, errors.New("failed to load fallback vehicle image")
	}

	if updateCallback == nil {
		updateCallback = func(current, previous RealmCache) {}
	}
	cache := &AuctionMemoryCache{
		mx:               &sync.Mutex{},
		data:             make(map[types.Realm]RealmCache),
		images:           make(map[string]image.Image, 25),
		onUpdateCallback: updateCallback,
	}

	cache.images["fallback"] = fallbackImage
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

func (cache AuctionMemoryCache) currentAuction(ctx context.Context, realm types.Realm) ([]AuctionVehicle, error) {
	domain, ok := realm.DomainBlitz()
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

	resultCh := make(chan AuctionVehicle, len(data.Results))
	var group errgroup.Group
	for _, item := range data.Results {
		if !item.Available || item.Type != "vehicle" {
			continue
		}
		group.Go(func() error {
			vehicle := AuctionVehicle{
				ID:            fmt.Sprint(item.Entity.ID),
				Image:         cache.images[item.Entity.PreviewImageURL],
				Total:         item.InitialCount,
				Available:     item.CurrentCount,
				IsPremium:     item.Entity.IsPremium,
				IsCollectible: item.Entity.IsCollectible,
				NextPriceDrop: time.Unix(int64(item.NextPriceTimestamp), 0),
				Price:         AuctionPrice{item.Price.toPrice(), item.NextPrice.toPrice()},
			}

			defer func() {
				if vehicle.Image == nil {
					vehicle.Image = cache.images["fallback"]
				}
				resultCh <- vehicle
			}()

			if vehicle.Image == nil && item.Entity.PreviewImageURL != "" {
				link, err := url.Parse(item.Entity.PreviewImageURL)
				if err != nil {
					log.Warn().Err(err).Str("url", link.String()).Msg("failed to parse auction vehicle preview url")
					return nil
				}

				ictx, cancel := context.WithTimeout(ctx, time.Second*5)
				defer cancel()

				loaded, err := images.SafeLoadFromURL(ictx, link, constants.ImageUploadMaxSize)
				if err != nil {
					log.Warn().Err(err).Str("url", link.String()).Msg("failed to load auction vehicle preview url")
					return nil
				}

				cache.images[item.Entity.PreviewImageURL] = loaded
				vehicle.Image = loaded
			}

			return nil
		})
	}
	err = group.Wait()
	close(resultCh)
	if err != nil {
		return nil, err
	}

	var result []AuctionVehicle
	for v := range resultCh {
		result = append(result, v)
	}
	return result, nil
}
