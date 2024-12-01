package blitzstars

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cufee/aftermath/internal/json"

	"github.com/cufee/aftermath/internal/stats/frame"
)

// Response from https://www.blitzstars.com/ average stats endpoint
type VehicleAverages struct {
	TankID  int `json:"tank_id"`
	Players int `json:"number_of_players"`
	All     struct {
		AvgBattles              float32 `json:"battles,omitempty"`
		AvgDroppedCapturePoints float32 `json:"dropped_capture_points,omitempty"`
	} `json:",omitempty"`
	Special struct {
		Winrate         float32 `json:"winrate,omitempty"`
		DamageRatio     float32 `json:"damageRatio,omitempty"`
		Kdr             float32 `json:"kdr,omitempty"`
		DamagePerBattle float32 `json:"damagePerBattle,omitempty"`
		KillsPerBattle  float32 `json:"killsPerBattle,omitempty"`
		HitsPerBattle   float32 `json:"hitsPerBattle,omitempty"`
		SpotsPerBattle  float32 `json:"spotsPerBattle,omitempty"`
		Wpm             float32 `json:"wpm,omitempty"`
		Dpm             float32 `json:"dpm,omitempty"`
		Kpm             float32 `json:"kpm,omitempty"`
		HitRate         float32 `json:"hitRate,omitempty"`
		SurvivalRate    float32 `json:"survivalRate,omitempty"`
	} `json:"special,omitempty"`
}

func (c client) CurrentTankAverages(ctx context.Context) (map[string]frame.StatsFrame, error) {
	req, err := http.NewRequest("GET", c.apiURL+"/tankaverages.json", nil)
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

	var averages []VehicleAverages
	err = json.NewDecoder(res.Body).Decode(&averages)
	if err != nil {
		return nil, err
	}

	averagesMap := make(map[string]frame.StatsFrame, len(averages))
	for _, average := range averages {
		battles := average.All.AvgBattles * float32(average.Players)

		averagesMap[fmt.Sprint(average.TankID)] = frame.StatsFrame{
			Battles:     frame.ValueInt(battles),
			BattlesWon:  frame.ValueInt(average.Special.Winrate * battles / 100),
			DamageDealt: frame.ValueInt(average.Special.DamagePerBattle * battles),

			ShotsHit:   frame.ValueInt(average.Special.HitsPerBattle * battles),
			ShotsFired: frame.ValueInt((average.Special.HitsPerBattle * battles) / (average.Special.HitRate / 100)),

			Frags:                frame.ValueInt(average.Special.KillsPerBattle * battles),
			EnemiesSpotted:       frame.ValueInt(average.Special.SpotsPerBattle * battles),
			DroppedCapturePoints: frame.ValueInt(average.All.AvgDroppedCapturePoints * battles),
		}

	}

	return averagesMap, nil
}
