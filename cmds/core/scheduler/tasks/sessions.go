package tasks

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/cufee/aftermath/cmds/core"
	"github.com/cufee/aftermath/internal/database"
)

func init() {
	defaultHandlers[database.TaskTypeRecordSessions] = TaskHandler{
		Process: func(client core.Client, task database.Task) (string, error) {
			if task.Data == nil {
				return "no data provided", errors.New("no data provided")
			}
			realm, ok := task.Data["realm"].(string)
			if !ok {
				return "invalid realm", errors.New("invalid realm")
			}

			return "did nothing for a task on realm " + realm, nil

			// accountErrs, err := cache.RefreshSessionsAndAccounts(models.SessionTypeDaily, nil, realm, task.Targets...)
			// if err != nil {
			// 	return "failed to refresh sessions on all account", err
			// }

			// var failedAccounts []int
			// for accountId, err := range accountErrs {
			// 	if err != nil && accountId != 0 {
			// 		failedAccounts = append(failedAccounts, accountId)
			// 	}
			// }
			// if len(failedAccounts) == 0 {
			// 	return "finished session update on all accounts", nil
			// }

			// // Retry failed accounts
			// task.Targets = failedAccounts
			// return "retrying failed accounts", errors.New("some accounts failed")
		},
		ShouldRetry: func(task *database.Task) bool {
			triesLeft, ok := task.Data["triesLeft"].(int32)
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
		Type:        database.TaskTypeRecordSessions,
		ReferenceID: "realm_" + realm,
		Data: map[string]any{
			"realm":     realm,
			"triesLeft": int32(3),
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

	// This update requires (2 + n) requests per n players
	return client.Database().CreateTasks(ctx, splitTaskByTargets(task, 90)...)
}
