package private

import (
	"net/http"

	"github.com/cufee/aftermath/cmds/core"
	"github.com/cufee/aftermath/cmds/core/scheduler"
)

func RestartStaleTasks(client core.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		scheduler.RestartTasksWorker(client)()
		w.Write([]byte("stale tasks restarted"))
	}
}
