package distype

import "encoding/json"

type Interaction struct {
	ID             Snowflake           `json:"id"`
	ApplicationID  Snowflake           `json:"application_id"`
	Type           InteractionType     `json:"type"`
	Data           InteractionData     `json:"data,omitempty"`
	GuildID        Optional[Snowflake] `json:"guild_id,omitempty"`
	Channel        Optional[Channel]   `json:"channel,omitempty"`
	ChannelID      Optional[Snowflake] `json:"channel_id,omitempty"`
	Member         Optional[Member]    `json:"member,omitempty"`
	User           Optional[User]      `json:"user,omitempty"`
	Token          string              `json:"token"`
	Version        int                 `json:"version"`
	Message        Optional[Message]   `json:"message,omitempty"`
	AppPermissions Optional[string]    `json:"app_permissions,omitempty"`
	Locale         Optional[string]    `json:"locale,omitempty"`
	GuildLocale    Optional[string]    `json:"guild_locale,omitempty"`
	Entitlements   []Entitlement       `json:"entitlements,omitempty"`
}

type interaction Interaction

type rawInteraction struct {
	interaction
	Data json.RawMessage `json:"data"`
}

func (i *Interaction) UnmarshalJSON(raw []byte) error {
	var tmp rawInteraction
	err := json.Unmarshal(raw, &tmp)
	if err != nil {
		return err
	}

	*i = Interaction(tmp.interaction)

	switch tmp.Type {
	case InteractionTypeApplicationCommand, InteractionTypeApplicationCommandAutocomplete:
		v := ApplicationCommandData{}
		err = json.Unmarshal(tmp.Data, &v)
		if err != nil {
			return err
		}
		i.Data = v
	case InteractionTypeMessageComponent:
		v := MessageComponentData{}
		err = json.Unmarshal(tmp.Data, &v)
		if err != nil {
			return err
		}
		i.Data = v
	case InteractionTypeModalSubmit:
		v := ModalSubmitData{}
		err = json.Unmarshal(tmp.Data, &v)
		if err != nil {
			return err
		}
		i.Data = v
	}
	return nil
}

type InteractionData interface {
	InteractionType() InteractionType
}

type InteractionType int

const (
	InteractionTypePing                           InteractionType = 1
	InteractionTypeApplicationCommand             InteractionType = 2
	InteractionTypeMessageComponent               InteractionType = 3
	InteractionTypeApplicationCommandAutocomplete InteractionType = 4
	InteractionTypeModalSubmit                    InteractionType = 5
)

type ApplicationCommandData struct {
	ID       Snowflake                  `json:"id"`
	Name     string                     `json:"name"`
	Type     ApplicationCommandType     `json:"type"`
	Resolved Optional[ResolvedData]     `json:"resolved,omitempty"`
	Options  []ApplicationCommandOption `json:"options,omitempty"`
	GuildID  Optional[Snowflake]        `json:"guild_id,omitempty"`
	TargetID Optional[Snowflake]        `json:"target_id,omitempty"`
}

func (ApplicationCommandData) InteractionType() InteractionType {
	return InteractionTypeApplicationCommand
}

type ApplicationCommandType int

const (
	ApplicationCommandTypeChatInput ApplicationCommandType = 1
	ApplicationCommandTypeUser      ApplicationCommandType = 2
	ApplicationCommandTypeMessage   ApplicationCommandType = 3
)

type ApplicationCommandOption struct {
	Name    string                       `json:"name"`
	Type    ApplicationCommandOptionType `json:"type"`
	Value   Optional[interface{}]        `json:"value,omitempty"`
	Options []ApplicationCommandOption   `json:"options,omitempty"`
	Focused Optional[bool]               `json:"focused,omitempty"`
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

type MessageComponentData struct {
	CustomID      string                 `json:"custom_id"`
	ComponentType MessageComponentType   `json:"component_type"`
	Values        []string               `json:"values,omitempty"`
	Resolved      Optional[ResolvedData] `json:"resolved,omitempty"`
}

func (MessageComponentData) InteractionType() InteractionType {
	return InteractionTypeMessageComponent
}

type ModalSubmitData struct {
	CustomID   string             `json:"custom_id"`
	Components []MessageComponent `json:"components"` // TODO: implement unmarshalJSON for MessageComponent
}

func (d *ModalSubmitData) UnmarshalJSON(data []byte) error {
	var v struct {
		CustomID      string                          `json:"custom_id"`
		RawComponents []unmarshalableMessageComponent `json:"components"`
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}

	d.CustomID = v.CustomID
	d.Components = make([]MessageComponent, len(v.RawComponents))
	for i, v := range v.RawComponents {
		d.Components[i] = v.MessageComponent
	}

	return err
}

func (ModalSubmitData) InteractionType() InteractionType {
	return InteractionTypeModalSubmit
}

type ResolvedData struct {
	Users       map[Snowflake]User       `json:"users,omitempty"`
	Members     map[Snowflake]Member     `json:"members,omitempty"`
	Roles       map[Snowflake]Role       `json:"roles,omitempty"`
	Channels    map[Snowflake]Channel    `json:"channels,omitempty"`
	Messages    map[Snowflake]Message    `json:"messages,omitempty"`
	Attachments map[Snowflake]Attachment `json:"attachments,omitempty"`
}
