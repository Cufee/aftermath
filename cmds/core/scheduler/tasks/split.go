package tasks

import "github.com/cufee/aftermath/internal/database/models"

func splitTaskByTargets(task models.Task, batchSize int) []models.Task {
	if len(task.Targets) <= batchSize {
		return []models.Task{task}
	}

	var tasks []models.Task
	subTasks := len(task.Targets) / batchSize

	for i := 0; i <= subTasks; i++ {
		subTask := task
		if len(task.Targets) > batchSize*(i+1) {
			subTask.Targets = (task.Targets[batchSize*i : batchSize*(i+1)])
		} else {
			subTask.Targets = (task.Targets[batchSize*i:])
		}
		tasks = append(tasks, subTask)
	}

	return tasks
}
