package discord

import "github.com/bwmarrin/discordgo"

type InteractionType int

const (
	IntegrationTypeGuild InteractionType = 0
	IntegrationTypeUser  InteractionType = 1
)

var InteractionTypeAll = []InteractionType{IntegrationTypeGuild, IntegrationTypeUser}

type InteractionContextType int

const (
	InteractionContextGuild          InteractionContextType = 0
	InteractionContextDirectMessage  InteractionContextType = 1
	InteractionContextPrivateChannel InteractionContextType = 2
)

var InteractionContextAll = []InteractionContextType{InteractionContextGuild, InteractionContextDirectMessage, InteractionContextPrivateChannel}

type ApplicationCommand struct {
	ID                string                           `json:"id,omitempty"`
	ApplicationID     string                           `json:"application_id,omitempty"`
	GuildID           string                           `json:"guild_id,omitempty"`
	Version           string                           `json:"version,omitempty"`
	Type              discordgo.ApplicationCommandType `json:"type,omitempty"`
	Name              string                           `json:"name"`
	NameLocalizations *map[discordgo.Locale]string     `json:"name_localizations,omitempty"`
	// NOTE: DefaultPermission will be soon deprecated. Use DefaultMemberPermissions and DMPermission instead.
	DefaultPermission        *bool                    `json:"default_permission,omitempty"`
	DefaultMemberPermissions *int64                   `json:"default_member_permissions,string,omitempty"`
	DMPermission             *bool                    `json:"dm_permission,omitempty"`
	NSFW                     *bool                    `json:"nsfw,omitempty"`
	IntegrationTypes         []InteractionType        `json:"integration_types,omitempty"`
	Contexts                 []InteractionContextType `json:"contexts,omitempty"`

	// NOTE: Chat commands only. Otherwise it mustn't be set.

	Description              string                                `json:"description,omitempty"`
	DescriptionLocalizations *map[discordgo.Locale]string          `json:"description_localizations,omitempty"`
	Options                  []*discordgo.ApplicationCommandOption `json:"options"`
}
