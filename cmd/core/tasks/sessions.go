package tasks

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/cufee/aftermath/cmd/core"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/logic"
	"github.com/rs/zerolog/log"
)

func init() {
	defaultHandlers[models.TaskTypeRecordSnapshots] = TaskHandler{
		Process: func(ctx context.Context, client core.Client, task *models.Task) error {
			if task.Data == nil {
				return errors.New("no data provided")
			}

			triesLeft, err := strconv.Atoi(task.Data["triesLeft"])
			if err != nil {
				return errors.Wrap(err, "failed to parse tries left")
			}
			task.Data["triesLeft"] = strconv.Itoa(triesLeft - 1)

			realm, ok := task.Data["realm"]
			if !ok {
				task.Data["triesLeft"] = "0" // do not retry
				return errors.New("invalid realm")
			}

			if len(task.Targets) > 100 {
				task.Data["triesLeft"] = "0" // do not retry
				return errors.New("invalid targets length")
			}
			if len(task.Targets) < 1 {
				task.Data["triesLeft"] = "0" // do not retry
				return errors.New("invalid targets length")
			}

			forceUpdate := task.Data["force"] == "true"

			log.Debug().Str("taskId", task.ID).Any("targets", task.Targets).Msg("started working on a session refresh task")

			accountErrors, err := logic.RecordAccountSnapshots(ctx, client.Wargaming(), client.Database(), realm, forceUpdate, task.Targets...)
			if err != nil {
				return err
			}

			// Retry failed accounts
			newTargets := make([]string, 0, len(accountErrors))
			for id, err := range accountErrors {
				if err == nil {
					continue
				}
				newTargets = append(newTargets, id)
				task.LogAttempt(models.TaskLog{
					Timestamp: time.Now(),
					Error:     err.Error(),
					Comment:   id,
				})
			}
			task.Targets = newTargets

			return nil
		},
	}
}

func CreateRecordSnapshotsTasks(client core.Client, realm string) error {
	realm = strings.ToUpper(realm)
	task := models.Task{
		Type:           models.TaskTypeRecordSnapshots,
		ReferenceID:    "realm_" + realm,
		ScheduledAfter: time.Now(),
		Data: map[string]string{
			"realm":     realm,
			"triesLeft": "3",
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	accounts, err := client.Database().GetRealmAccountIDs(ctx, realm)
	if err != nil {
		return err
	}
	if len(accounts) < 1 {
		return errors.New("no accounts on realm " + realm)
	}
	task.Targets = append(task.Targets, accounts...)

	// This update requires (2 + n) requests per n players
	tasks := splitTaskByTargets(task, 50)
	err = client.Database().CreateTasks(ctx, tasks...)
	if err != nil {
		return err
	}

	log.Debug().Int("count", len(tasks)).Msg("scheduled realm account snapshots tasks")
	return nil
}