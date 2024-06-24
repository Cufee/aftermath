package private

import (
	"net/http"
	"slices"
	"strings"

	"github.com/cufee/aftermath/cmds/core"
	"github.com/cufee/aftermath/cmds/core/tasks"
)

var validRealms = []string{"na", "eu", "as"}

func SaveRealmSnapshots(client core.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		realm := r.PathValue("realm")
		if realm == "" {
			http.Error(w, "realm is required", http.StatusBadRequest)
			return
		}
		if !slices.Contains(validRealms, strings.ToLower(realm)) {
			http.Error(w, realm+" is not a valid realm", http.StatusBadRequest)
			return
		}

		err := tasks.CreateRecordSnapshotsTasks(client, realm)
		if err != nil {
			http.Error(w, "failed to create tasks: "+err.Error(), http.StatusBadRequest)
			return
		}

		w.Write([]byte("tasks scheduled"))
	}
}
