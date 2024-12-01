package blitzstars

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cufee/aftermath/internal/json"

	"github.com/cufee/am-wg-proxy-next/v2/types"
)

type TankHistoryEntry struct {
	TankID          int              `json:"tank_id"`
	LastBattleTime  int              `json:"last_battle_time"`
	BattlesLifeTime int              `json:"battle_life_time"`
	MarkOfMastery   int              `json:"mark_of_mastery"`
	Stats           types.StatsFrame `json:"all"`
}

func (c client) AccountTankHistories(ctx context.Context, accountId string) (map[int][]TankHistoryEntry, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/tankhistories/for/%s", c.apiURL, accountId), nil)
	if err != nil {
		return nil, err
	}

	res, err := c.http.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status code: %d", res.StatusCode)
	}

	var histories []TankHistoryEntry
	err = json.NewDecoder(res.Body).Decode(&histories)
	if err != nil {
		return nil, err
	}

	var historiesMap = make(map[int][]TankHistoryEntry, len(histories))
	for _, entry := range histories {
		historiesMap[entry.TankID] = append(historiesMap[entry.TankID], entry)
	}

	return historiesMap, nil
}
