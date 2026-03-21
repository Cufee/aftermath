package averages

import (
	"context"
	"sync"

	"github.com/cufee/aftermath/internal/log"
	"github.com/cufee/aftermath/internal/stats/frame"
	"github.com/pkg/errors"
)

var ErrServiceUnavailable = errors.New("tank averages unavailable")

type Source interface {
	CurrentTankAverages(ctx context.Context) (map[string]frame.StatsFrame, error)
}

type Client struct {
	primary   Source
	secondary Source
}

// NewClient creates a merged averages client that fetches from both sources
// concurrently and merges results, preferring the primary source for any
// overlapping tank IDs.
func NewClient(primary, secondary Source) *Client {
	return &Client{primary: primary, secondary: secondary}
}

func (c *Client) CurrentTankAverages(ctx context.Context) (map[string]frame.StatsFrame, error) {
	var wg sync.WaitGroup

	var primaryResult map[string]frame.StatsFrame
	var primaryErr error
	var secondaryResult map[string]frame.StatsFrame
	var secondaryErr error

	wg.Go(func() {
		primaryResult, primaryErr = c.primary.CurrentTankAverages(ctx)
	})
	wg.Go(func() {
		secondaryResult, secondaryErr = c.secondary.CurrentTankAverages(ctx)
	})
	wg.Wait()

	if primaryErr != nil {
		log.Err(primaryErr).Msg("primary averages source failed")
	}
	if secondaryErr != nil {
		log.Err(secondaryErr).Msg("secondary averages source failed")
	}

	if primaryErr != nil && secondaryErr != nil {
		return nil, errors.Wrap(ErrServiceUnavailable, "all averages sources failed")
	}

	// start with secondary, then overwrite with primary so primary wins on overlap
	merged := make(map[string]frame.StatsFrame, max(len(primaryResult), len(secondaryResult)))
	for id, avg := range secondaryResult {
		merged[id] = avg
	}
	for id, avg := range primaryResult {
		merged[id] = avg
	}

	return merged, nil
}
