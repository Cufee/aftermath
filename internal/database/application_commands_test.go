package database

import (
	"context"
	"testing"

	"github.com/cufee/aftermath/internal/database/models"
	"github.com/matryer/is"
)

func TestApplicationCommands(t *testing.T) {
	client := MustTestClient(t)
	is := is.New(t)

	t.Run("get command by hash", func(t *testing.T) {
		err := client.UpsertCommands(context.Background(), models.ApplicationCommand{
			ID:      "command-1",
			Hash:    "hash-1",
			Name:    "command-name-1",
			Version: "command-1-1",
		})
		is.NoErr(err)

		cmd, err := client.GetCommandsByHash(context.Background(), "hash-1")
		is.NoErr(err)
		is.True(len(cmd) == 1)
		is.True(cmd[0].ID == "command-1")
	})
	t.Run("get command by id", func(t *testing.T) {
		err := client.UpsertCommands(context.Background(), models.ApplicationCommand{
			ID:      "command-2",
			Hash:    "hash-2",
			Name:    "command-name-2",
			Version: "command-2-1",
		})
		is.NoErr(err)

		cmd, err := client.GetCommandsByID(context.Background(), "command-2")
		is.NoErr(err)
		is.True(len(cmd) == 1)
		is.True(cmd[0].ID == "command-2")
	})
	t.Run("update command hash", func(t *testing.T) {
		err := client.UpsertCommands(context.Background(), models.ApplicationCommand{
			ID:      "command-3",
			Hash:    "hash-3",
			Name:    "command-name-3",
			Version: "command-3-1",
		})
		is.NoErr(err)

		cmd, err := client.GetCommandsByHash(context.Background(), "hash-3")
		is.NoErr(err)
		is.True(len(cmd) == 1)
		is.True(cmd[0].ID == "command-3")

		err = client.UpsertCommands(context.Background(), models.ApplicationCommand{
			ID:      "command-3",
			Hash:    "hash-4",
			Name:    "command-name-3",
			Version: "command-3-2",
		})
		is.NoErr(err)

		cmd, err = client.GetCommandsByHash(context.Background(), "hash-4")
		is.NoErr(err)
		is.True(len(cmd) == 1)
		is.True(cmd[0].ID == "command-3")
		is.True(cmd[0].Version == "command-3-2")
	})
}
