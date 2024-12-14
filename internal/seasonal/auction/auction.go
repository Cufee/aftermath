package auction

import (
	"context"
	"errors"
	"image"
	"slices"
	"time"

	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/external/wargaming"
	"github.com/cufee/aftermath/internal/log"
	"github.com/cufee/am-wg-proxy-next/v2/types"
	"golang.org/x/sync/errgroup"
)

type cacheType struct {
	auction *wargaming.AuctionMemoryCache
	db      database.GlossaryClient
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
	cache = &cacheType{auction, db}
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
	data := make([]Vehicle, len(current.Vehicles))
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
