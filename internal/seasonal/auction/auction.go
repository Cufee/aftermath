package auction

import (
	"bytes"
	"context"
	"errors"
	"image"
	"image/png"
	"slices"
	"sync"
	"time"

	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/external/wargaming"
	"github.com/cufee/aftermath/internal/log"
	"github.com/cufee/am-wg-proxy-next/v2/types"
	"golang.org/x/sync/errgroup"
	"golang.org/x/text/language"
)

type cacheType struct {
	auction *wargaming.AuctionMemoryCache
	db      database.GlossaryClient

	mx       *sync.Mutex
	rendered map[types.Realm][][]byte
}

type Vehicle struct {
	models.Vehicle
	Image image.Image

	Available int
	Total     int

	Price         wargaming.AuctionPrice
	NextPriceDrop time.Time
}

var cache *cacheType

func InitAuctionCache(db database.GlossaryClient) error {
	onUpdate := func(current, previous wargaming.RealmCache) {}
	auction, err := wargaming.NewAuctionCache(onUpdate)
	if err != nil {
		return err
	}
	cache = &cacheType{auction, db, &sync.Mutex{}, make(map[types.Realm][][]byte)}

	return nil
}

func UpdateAuctionCache(realms ...types.Realm) {
	if cache == nil {
		log.Error().Msg("auction cache update called before the cache init")
		return
	}

	var group errgroup.Group
	for _, realm := range realms {
		group.Go(func() error {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
			defer cancel()

			err := cache.auction.Update(ctx, realm)
			if err != nil {
				log.Err(err).Str("realm", realm.String()).Msg("failed to update auction cache")
			}

			err = cache.updateAuctionImagesCache(ctx, realm)
			if err != nil {
				log.Err(err).Str("realm", realm.String()).Msg("failed to update auction image cache")
			}
			return nil
		})
	}
	group.Wait()
}

var (
	ErrRealmNotCached = errors.New("realm not cached")
)

type AuctionVehicles struct {
	Vehicles      []Vehicle
	LastUpdate    time.Time
	NextUpdate    time.Time
	NextPriceDrop time.Time
}

func CurrentAuction(ctx context.Context, realm types.Realm) (AuctionVehicles, error) {
	current, ok := cache.auction.Current(realm)
	if !ok {
		return AuctionVehicles{}, ErrRealmNotCached
	}

	var ids []string
	for _, v := range current.Vehicles {
		ids = append(ids, v.ID)
	}

	glossary, err := cache.db.GetVehicles(ctx, ids)
	if err != nil {
		return AuctionVehicles{}, err
	}

	var nextPriceDrop time.Time
	data := make([]Vehicle, 0, len(current.Vehicles))
	for _, vehicle := range current.Vehicles {
		v := Vehicle{
			Image:         vehicle.Image,
			Price:         vehicle.Price,
			Total:         vehicle.Total,
			Available:     vehicle.Available,
			Vehicle:       glossary[vehicle.ID],
			NextPriceDrop: vehicle.NextPriceDrop,
		}
		// make sure these values are correct
		v.Premium = vehicle.IsPremium
		v.Collectible = vehicle.IsCollectible
		data = append(data, v)

		if v.NextPriceDrop.Before(nextPriceDrop) {
			nextPriceDrop = v.NextPriceDrop
		}
	}

	slices.SortFunc(data, func(a, b Vehicle) int {
		return a.Available - b.Available
	})

	return AuctionVehicles{
		Vehicles:      data,
		LastUpdate:    current.LastUpdate,
		NextPriceDrop: nextPriceDrop,
	}, nil
}

func CurrentAuctionImages(ctx context.Context, realm types.Realm) ([][]byte, AuctionVehicles, error) {
	data, err := CurrentAuction(ctx, realm)
	if err != nil {
		return nil, AuctionVehicles{}, err
	}
	if images, ok := cache.rendered[realm]; ok && len(images) > 0 {
		return images, data, nil
	}
	return nil, AuctionVehicles{}, ErrRealmNotCached
}

func (c *cacheType) updateAuctionImagesCache(ctx context.Context, realm types.Realm) error {
	data, err := CurrentAuction(ctx, realm)
	if err != nil {
		return err
	}

	cards, err := AuctionCards(data, language.English)
	if err != nil {
		return err
	}

	var imagesMx sync.Mutex
	images := make([][]byte, len(cards))

	var group errgroup.Group
	for i, img := range cards {
		group.Go(func() error {
			var buf bytes.Buffer
			err = png.Encode(&buf, img)
			if err != nil {
				return err
			}
			imagesMx.Lock()
			defer imagesMx.Unlock()

			images[i] = buf.Bytes()
			return nil
		})
	}
	err = group.Wait()
	if err != nil {
		return err
	}

	cache.mx.Lock()
	defer cache.mx.Unlock()

	cache.rendered[realm] = images
	return nil
}
