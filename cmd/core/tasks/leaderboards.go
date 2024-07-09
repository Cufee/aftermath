package tasks

import (
	"context"
	"slices"
	"strings"
	"time"

	"github.com/cufee/aftermath/cmd/core"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/log"
	"github.com/cufee/aftermath/internal/logic"
	"github.com/pkg/errors"
)

func init() {
	defaultHandlers[models.TaskTypeAchievementsLeaderboardUpdate] = TaskHandler{
		Process: func(ctx context.Context, client core.Client, task *models.Task) error {
			realm, ok := task.Data["realm"]
			if !ok {
				return errors.New("invalid realm")
			}
			scoreType, ok := task.Data["type"]
			if !ok || !slices.Contains(models.ScoreType("").Values(), scoreType) {
				return errors.New("invalid score type")
			}

			if len(task.Targets) > 100 {
				return errors.New("invalid targets length")
			}
			if len(task.Targets) < 1 {
				return errors.New("invalid targets length")
			}

			accountErrors, err := logic.RecordCurrentAchievementsLeaderboards(ctx, client.Wargaming(), client.Database(), models.ScoreType(scoreType), realm, false, task.Targets...)
			if err != nil {
				return err // implicitly retry all targets
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

func CreateUpdateLeaderboardsTasks(client core.Client, realm string, scoreType models.ScoreType) error {
	realm = strings.ToUpper(realm)
	task := models.Task{
		Type:           models.TaskTypeAchievementsLeaderboardUpdate,
		ReferenceID:    "realm_" + realm,
		ScheduledAfter: time.Now(),
		Data: map[string]string{
			"type":  string(scoreType),
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
		return errors.New("no accounts on realm " + realm)
	}
	task.Targets = append(task.Targets, accounts...)

	// This update requires (2 + n) requests per n players
	// 1 - get all accounts last battle time
	// 1 - get all account achievements
	// n - get vehicle achievements for each account
	tasks := splitTaskByTargets(task, 50)
	err = client.Database().CreateTasks(ctx, tasks...)
	if err != nil {
		return err
	}

	log.Debug().Int("count", len(tasks)).Msg("scheduled realm leaderboard update tasks")
	return nil
}
