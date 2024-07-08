package tasks

import (
	"github.com/cufee/aftermath/cmd/core"
)

// func init() {
// 	defaultHandlers[models.TaskTypeAchievementsLeaderboardUpdate] = TaskHandler{
// 		Process: func(ctx context.Context, client core.Client, task *models.Task) error {
// 			forceUpdate := task.Data["force"] == "true"
// 			realm, ok := task.Data["realm"]
// 			if !ok {
// 				return errors.New("invalid realm")
// 			}
// 			if len(task.Targets) > 100 {
// 				return errors.New("invalid targets length")
// 			}
// 			if len(task.Targets) < 1 {
// 				return errors.New("invalid targets length")
// 			}

// 			log.Debug().Str("taskId", task.ID).Any("targets", task.Targets).Msg("started working on a session refresh task")
// 			accountErrors, err := logic.RecordAccountSnapshots(ctx, client.Wargaming(), client.Database(), realm, forceUpdate, task.Targets...)
// 			if err != nil {
// 				return err
// 			}

// 			return nil
// 		},
// 	}
// }

func CreateUpdateLeaderboardsTasks(client core.Client, realm string) error {
	//
	return nil
}
