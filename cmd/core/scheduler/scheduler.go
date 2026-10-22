package scheduler

import (
	"context"
	"fmt"
	"time"

	"github.com/cufee/aftermath/cmd/core"
	"github.com/cufee/aftermath/internal/log"
	"github.com/go-co-op/gocron/v2"
)

type scheduler struct {
	cron gocron.Scheduler
	jobs []gocron.Job
}

func New() (*scheduler, error) {
	// v1 waited for running jobs indefinitely. Keep that behavior in practice,
	// including the averages worker's fifteen-minute jitter, instead of v2's
	// default ten-second timeout. Reserve room for gocron's two-second margin.
	cron, err := gocron.NewScheduler(gocron.WithLocation(time.UTC),
		gocron.WithStopTimeout(time.Duration(1<<63-1)-2*time.Second))
	if err != nil {
		return nil, err
	}
	return &scheduler{cron: cron}, nil
}

func (s *scheduler) Add(cron string, fn func()) (gocron.Job, error) {
	j, err := s.cron.NewJob(gocron.CronJob(cron, false), gocron.NewTask(fn))
	if err != nil {
		return nil, err
	}
	s.jobs = append(s.jobs, j)
	return j, nil
}

func (s *scheduler) Start(ctx context.Context) (func(), error) {
	s.cron.Start()
	return func() {
		if err := s.cron.Shutdown(); err != nil {
			log.Err(err).Msg("failed to shut down scheduler")
		}
	}, nil
}

func RegisterDefaultTasks(s *scheduler, coreClient core.Client) error {
	jobs := []struct {
		cron string
		fn   func()
	}{
		{"0 * * * *", RestartTasksWorker(coreClient)},
		{"0 15 * * *", CleanupPubSubWorker(coreClient)},
		{"0 5 * * *", CreateCleanupTaskWorker(coreClient)},
		// Refresh around the time WG releases game updates.
		{"10 12 * * *", UpdateGlossaryWorker(coreClient)},
		{"10 10 * * *", UpdateGlossaryWorker(coreClient)},
		{"0 0 * * *", UpdateAveragesWorker(coreClient)},
		{"0 9 * * *", CreateSnapshotTasksWorker(coreClient, "NA")},
		{"0 1 * * *", CreateSnapshotTasksWorker(coreClient, "EU")},
		{"0 18 * * *", CreateSnapshotTasksWorker(coreClient, "AS")},
	}
	for _, job := range jobs {
		if _, err := s.Add(job.cron, job.fn); err != nil {
			_ = s.cron.Shutdown()
			return fmt.Errorf("register scheduled task %q: %w", job.cron, err)
		}
	}
	return nil
}
