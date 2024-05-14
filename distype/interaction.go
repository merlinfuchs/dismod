package distype

import "encoding/json"

type Interaction struct {
	ID             Snowflake       `json:"id"`
	ApplicationID  Snowflake       `json:"application_id"`
	Type           InteractionType `json:"type"`
	Data           InteractionData `json:"data,omitempty"`
	GuildID        *Snowflake      `json:"guild_id,omitempty"`
	Channel        *Channel        `json:"channel,omitempty"`
	ChannelID      *Snowflake      `json:"channel_id,omitempty"`
	Member         *Member         `json:"member,omitempty"`
	User           *User           `json:"user,omitempty"`
	Token          string          `json:"token"`
	Version        int             `json:"version"`
	Message        *Message        `json:"message,omitempty"`
	AppPermissions *string         `json:"app_permissions,omitempty"`
	Locale         *string         `json:"locale,omitempty"`
	GuildLocale    *string         `json:"guild_locale,omitempty"`
	Entitlements   []Entitlement   `json:"entitlements,omitempty"`
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
	ID       Snowflake                      `json:"id"`
	Name     string                         `json:"name"`
	Type     ApplicationCommandType         `json:"type"`
	Resolved *ResolvedData                  `json:"resolved,omitempty"`
	Options  []ApplicationCommandDataOption `json:"options,omitempty"`
	GuildID  *Snowflake                     `json:"guild_id,omitempty"`
	TargetID *Snowflake                     `json:"target_id,omitempty"`
}

func (ApplicationCommandData) InteractionType() InteractionType {
	return InteractionTypeApplicationCommand
}

type ApplicationCommandDataOption struct {
	Name    string                         `json:"name"`
	Type    ApplicationCommandOptionType   `json:"type"`
	Value   interface{}                    `json:"value,omitempty"`
	Options []ApplicationCommandDataOption `json:"options,omitempty"`
	Focused *bool                          `json:"focused,omitempty"`
}

type MessageComponentData struct {
	CustomID      string               `json:"custom_id"`
	ComponentType MessageComponentType `json:"component_type"`
	Values        []string             `json:"values,omitempty"`
	Resolved      *ResolvedData        `json:"resolved,omitempty"`
}

func (MessageComponentData) InteractionType() InteractionType {
	return InteractionTypeMessageComponent
}

type ModalSubmitData struct {
	CustomID   string             `json:"custom_id"`
	Components []MessageComponent `json:"components"`
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

type InteractionResponse struct {
	Type InteractionResponseType `json:"type"`
	Data InteractionResponseData `json:"data,omitempty"`
}

func (i *InteractionResponse) UnmarshalJSON(raw []byte) error {
	var tmp struct {
		Type InteractionResponseType `json:"type"`
		Data json.RawMessage         `json:"data,omitempty"`
	}
	err := json.Unmarshal(raw, &tmp)
	if err != nil {
		return err
	}

	i.Type = tmp.Type

	switch tmp.Type {
	case InteractionResponseTypeChannelMessageWithSource | InteractionResponseTypeDeferredChannelMessageWithSource:
		v := InteractionMessageCreateResponse{}
		err = json.Unmarshal(tmp.Data, &v)
		if err != nil {
			return err
		}
		i.Data = v
	case InteractionResponseTypeUpdateMessage | InteractionResponseTypeDeferredUpdateMessage:
		v := InteractionMessageUpdateResponse{}
		err = json.Unmarshal(tmp.Data, &v)
		if err != nil {
			return err
		}
		i.Data = v
	case InteractionResponseTypeApplicationCommandAutocompleteResult:
		v := InteractionAutocompleteResponse{}
		err = json.Unmarshal(tmp.Data, &v)
		if err != nil {
			return err
		}
		i.Data = v
	case InteractionResponseTypeModal:
		v := InteractionModalResponse{}
		err = json.Unmarshal(tmp.Data, &v)
		if err != nil {
			return err
		}
		i.Data = v
	}
	return nil
}

type InteractionResponseType int

const (
	InteractionResponseTypePong                                 InteractionResponseType = 1
	InteractionResponseTypeChannelMessageWithSource             InteractionResponseType = 4
	InteractionResponseTypeDeferredChannelMessageWithSource     InteractionResponseType = 5
	InteractionResponseTypeDeferredUpdateMessage                InteractionResponseType = 6
	InteractionResponseTypeUpdateMessage                        InteractionResponseType = 7
	InteractionResponseTypeApplicationCommandAutocompleteResult InteractionResponseType = 8
	InteractionResponseTypeModal                                InteractionResponseType = 9
	InteractionResponseTypePremiumRequired                      InteractionResponseType = 10
)

type InteractionResponseData interface {
	InteractionResponseType() InteractionResponseType
}

type InteractionMessageCreateResponse = MessageCreateParams

func (InteractionMessageCreateResponse) InteractionResponseType() InteractionResponseType {
	return InteractionResponseTypeChannelMessageWithSource
}

type InteractionMessageUpdateResponse = MessageEditParams

func (InteractionMessageUpdateResponse) InteractionResponseType() InteractionResponseType {
	return InteractionResponseTypeUpdateMessage
}

type InteractionAutocompleteResponse struct{} // TODO

func (InteractionAutocompleteResponse) InteractionResponseType() InteractionResponseType {
	return InteractionResponseTypeApplicationCommandAutocompleteResult
}

type InteractionModalResponse struct{} // TODO

func (InteractionModalResponse) InteractionResponseType() InteractionResponseType {
	return InteractionResponseTypeModal
}

type InteractionCreateEvent = Interaction

type InteractionResponseCreateRequest struct {
	InteractionID    Snowflake `json:"interaction_id"`
	InteractionToken string    `json:"interaction_token"`
	InteractionResponse
}

type InteractionResponseCreateResponse struct{}

type InteractionResponseGetRequest struct {
	ApplicationID    Snowflake `json:"application_id"`
	InteractionToken string    `json:"interaction_token"`
}

type InteractionResponseGetResponse = Message

type InteractionResponseEditRequest struct {
	ApplicationID     Snowflake `json:"application_id"`
	InteractionToken  string    `json:"interaction_token"`
	MessageEditParams `tstype:",extends"`
}

type InteractionResponseEditResponse = Message

type InteractionResponseDeleteRequest struct {
	ApplicationID    Snowflake `json:"application_id"`
	InteractionToken string    `json:"interaction_token"`
}

type InteractionResponseDeleteResponse struct{}

type InteractionFollowupCreateRequest struct {
	ApplicationID       Snowflake `json:"application_id"`
	InteractionToken    string    `json:"interaction_token"`
	MessageCreateParams `tstype:",extends"`
}

type InteractionFollowupCreateResponse = Message

type InteractionFollowupGetRequest struct {
	ApplicationID    Snowflake `json:"application_id"`
	InteractionToken string    `json:"interaction_token"`
	MessageID        Snowflake `json:"message_id"`
}

type InteractionFollowupGetResponse = Message

type InteractionFollowupEditRequest struct {
	ApplicationID     Snowflake `json:"application_id"`
	InteractionToken  string    `json:"interaction_token"`
	MessageID         Snowflake `json:"message_id"`
	MessageEditParams `tstype:",extends"`
}

type InteractionFollowupEditResponse = Message

type InteractionFollowupDeleteRequest struct {
	ApplicationID    Snowflake `json:"application_id"`
	InteractionToken string    `json:"interaction_token"`
	MessageID        Snowflake `json:"message_id"`
}

type InteractionFollowupDeleteResponse struct{}
