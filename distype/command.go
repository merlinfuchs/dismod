package distype

type ApplicationCommand struct {
	ID                       Snowflake                  `json:"id"`
	Type                     *ApplicationCommandType    `json:"type,omitempty"`
	ApplicationID            Snowflake                  `json:"application_id"`
	GuildID                  *Snowflake                 `json:"guild_id,omitempty"`
	Name                     string                     `json:"name"`
	NameLocalizations        map[string]string          `json:"name_localizations,omitempty"`
	Description              string                     `json:"description"`
	DescriptionLocalizations map[string]string          `json:"description_localizations,omitempty"`
	Options                  []ApplicationCommandOption `json:"options,omitempty"`
	DefaultMemberPermissions Nullable[string]           `json:"default_member_permission,omitempty"`
	DMPermission             *bool                      `json:"dm_permission,omitempty"`
	NSFW                     *bool                      `json:"nsfw,omitempty"`
	Version                  Snowflake                  `json:"version"`
}

type ApplicationCommandType int

const (
	ApplicationCommandTypeChatInput ApplicationCommandType = 1
	ApplicationCommandTypeUser      ApplicationCommandType = 2
	ApplicationCommandTypeMessage   ApplicationCommandType = 3
)

type ApplicationCommandOption struct {
	Type                 ApplicationCommandOptionType     `json:"type"`
	Name                 string                           `json:"name"`
	NameLocations        map[string]string                `json:"name_locations,omitempty"`
	Description          string                           `json:"description"`
	DescriptionLocations map[string]string                `json:"description_locations,omitempty"`
	Required             *bool                            `json:"required,omitempty"`
	Choices              []ApplicationCommandOptionChoice `json:"choices,omitempty"`
	Options              []ApplicationCommandOption       `json:"options,omitempty"`
	ChannelTypes         []ChannelType                    `json:"channel_types,omitempty"`
	MinValue             *int                             `json:"min_value,omitempty"`
	MaxValue             *int                             `json:"max_value,omitempty"`
	MinLength            *int                             `json:"min_length,omitempty"`
	MaxLength            *int                             `json:"max_length,omitempty"`
	Autocomplete         *bool                            `json:"autocomplete,omitempty"`
}

type ApplicationCommandOptionType int

const (
	ApplicationCommandOptionTypeSubCommand      ApplicationCommandOptionType = 1
	ApplicationCommandOptionTypeSubCommandGroup ApplicationCommandOptionType = 2
	ApplicationCommandOptionTypeString          ApplicationCommandOptionType = 3
	ApplicationCommandOptionTypeInteger         ApplicationCommandOptionType = 4
	ApplicationCommandOptionTypeBoolean         ApplicationCommandOptionType = 5
	ApplicationCommandOptionTypeUser            ApplicationCommandOptionType = 6
	ApplicationCommandOptionTypeChannel         ApplicationCommandOptionType = 7
	ApplicationCommandOptionTypeRole            ApplicationCommandOptionType = 8
	ApplicationCommandOptionTypeMentionable     ApplicationCommandOptionType = 9
	ApplicationCommandOptionTypeNumber          ApplicationCommandOptionType = 10
	ApplicationCommandOptionTypeAttachment      ApplicationCommandOptionType = 11
)

type ApplicationCommandOptionChoice struct {
	Name              string            `json:"name"`
	NameLocalizations map[string]string `json:"name_localizations,omitempty"`
	Value             interface{}       `json:"value"`
}
