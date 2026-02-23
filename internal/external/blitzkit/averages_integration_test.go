package blitzkit

import (
	"context"
	"testing"
	"time"

	"github.com/matryer/is"
)

func TestCurrentTankAveragesIntegration(t *testing.T) {
	is := is.New(t)

	client, err := NewClient(time.Second * 10)
	is.NoErr(err)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	averages, err := client.CurrentTankAverages(ctx)
	is.NoErr(err)
	is.True(len(averages) > 0)

	var hasValidEntry bool
	for _, avg := range averages {
		if avg.Battles > 0 && avg.DamageDealt > 0 && avg.ShotsFired > 0 {
			hasValidEntry = true
			break
		}
	}

	is.True(hasValidEntry)
}
