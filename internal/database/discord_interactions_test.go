package database

import (
	"context"
	"fmt"
	"testing"

	"github.com/cufee/aftermath/internal/database/gen/table"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/matryer/is"
	"golang.org/x/text/language"
)

func TestDiscordInteraction(t *testing.T) {
	client := MustTestClient(t)
	is := is.New(t)

	defer client.db.Exec(fmt.Sprintf("DELETE FROM %s;", table.DiscordInteraction.TableName()))

	user, err := client.GetOrCreateUserByID(context.Background(), "user-TestDiscordInteraction")
	is.NoErr(err)

	defer client.db.Exec(fmt.Sprintf("DELETE FROM %s WHERE id = '%s';", table.User.TableName(), user.ID))

	t.Run("create and get discord interaction", func(t *testing.T) {
		data := models.DiscordInteraction{
			Snowflake: "flake-1",
			Result:    "result-1",
			UserID:    user.ID,
			ChannelID: "channel-1",
			GuildID:   "guild-1",
			MessageID: "message-1",
			EventID:   "event-1",
			Locale:    language.English,
			Type:      models.InteractionTypeCommand,
		}
		inserted, err := client.CreateDiscordInteraction(context.Background(), data)
		is.NoErr(err)
		is.True(inserted.ID != "")
		is.Equal(inserted.UserID, data.UserID)
		is.Equal(inserted.Result, data.Result)
		is.Equal(inserted.Snowflake, data.Snowflake)

		found, err := client.GetDiscordInteraction(context.Background(), inserted.ID)
		is.NoErr(err)
		is.Equal(found.ID, inserted.ID)
	})

	t.Run("create and find discord interaction", func(t *testing.T) {
		data1 := models.DiscordInteraction{
			Snowflake: "flake-2",
			Result:    "result-2",
			UserID:    user.ID,
			ChannelID: "channel-2",
			GuildID:   "guild-2",
			MessageID: "message-2",
			EventID:   "event-2",
			Locale:    language.English,
			Type:      models.InteractionTypeCommand,
		}
		data2 := models.DiscordInteraction{
			Snowflake: "flake-3",
			Result:    "result-3",
			UserID:    user.ID,
			ChannelID: "channel-2",
			GuildID:   "guild-3",
			MessageID: "message-3",
			EventID:   "event-3",
			Locale:    language.English,
			Type:      models.InteractionTypeCommand,
		}
		{
			_, err := client.CreateDiscordInteraction(context.Background(), data1)
			is.NoErr(err)
			_, err = client.CreateDiscordInteraction(context.Background(), data2)
			is.NoErr(err)
		}
		{
			found, err := client.FindDiscordInteractions(context.Background(), WithEventID(data1.EventID))
			is.NoErr(err)
			is.True(len(found) == 1)
			is.Equal(found[0].EventID, data1.EventID)
		}
		{
			found, err := client.FindDiscordInteractions(context.Background(), WithMessageID(data2.MessageID))
			is.NoErr(err)
			is.True(len(found) == 1)
			is.Equal(found[0].MessageID, data2.MessageID)
		}
		{
			found, err := client.FindDiscordInteractions(context.Background(), WithChannelID(data1.ChannelID))
			is.NoErr(err)
			is.True(len(found) == 2)
			is.True(found[0].Snowflake != found[1].Snowflake)
		}
	})
}
