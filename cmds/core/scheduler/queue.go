package scheduler

import (
	"context"
	"sync"
	"time"

	"github.com/cufee/aftermath/cmds/core"
	"github.com/cufee/aftermath/cmds/core/scheduler/tasks"
	"github.com/cufee/aftermath/internal/database"
	"github.com/rs/zerolog/log"
)

type Queue struct {
	limiter          chan struct{}
	concurrencyLimit int
	lastTaskRun      time.Time

	handlers map[database.TaskType]tasks.TaskHandler
	core     core.Client
}

func (q *Queue) ConcurrencyLimit() int {
	return q.concurrencyLimit
}

func (q *Queue) ActiveWorkers() int {
	return len(q.limiter)
}

func (q *Queue) LastTaskRun() time.Time {
	return q.lastTaskRun
}

func NewQueue(client core.Client, handlers map[database.TaskType]tasks.TaskHandler, concurrencyLimit int) *Queue {
	return &Queue{
		core:             client,
		handlers:         handlers,
		concurrencyLimit: concurrencyLimit,
		limiter:          make(chan struct{}, concurrencyLimit),
	}
}

func (q *Queue) Process(callback func(error), tasks ...database.Task) {
	var err error
	if callback != nil {
		defer callback(err)
	}
	if len(tasks) == 0 {
		log.Debug().Msg("no tasks to process")
		return
	}

	log.Debug().Msgf("processing %d tasks", len(tasks))

	var wg sync.WaitGroup
	q.lastTaskRun = time.Now()
	processedTasks := make(chan database.Task, len(tasks))
	for _, task := range tasks {
		wg.Add(1)
		go func(t database.Task) {
			q.limiter <- struct{}{}
			defer func() {
				processedTasks <- t
				wg.Done()
				<-q.limiter
				log.Debug().Msgf("finished processing task %s", t.ID)
			}()
			log.Debug().Msgf("processing task %s", t.ID)

			handler, ok := q.handlers[t.Type]
			if !ok {
				t.Status = database.TaskStatusFailed
				t.LogAttempt(database.TaskLog{
					Targets:   t.Targets,
					Timestamp: time.Now(),
					Error:     "missing task type handler",
				})
				return
			}

			attempt := database.TaskLog{
				Targets:   t.Targets,
				Timestamp: time.Now(),
			}

			message, err := handler.Process(nil, t)
			attempt.Comment = message
			if err != nil {
				attempt.Error = err.Error()
				t.Status = database.TaskStatusFailed
			} else {
				t.Status = database.TaskStatusComplete
			}
			t.LogAttempt(attempt)
		}(task)
	}

	wg.Wait()
	close(processedTasks)

	rescheduledCount := 0
	processedSlice := make([]database.Task, 0, len(processedTasks))
	for task := range processedTasks {
		handler, ok := q.handlers[task.Type]
		if !ok {
			continue
		}

		if task.Status == database.TaskStatusFailed && handler.ShouldRetry(&task) {
			rescheduledCount++
			task.Status = database.TaskStatusScheduled
		}
		processedSlice = append(processedSlice, task)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err = q.core.Database().UpdateTasks(ctx, processedSlice...)
	if err != nil {
		return
	}

	log.Debug().Msgf("processed %d tasks, %d rescheduled", len(processedSlice), rescheduledCount)
}
