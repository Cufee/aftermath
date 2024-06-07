package tasks

import "github.com/cufee/aftermath/internal/database"

func splitTaskByTargets(task database.Task, batchSize int) []database.Task {
	if len(task.Targets) <= batchSize {
		return []database.Task{task}
	}

	var tasks []database.Task
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
