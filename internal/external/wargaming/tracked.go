package wargaming

import (
	"context"

	"github.com/cufee/am-wg-proxy-next/v2/types"
)

type requestObserver interface {
	Record(source, operation string, failed bool)
}

type trackedClient struct {
	Client
	observer requestObserver
}

func NewTrackedClient(client Client, observer requestObserver) Client {
	return &trackedClient{
		Client:   client,
		observer: observer,
	}
}

func (c *trackedClient) SearchAccounts(ctx context.Context, realm types.Realm, query string, opts ...types.Option) ([]types.Account, error) {
	data, err := c.Client.SearchAccounts(ctx, realm, query, opts...)
	c.observer.Record("wargaming", "search_accounts", err != nil)
	return data, err
}

func (c *trackedClient) AccountByID(ctx context.Context, realm types.Realm, id string, opts ...types.Option) (types.ExtendedAccount, error) {
	data, err := c.Client.AccountByID(ctx, realm, id, opts...)
	c.observer.Record("wargaming", "account_by_id", err != nil)
	return data, err
}

func (c *trackedClient) BatchAccountByID(ctx context.Context, realm types.Realm, ids []string, opts ...types.Option) (map[string]types.ExtendedAccount, error) {
	data, err := c.Client.BatchAccountByID(ctx, realm, ids, opts...)
	c.observer.Record("wargaming", "batch_account_by_id", err != nil)
	return data, err
}

func (c *trackedClient) AccountClan(ctx context.Context, realm types.Realm, id string, opts ...types.Option) (types.ClanMember, error) {
	data, err := c.Client.AccountClan(ctx, realm, id, opts...)
	c.observer.Record("wargaming", "account_clan", err != nil)
	return data, err
}

func (c *trackedClient) BatchAccountClan(ctx context.Context, realm types.Realm, ids []string, opts ...types.Option) (map[string]types.ClanMember, error) {
	data, err := c.Client.BatchAccountClan(ctx, realm, ids, opts...)
	c.observer.Record("wargaming", "batch_account_clan", err != nil)
	return data, err
}

func (c *trackedClient) AccountVehicles(ctx context.Context, realm types.Realm, id string, vehicles []string, opts ...types.Option) ([]types.VehicleStatsFrame, error) {
	data, err := c.Client.AccountVehicles(ctx, realm, id, vehicles, opts...)
	c.observer.Record("wargaming", "account_vehicles", err != nil)
	return data, err
}

func (c *trackedClient) AccountAchievements(ctx context.Context, realm types.Realm, id string, opts ...types.Option) (types.AchievementsFrame, error) {
	data, err := c.Client.AccountAchievements(ctx, realm, id, opts...)
	c.observer.Record("wargaming", "account_achievements", err != nil)
	return data, err
}

func (c *trackedClient) AccountVehicleAchievements(ctx context.Context, realm types.Realm, id string, opts ...types.Option) (map[string]types.AchievementsFrame, error) {
	data, err := c.Client.AccountVehicleAchievements(ctx, realm, id, opts...)
	c.observer.Record("wargaming", "account_vehicle_achievements", err != nil)
	return data, err
}

func (c *trackedClient) BatchAccountAchievements(ctx context.Context, realm types.Realm, ids []string, opts ...types.Option) (map[string]types.AchievementsFrame, error) {
	data, err := c.Client.BatchAccountAchievements(ctx, realm, ids, opts...)
	c.observer.Record("wargaming", "batch_account_achievements", err != nil)
	return data, err
}

func (c *trackedClient) SearchClans(ctx context.Context, realm types.Realm, query string, opts ...types.Option) ([]types.Clan, error) {
	data, err := c.Client.SearchClans(ctx, realm, query, opts...)
	c.observer.Record("wargaming", "search_clans", err != nil)
	return data, err
}

func (c *trackedClient) ClanByID(ctx context.Context, realm types.Realm, id string, opts ...types.Option) (types.ExtendedClan, error) {
	data, err := c.Client.ClanByID(ctx, realm, id, opts...)
	c.observer.Record("wargaming", "clan_by_id", err != nil)
	return data, err
}

func (c *trackedClient) BatchClanByID(ctx context.Context, realm types.Realm, ids []string, opts ...types.Option) (map[string]types.ExtendedClan, error) {
	data, err := c.Client.BatchClanByID(ctx, realm, ids, opts...)
	c.observer.Record("wargaming", "batch_clan_by_id", err != nil)
	return data, err
}

func (c *trackedClient) VehicleGlossary(ctx context.Context, realm types.Realm, vehicleID string, opts ...types.Option) (types.VehicleDetails, error) {
	data, err := c.Client.VehicleGlossary(ctx, realm, vehicleID, opts...)
	c.observer.Record("wargaming", "vehicle_glossary", err != nil)
	return data, err
}

func (c *trackedClient) CompleteVehicleGlossary(ctx context.Context, realm types.Realm, opts ...types.Option) (map[string]types.VehicleDetails, error) {
	data, err := c.Client.CompleteVehicleGlossary(ctx, realm, opts...)
	c.observer.Record("wargaming", "complete_vehicle_glossary", err != nil)
	return data, err
}
