package blitzkit

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cufee/aftermath/internal/json"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"

	"github.com/cufee/aftermath/internal/stats/frame"
)

const supportedManifestVersion = 1
const assetsURL = "https://raw.githubusercontent.com/blitzkit/assets/main"

type averagesManifest struct {
	Version int    `json:"version"`
	Latest  uint32 `json:"latest"`
}

func (c client) CurrentTankAverages(ctx context.Context) (map[string]frame.StatsFrame, error) {
	manifest, err := c.fetchManifest(ctx)
	if err != nil {
		return nil, err
	}

	averages, err := c.fetchAverages(ctx, manifest.Latest)
	if err != nil {
		return nil, err
	}

	averagesMap := make(map[string]frame.StatsFrame, len(averages.GetAverages()))
	for tankID, average := range averages.GetAverages() {
		if average == nil || average.Mu == nil {
			continue
		}

		mu := average.GetMu()
		if mu.GetBattles() <= 0 {
			continue
		}

		averagesMap[fmt.Sprint(tankID)] = frame.StatsFrame{
			Battles:              frame.ValueInt(mu.GetBattles()),
			BattlesWon:           frame.ValueInt(mu.GetWins()),
			DamageDealt:          frame.ValueInt(mu.GetDamageDealt()),
			ShotsHit:             frame.ValueInt(mu.GetHits()),
			ShotsFired:           frame.ValueInt(mu.GetShots()),
			Frags:                frame.ValueInt(mu.GetFrags()),
			EnemiesSpotted:       frame.ValueInt(mu.GetSpotted()),
			DroppedCapturePoints: frame.ValueInt(mu.GetDroppedCapturePoints()),
		}
	}

	return averagesMap, nil
}

func (c client) fetchManifest(ctx context.Context) (averagesManifest, error) {
	url := assetsURL + "/averages/manifest.json"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return averagesManifest{}, err
	}

	res, err := c.http.Do(req.WithContext(ctx))
	if err != nil {
		return averagesManifest{}, err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return averagesManifest{}, errors.Wrapf(ErrServiceUnavailable, "manifest request bad status code: %d", res.StatusCode)
	}

	var manifest averagesManifest
	if err := json.NewDecoder(res.Body).Decode(&manifest); err != nil {
		return averagesManifest{}, err
	}

	if manifest.Version != supportedManifestVersion {
		return averagesManifest{}, errors.Wrapf(ErrServiceUnavailable, "unsupported manifest version: %d", manifest.Version)
	}
	if manifest.Latest == 0 {
		return averagesManifest{}, errors.Wrap(ErrServiceUnavailable, "invalid manifest latest value")
	}

	return manifest, nil
}

func (c client) fetchAverages(ctx context.Context, latest uint32) (*AverageDefinitions, error) {
	url := fmt.Sprintf("%s/averages/%d.pb", assetsURL, latest)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.http.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, errors.Wrapf(ErrServiceUnavailable, "averages request bad status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var averages AverageDefinitions
	if err := proto.Unmarshal(body, &averages); err != nil {
		return nil, errors.Wrapf(ErrServiceUnavailable, "failed to decode averages payload: %v", err)
	}

	return &averages, nil
}
