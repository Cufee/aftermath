package scheduler

import (
	"time"

	"github.com/cufee/aftermath/cmds/core"
	"github.com/cufee/aftermath/cmds/core/scheduler/tasks"
	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"
)

func StartCronJobs(client core.Client, queue *tasks.Queue) {
	defer log.Info().Msg("started cron scheduler")

	c := gocron.NewScheduler(time.UTC)
	// Tasks
	c.Cron("* * * * *").Do(runTasksWorker(queue))
	c.Cron("0 * * * *").Do(restartTasksWorker(queue))

	// Glossary - Do it around the same time WG releases game updates
	c.Cron("0 10 * * *").Do(UpdateGlossaryWorker(client))
	c.Cron("0 12 * * *").Do(UpdateGlossaryWorker(client))
	// c.AddFunc("40 9 * * 0", updateAchievementsWorker)

	// Averages - Update averages shortly after session refreshes
	c.Cron("0 10 * * *").Do(UpdateAveragesWorker(client))
	c.Cron("0 2 * * *").Do(UpdateAveragesWorker(client))
	c.Cron("0 19 * * *").Do(UpdateAveragesWorker(client))

	// Sessions
	c.Cron("0 9 * * *").Do(createSessionTasksWorker(client, "NA"))  // NA
	c.Cron("0 1 * * *").Do(createSessionTasksWorker(client, "EU"))  // EU
	c.Cron("0 18 * * *").Do(createSessionTasksWorker(client, "AS")) // Asia

	// Refresh WN8
	// "45 9 * * *" 	// NA
	// "45 1 * * *" 	// EU
	// "45 18 * * *" 	// Asia

	// Configurations
	c.Cron("0 0 */7 * *").Do(rotateBackgroundPresetsWorker(client))

	// Start the Cron job scheduler
	c.StartAsync()
}
