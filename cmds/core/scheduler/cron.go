package scheduler

import (
	"time"

	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"
)

func (queue *Queue) StartCronJobsAsync() {
	defer log.Info().Msg("started cron scheduler")

	c := gocron.NewScheduler(time.UTC)
	// Tasks
	c.Cron("* * * * *").Do(RunTasksWorker(queue))
	// some tasks might be stuck due to a panic or restart, restart them
	c.Cron("0 * * * *").Do(RestartTasksWorker(queue))

	// Glossary - Do it around the same time WG releases game updates
	c.Cron("0 10 * * *").Do(UpdateGlossaryWorker(queue.core))
	c.Cron("0 12 * * *").Do(UpdateGlossaryWorker(queue.core))
	// c.AddFunc("40 9 * * 0", updateAchievementsWorker)

	// Averages - Update averages once daily
	c.Cron("0 0 * * *").Do(UpdateAveragesWorker(queue.core))

	// Sessions
	c.Cron("0 9 * * *").Do(CreateSessionTasksWorker(queue.core, "NA"))  // NA
	c.Cron("0 1 * * *").Do(CreateSessionTasksWorker(queue.core, "EU"))  // EU
	c.Cron("0 18 * * *").Do(CreateSessionTasksWorker(queue.core, "AS")) // Asia

	// Refresh WN8
	// "45 9 * * *" 	// NA
	// "45 1 * * *" 	// EU
	// "45 18 * * *" 	// Asia

	// Configurations
	c.Cron("0 0 */7 * *").Do(RotateBackgroundPresetsWorker(queue.core))

	// Start the Cron job scheduler
	c.StartAsync()
}
