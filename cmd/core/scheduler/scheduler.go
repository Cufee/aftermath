package scheduler

import (
	"context"
	"time"

	"github.com/cufee/aftermath/cmd/core"
	"github.com/go-co-op/gocron"
)

type scheduler struct {
	cron *gocron.Scheduler
	jobs []*gocron.Job
}

func New() *scheduler {
	return &scheduler{cron: gocron.NewScheduler(time.UTC)}
}

func (s *scheduler) Add(cron string, fn func()) (*gocron.Job, error) {
	j, err := s.cron.Cron(cron).Do(fn)
	if err != nil {
		return nil, err
	}
	s.jobs = append(s.jobs, j)
	return j, nil
}

func (s *scheduler) Start(ctx context.Context) (func(), error) {
	go s.cron.StartBlocking()
	return s.cron.Stop, nil
}

func RegisterDefaultTasks(s *scheduler, coreClient core.Client) {
	s.Add("0 * * * *", RestartTasksWorker(coreClient))
	s.Add("0 5 * * *", CreateCleanupTaskWorker(coreClient)) // delete expired documents

	// Glossary - Do it around the same time WG releases game updates
	s.Add("0 10 * * *", UpdateGlossaryWorker(coreClient))
	s.Add("0 12 * * *", UpdateGlossaryWorker(coreClient))
	// c.AddFunc("40 9 * * 0", updateAchievementsWorker)

	// Averages - Update averages once daily
	s.Add("0 0 * * *", UpdateAveragesWorker(coreClient))

	// Snapshots
	s.Add("0 9 * * *", CreateSnapshotTasksWorker(coreClient, "NA"))  // NA
	s.Add("0 1 * * *", CreateSnapshotTasksWorker(coreClient, "EU"))  // EU
	s.Add("0 18 * * *", CreateSnapshotTasksWorker(coreClient, "AS")) // Asia

	// Achievement leaderboards
	// ideally, this should not delay snapshots
	s.Add("25 * * * *", CreateSnapshotTasksWorker(coreClient, "NA")) // NA
	s.Add("30 * * * *", CreateSnapshotTasksWorker(coreClient, "EU")) // EU
	s.Add("35 * * * *", CreateSnapshotTasksWorker(coreClient, "AS")) // Asia

	// Configurations
	s.Add("0 0 */7 * *", RotateBackgroundPresetsWorker(coreClient))
}
