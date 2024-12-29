package constants

import (
	"os"
	"strconv"
)

var DevMode = os.Getenv("DEV_MODE") == "true"

var (
	SchedulerEnabled           = os.Getenv("SCHEDULER_ENABLED") == "true"
	SchedulerConcurrentWorkers int
)

func init() {
	SchedulerConcurrentWorkers, _ = strconv.Atoi(MustGetEnv("SCHEDULER_CONCURRENT_WORKERS", func() bool { return !SchedulerEnabled }))
}
