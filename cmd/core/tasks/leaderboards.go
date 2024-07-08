package tasks

import (
	"github.com/cufee/aftermath/cmd/core"
)

func init() {
	// defaultHandlers[models.TaskTypeAchievementsLeaderboardUpdate] = TaskHandler{
	// 	Process: func(ctx context.Context, client core.Client, task *models.Task) error {
	// 		realm, ok := task.Data["realm"]
	// 		if !ok {
	// 			return errors.New("invalid realm")
	// 		}
	// 		if len(task.Targets) > 100 {
	// 			return errors.New("invalid targets length")
	// 		}
	// 		if len(task.Targets) < 1 {
	// 			return errors.New("invalid targets length")
	// 		}

	// 		ctx, cancel := context.WithCancel(ctx)
	// 		defer cancel()

	// 		lastBattles, err := client.Wargaming().BatchAccountByID(ctx, realm, task.Targets, "account_id", "last_battle_time")
	// 		if err != nil {
	// 			return errors.Wrap(err, "failed to get account data from wargaming")
	// 		}

	// 		client.Database().GetAchievementsSnapshots(ctx, )

	// 		// calculate the current score for each player

	// 		// save scores to database

	// 		return nil
	// 	},
	// }
}

func CreateUpdateLeaderboardsTasks(client core.Client, realm string) error {
	//
	return nil
}
