package constants

import (
	"os"
	"strconv"
)

var (
	SchedulerEnabled           = os.Getenv("SCHEDULER_ENABLED") == "true"
	SchedulerConcurrentWorkers int
)

func init() {
	SchedulerConcurrentWorkers, _ = strconv.Atoi(mustGetEnv("SCHEDULER_CONCURRENT_WORKERS", func() bool { return !SchedulerEnabled }))
}
