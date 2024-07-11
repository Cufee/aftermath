package scheduler

import (
	"context"
	"time"

	"github.com/cufee/aftermath/cmd/core"
	"github.com/cufee/aftermath/internal/database/models"
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
	s.Add("10 12 * * *", UpdateGlossaryWorker(coreClient))
	s.Add("10 10 * * *", UpdateGlossaryWorker(coreClient))
	// c.AddFunc("40 9 * * 0", updateAchievementsWorker)

	// Averages - Update averages once daily
	s.Add("0 0 * * *", UpdateAveragesWorker(coreClient))

	// Snapshots
	s.Add("0 9 * * *", CreateSnapshotTasksWorker(coreClient, "NA"))  // NA
	s.Add("0 1 * * *", CreateSnapshotTasksWorker(coreClient, "EU"))  // EU
	s.Add("0 18 * * *", CreateSnapshotTasksWorker(coreClient, "AS")) // Asia

	// Achievement leaderboards. ideally, this should not delay snapshots
	// hourly
	s.Add("25 * * * *", CreateLeaderboardTasksWorker(coreClient, "NA", models.LeaderboardScoreHourly)) // NA
	s.Add("30 * * * *", CreateLeaderboardTasksWorker(coreClient, "EU", models.LeaderboardScoreHourly)) // EU
	s.Add("35 * * * *", CreateLeaderboardTasksWorker(coreClient, "AS", models.LeaderboardScoreHourly)) // Asia
	// daily
	s.Add("0 8 * * *", CreateLeaderboardTasksWorker(coreClient, "NA", models.LeaderboardScoreDaily))  // NA
	s.Add("0 0 * * *", CreateLeaderboardTasksWorker(coreClient, "EU", models.LeaderboardScoreDaily))  // EU
	s.Add("0 17 * * *", CreateLeaderboardTasksWorker(coreClient, "AS", models.LeaderboardScoreDaily)) // Asia

	// Configurations
	s.Add("0 0 */7 * *", RotateBackgroundPresetsWorker(coreClient))
}
