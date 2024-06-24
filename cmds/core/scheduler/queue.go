package scheduler

import (
	"context"
	"runtime/debug"
	"sync"
	"time"

	"github.com/cufee/aftermath/cmds/core"
	"github.com/cufee/aftermath/cmds/core/scheduler/tasks"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/rs/zerolog/log"
)

type Queue struct {
	limiter          chan struct{}
	concurrencyLimit int
	lastTaskRun      time.Time

	handlers map[models.TaskType]tasks.TaskHandler
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

func NewQueue(client core.Client, handlers map[models.TaskType]tasks.TaskHandler, concurrencyLimit int) *Queue {
	return &Queue{
		core:             client,
		handlers:         handlers,
		concurrencyLimit: concurrencyLimit,
		limiter:          make(chan struct{}, concurrencyLimit),
	}
}

func (q *Queue) Process(callback func(error), tasks ...models.Task) {
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
	processedTasks := make(chan models.Task, len(tasks))
	for _, task := range tasks {
		wg.Add(1)
		go func(t models.Task) {
			defer func() {
				if r := recover(); r != nil {
					log.Error().Str("stack", string(debug.Stack())).Msg("panic in task handler")
					t.Status = models.TaskStatusFailed
					t.LogAttempt(models.TaskLog{
						Targets:   t.Targets,
						Timestamp: time.Now(),
						Error:     "task handler caused a panic",
					})
				} else {
					log.Debug().Msgf("finished processing task %s", t.ID)
				}
				processedTasks <- t
				<-q.limiter
				wg.Done()
			}()

			q.limiter <- struct{}{}
			log.Debug().Msgf("processing task %s", t.ID)

			handler, ok := q.handlers[t.Type]
			if !ok {
				t.Status = models.TaskStatusFailed
				t.LogAttempt(models.TaskLog{
					Targets:   t.Targets,
					Timestamp: time.Now(),
					Error:     "missing task type handler",
				})
				return
			}
			t.LastRun = time.Now()

			attempt := models.TaskLog{
				Targets:   t.Targets,
				Timestamp: time.Now(),
			}

			message, err := handler.Process(q.core, t)
			attempt.Comment = message
			if err != nil {
				attempt.Error = err.Error()
				t.Status = models.TaskStatusFailed
			} else {
				t.Status = models.TaskStatusComplete
			}
			t.LogAttempt(attempt)
		}(task)
	}

	wg.Wait()
	close(processedTasks)

	rescheduledCount := 0
	processedSlice := make([]models.Task, 0, len(processedTasks))
	for task := range processedTasks {
		handler, ok := q.handlers[task.Type]
		if !ok {
			continue
		}

		if task.Status == models.TaskStatusFailed && handler.ShouldRetry(&task) {
			rescheduledCount++
			task.Status = models.TaskStatusScheduled
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
