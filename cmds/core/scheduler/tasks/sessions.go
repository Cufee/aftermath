package tasks

import (
	"context"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/cufee/aftermath/cmds/core"
	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/logic"
	"github.com/rs/zerolog/log"
)

func init() {
	defaultHandlers[database.TaskTypeRecordSessions] = TaskHandler{
		Process: func(client core.Client, task database.Task) (string, error) {
			if task.Data == nil {
				return "no data provided", errors.New("no data provided")
			}
			realm, ok := task.Data["realm"].(string)
			if !ok {
				task.Data["triesLeft"] = int(0) // do not retry
				return "invalid realm", errors.New("invalid realm")
			}
			if len(task.Targets) > 100 {
				task.Data["triesLeft"] = int(0) // do not retry
				return "cannot process 100+ accounts at a time", errors.New("invalid targets length")
			}
			if len(task.Targets) < 1 {
				task.Data["triesLeft"] = int(0) // do not retry
				return "targed ids cannot be left blank", errors.New("invalid targets length")
			}
			forceUpdate, _ := task.Data["force"].(bool)

			log.Debug().Str("taskId", task.ID).Any("targets", task.Targets).Msg("started working on a session refresh task")

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
			defer cancel()

			accountErrors, err := logic.RecordAccountSnapshots(ctx, client.Wargaming(), client.Database(), realm, forceUpdate, task.Targets...)
			if err != nil {
				return "failed to record sessions", err
			}

			if len(accountErrors) == 0 {
				return "finished session update on all accounts", nil
			}

			// Retry failed accounts
			task.Targets = make([]string, len(accountErrors))
			for id := range accountErrors {
				task.Targets = append(task.Targets, id)
			}
			return "retrying failed accounts", errors.New("some accounts failed")
		},
		ShouldRetry: func(task *database.Task) bool {
			triesLeft, ok := task.Data["triesLeft"].(int)
			if !ok {
				return false
			}
			if triesLeft <= 0 {
				return false
			}

			triesLeft -= 1
			task.Data["triesLeft"] = triesLeft
			task.ScheduledAfter = time.Now().Add(5 * time.Minute) // Backoff for 5 minutes to avoid spamming
			return true
		},
	}
}

func CreateSessionUpdateTasks(client core.Client, realm string) error {
	realm = strings.ToUpper(realm)
	task := database.Task{
		Type:           database.TaskTypeRecordSessions,
		ReferenceID:    "realm_" + realm,
		ScheduledAfter: time.Now(),
		Data: map[string]any{
			"realm":     realm,
			"triesLeft": int(3),
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	accounts, err := client.Database().GetRealmAccounts(ctx, realm)
	if err != nil {
		return err
	}
	if len(accounts) < 1 {
		return nil
	}
	for _, account := range accounts {
		task.Targets = append(task.Targets, account.ID)
	}

	// This update requires (2 + n) requests per n players
	return client.Database().CreateTasks(ctx, splitTaskByTargets(task, 90)...)
}
