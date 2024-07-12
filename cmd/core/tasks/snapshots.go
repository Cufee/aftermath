package tasks

import (
	"context"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/cufee/aftermath/cmd/core"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/log"
	"github.com/cufee/aftermath/internal/logic"
)

func init() {
	defaultHandlers[models.TaskTypeRecordSnapshots] = TaskHandler{
		Process: func(ctx context.Context, client core.Client, task *models.Task) error {
			forceUpdate := task.Data["force"] == "true"
			realm, ok := task.Data["realm"]
			if !ok {
				return errors.New("invalid realm")
			}
			if len(task.Targets) > 100 {
				return errors.New("invalid targets length")
			}
			if len(task.Targets) < 1 {
				return errors.New("invalid targets length")
			}

			log.Debug().Str("taskId", task.ID).Any("targets", task.Targets).Msg("started working on a session refresh task")
			accountErrors, err := logic.RecordAccountSnapshots(ctx, client.Wargaming(), client.Database(), realm, forceUpdate, task.Targets...)
			if err != nil {
				return err
			}

			// retry failed accounts
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

			if len(newTargets) > 0 {
				task.Targets = newTargets
				return errors.New("some accounts failed")
			}
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
			"realm": realm,
		},
		TriesLeft: 3,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	accounts, err := client.Database().GetRealmAccountIDs(ctx, realm)
	if err != nil {
		return err
	}
	if len(accounts) < 1 {
		log.Warn().Str("realm", realm).Msg("no accounts on realm")
		return nil
	}
	task.Targets = append(task.Targets, accounts...)

	// This update requires (3 + 2n) requests per n players
	// 1 - get all account stats
	// 1 - get all clan profiles
	// 1 - get all account achievements
	// n - get vehicle stats for each account
	// n - get vehicle achievements for each account
	tasks := splitTaskByTargets(task, 50)
	err = client.Database().CreateTasks(ctx, tasks...)
	if err != nil {
		return err
	}

	log.Debug().Int("count", len(tasks)).Msg("scheduled realm account snapshots tasks")
	return nil
}
