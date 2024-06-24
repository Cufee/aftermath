package queue

import (
	"context"
	"fmt"
	"runtime/debug"
	"sync"
	"time"

	"github.com/cufee/aftermath/cmds/core"
	"github.com/cufee/aftermath/cmds/core/scheduler"
	"github.com/cufee/aftermath/cmds/core/tasks"
	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type queue struct {
	queueLock *sync.Mutex
	queued    chan models.Task
	handlers  map[models.TaskType]tasks.TaskHandler

	workerLimit   int
	workerLimiter chan struct{}
	workerTimeout time.Duration

	activeTasks   map[string]struct{}
	activeTasksMx *sync.Mutex
	newCoreClient func() (core.Client, error)
}

func New(workerLimit int, newCoreClient func() (core.Client, error)) *queue {
	return &queue{
		newCoreClient: newCoreClient,
		handlers:      make(map[models.TaskType]tasks.TaskHandler),

		workerLimit:   workerLimit,
		workerTimeout: time.Second * 15,
		workerLimiter: make(chan struct{}, workerLimit),

		activeTasksMx: &sync.Mutex{},
		activeTasks:   make(map[string]struct{}),

		queueLock: &sync.Mutex{},
		queued:    make(chan models.Task, workerLimit*2),
	}
}

func (q *queue) SetHandlers(handlers map[models.TaskType]tasks.TaskHandler) error {
	for t, h := range handlers {
		if _, ok := q.handlers[t]; ok {
			return fmt.Errorf("handler for %s is already registered", t)
		}
		q.handlers[t] = h
	}
	return nil
}

func (q *queue) Start(ctx context.Context) (func(), error) {
	qctx, qCancel := context.WithCancel(ctx)

	coreClint, err := q.newCoreClient()
	if err != nil {
		qCancel()
		return nil, err
	}

	s := scheduler.New()
	// on cron, pull and enqueue available tasks
	s.Add("* * * * *", func() {
		pctx, cancel := context.WithTimeout(qctx, time.Second*5)
		defer cancel()
		err := q.pullAndEnqueueTasks(pctx, coreClint.Database())
		if err != nil {
			log.Err(err).Msg("failed to pull tasks")
		}
	})
	stopScheduler, err := s.Start(qctx)
	if err != nil {
		qCancel()
		return nil, err
	}

	go q.startWorkers(qctx, func(_ string) {
		if len(q.queued) < q.workerLimit {
			err := q.pullAndEnqueueTasks(qctx, coreClint.Database())
			if err != nil {
				log.Err(err).Msg("failed to pull tasks")
			}
		}
	})

	return func() {
		qCancel()
		stopScheduler()

		q.activeTasksMx.Lock()
		defer q.activeTasksMx.Unlock()

		var abandonedIDs []string
		for id := range q.activeTasks {
			abandonedIDs = append(abandonedIDs, id)
		}

		cctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		err = coreClint.Database().AbandonTasks(cctx, abandonedIDs...)
		if err != nil {
			log.Err(err).Msg("failed to abandon tasks")
		}
	}, nil
}

func (q *queue) Enqueue(task models.Task) {
	q.queued <- task
}

func (q *queue) pullAndEnqueueTasks(ctx context.Context, db database.Client) error {
	if ok := q.queueLock.TryLock(); !ok {
		return errors.New("queue is locked")
	}
	defer q.queueLock.Unlock()

	log.Debug().Msg("pulling available tasks into queue")

	tasks, err := db.GetAndStartTasks(ctx, q.workerLimit)
	if err != nil && !database.IsNotFound(err) {
		return err
	}
	log.Debug().Msgf("pulled %d tasks into queue", len(tasks))

	go func(tasks []models.Task) {
		for _, t := range tasks {
			q.Enqueue(t)
		}
	}(tasks)

	return nil
}

func (q *queue) startWorkers(ctx context.Context, onComplete func(id string)) {
	for {
		select {
		case task := <-q.queued:
			go func() {
				q.workerLimiter <- struct{}{}
				q.activeTasksMx.Lock()
				q.activeTasks[task.ID] = struct{}{}
				q.activeTasksMx.Unlock()
				defer func() {
					<-q.workerLimiter
					q.activeTasksMx.Lock()
					delete(q.activeTasks, task.ID)
					q.activeTasksMx.Unlock()
					if onComplete != nil {
						onComplete(task.ID)
					}
				}()

				log.Debug().Str("taskId", task.ID).Msg("worker started processing a task")
				defer log.Debug().Str("taskId", task.ID).Msg("worker finished processing a task")

				defer func() {
					if r := recover(); r != nil {
						event := log.Error().Str("stack", string(debug.Stack())).Str("taskId", task.ID)
						defer event.Msg("panic in queue worker")

						coreClient, err := q.newCoreClient()
						if err != nil {
							event.AnErr("core", err).Str("additional", "failed to create a core client")
							return
						}
						task.Status = models.TaskStatusFailed
						task.LogAttempt(models.TaskLog{
							Timestamp: time.Now(),
							Comment:   "task caused a panic in worker handler",
						})
						err = coreClient.Database().UpdateTasks(context.Background(), task)
						if err != nil {
							event.AnErr("updateTasks", err).Str("additional", "failed to update a task")
						}
					}
				}()

				wctx, cancel := context.WithTimeout(ctx, q.workerTimeout)
				defer cancel()

				coreClient, err := q.newCoreClient()
				if err != nil {
					log.Err(err).Msg("failed to create a new core client for a task worker")
					return
				}

				handler, ok := q.handlers[task.Type]
				if !ok {
					task.Status = models.TaskStatusFailed
					task.LogAttempt(models.TaskLog{
						Error:     "task missing a handler",
						Comment:   "task missing a handler",
						Timestamp: time.Now(),
					})
					err := coreClient.Database().UpdateTasks(wctx, task)
					if err != nil {
						log.Err(err).Msg("failed to update a task")
					}
					return
				}

				msg, err := handler.Process(coreClient, task)
				task.Status = models.TaskStatusComplete
				task.LogAttempt(models.TaskLog{
					Error:     err.Error(),
					Timestamp: time.Now(),
					Comment:   msg,
				})

				if err != nil {
					if handler.ShouldRetry(&task) {
						task.Status = models.TaskStatusScheduled
					} else {
						task.Status = models.TaskStatusFailed
					}
				}

				err = coreClient.Database().UpdateTasks(wctx, task)
				if err != nil {
					log.Err(err).Msg("failed to update a task")
				}
			}()

		case <-ctx.Done():
			return
		}
	}
}
