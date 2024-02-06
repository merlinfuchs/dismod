package distype

import (
	"encoding/json"
	"time"
)

type Message struct {
	ID                   Snowflake                      `json:"id"`
	ChannelID            Snowflake                      `json:"channel_id"`
	Author               User                           `json:"author"`
	Content              string                         `json:"content"`
	Timestamp            time.Time                      `json:"timestamp"`
	EditedTimestamp      Nullable[time.Time]            `json:"edited_timestamp"`
	TTS                  bool                           `json:"tts"`
	MentionEveryone      bool                           `json:"mention_everyone"`
	Mentions             []User                         `json:"mentions"`
	MentionRoles         []Snowflake                    `json:"mention_roles"`
	MentionChannels      []ChannelMention               `json:"mention_channels,omitempty"`
	Attachments          []Attachment                   `json:"attachments"`
	Embeds               []Embed                        `json:"embeds"`
	Reactions            []Reaction                     `json:"reactions,omitempty"`
	Nonce                Optional[IntOrString]          `json:"nonce,omitempty"`
	Pinned               bool                           `json:"pinned"`
	WebhookID            Optional[Snowflake]            `json:"webhook_id,omitempty"`
	Type                 MessageType                    `json:"type"`
	Activity             Optional[MessageActivity]      `json:"activity,omitempty"`
	Application          Optional[Application]          `json:"application,omitempty"`
	ApplicationID        Optional[Snowflake]            `json:"application_id,omitempty"`
	MessageReference     Optional[MessageReference]     `json:"message_reference,omitempty"`
	Flags                Optional[MessageFlags]         `json:"flags,omitempty"`
	ReferencedMessage    Optional[Nullable[Message]]    `json:"referenced_message,omitempty"`
	Interaction          Optional[MessageInteraction]   `json:"interaction,omitempty"`
	Thread               Optional[Channel]              `json:"thread,omitempty"`
	Components           []MessageComponent             `json:"components,omitempty"`
	StickerItems         []MessageStickerItem           `json:"sticker_items,omitempty"`
	Stickers             []Sticker                      `json:"stickers,omitempty"`
	Position             Optional[int]                  `json:"position,omitempty"`
	RoleSubscriptionData Optional[RoleSubscriptionData] `json:"role_subscription_data,omitempty"`
	Resolved             Optional[ResolvedData]         `json:"resolved,omitempty"`
}

func (m *Message) UnmarshalJSON(data []byte) error {
	type message Message
	var v struct {
		message
		RawComponents []unmarshalableMessageComponent `json:"components"`
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	*m = Message(v.message)
	m.Components = make([]MessageComponent, len(v.RawComponents))
	for i, v := range v.RawComponents {
		m.Components[i] = v.MessageComponent
	}
	return err
}

type MessageType int

const (
	MessageTypeDefault                                 MessageType = 0
	MessageTypeRecipientAdd                            MessageType = 1
	MessageTypeRecipientRemove                         MessageType = 2
	MessageTypeCall                                    MessageType = 3
	MessageTypeChannelNameChange                       MessageType = 4
	MessageTypeChannelIconChange                       MessageType = 5
	MessageTypeChannelPinnedMessage                    MessageType = 6
	MessageTypeUserJoin                                MessageType = 7
	MessageTypeGuildBoost                              MessageType = 8
	MessageTypeGuildBoostTier1                         MessageType = 9
	MessageTypeGuildBoostTier2                         MessageType = 10
	MessageTypeGuildBoostTier3                         MessageType = 11
	MessageTypeChannelFollowAdd                        MessageType = 12
	MessageTypeGuildDiscoveryDisqualified              MessageType = 14
	MessageTypeGuildDiscoveryRequalified               MessageType = 15
	MessageTypeGuildDiscoveryGracePeriodInitialWarning MessageType = 16
	MessageTypeGuildDiscoveryGracePeriodFinalWarning   MessageType = 17
	MessageTypeThreadCreated                           MessageType = 18
	MessageTypeReply                                   MessageType = 19
	MessageTypeChatInputCommand                        MessageType = 20
	MessageTypeThreadStarterMessage                    MessageType = 21
	MessageTypeGuildInviteReminder                     MessageType = 22
	MessageTypeContextMenuCommand                      MessageType = 23
	MessageTypeAutoModerationAction                    MessageType = 24
	MessageTypeRoleSubscriptionPurchase                MessageType = 25
	MessageTypeInteractionPremiumUpsell                MessageType = 26
	MessageTypeStageStart                              MessageType = 27
	MessageTypeStageEnd                                MessageType = 28
	MessageTypeStageSpeaker                            MessageType = 29
	MessageTypeStageTopic                              MessageType = 31
	MessageTypeGuildApplicationPremiumSubscription     MessageType = 32
)

type MessageActivity struct {
	Type    MessageActivityType `json:"type"`
	PartyID string              `json:"party_id"`
}

type MessageActivityType int

const (
	MessageActivityTypeJoin        MessageActivityType = 1
	MessageActivityTypeSpectate    MessageActivityType = 2
	MessageActivityTypeListen      MessageActivityType = 3
	MessageActivityTypeJoinRequest MessageActivityType = 5
)

type MessageFlags int

const (
	MessageFlagsCrossposted                      MessageFlags = 1 << 0
	MessageFlagsIsCrosspost                      MessageFlags = 1 << 1
	MessageFlagsSuppressEmbeds                   MessageFlags = 1 << 2
	MessageFlagsSourceMessageDeleted             MessageFlags = 1 << 3
	MessageFlagsUrgent                           MessageFlags = 1 << 4
	MessageFlagsHasThread                        MessageFlags = 1 << 5
	MessageFlagsEphemeral                        MessageFlags = 1 << 6
	MessageFlagsLoading                          MessageFlags = 1 << 7
	MessageFlagsFailedToMentionSomeRolesInThread MessageFlags = 1 << 8
	MessageFlagsSuppressNotifications            MessageFlags = 1 << 12
	MessageFlagsIsVoiceMessage                   MessageFlags = 1 << 13
)

type MessageReference struct {
	MessageID       Optional[Snowflake] `json:"message_id,omitempty"`
	ChannelID       Optional[Snowflake] `json:"channel_id,omitempty"`
	GuildID         Optional[Snowflake] `json:"guild_id,omitempty"`
	FailIfNotExists Optional[bool]      `json:"fail_if_not_exists,omitempty"`
}

type Reaction struct {
	Count        int                  `json:"count"`
	CountDetails ReactionCountDetails `json:"count_details"`
	Me           bool                 `json:"me"`
	MeBurst      bool                 `json:"me_burst"`
	Emoji        Emoji                `json:"emoji"`
	BurstColors  []string             `json:"burst_colors"`
}

type ReactionCountDetails struct {
	Burst  int `json:"burst"`
	Normal int `json:"normal"`
}

type ChannelMention struct {
	ID      Snowflake   `json:"id"`
	GuildID Snowflake   `json:"guild_id"`
	Type    ChannelType `json:"type"`
	Name    string      `json:"name"`
}

type Attachment struct {
	ID          Snowflake                 `json:"id"`
	Filename    string                    `json:"filename"`
	Description Optional[string]          `json:"description,omitempty"`
	ContentType Optional[string]          `json:"content_type,omitempty"`
	Size        int                       `json:"size"`
	URL         string                    `json:"url"`
	ProxyURL    string                    `json:"proxy_url"`
	Height      Optional[int]             `json:"height,omitempty"`
	Width       Optional[int]             `json:"width,omitempty"`
	Ephemeral   Optional[bool]            `json:"ephemeral,omitempty"`
	DurationSec Optional[float64]         `json:"duration_secs,omitempty"`
	Waveform    Optional[string]          `json:"waveform,omitempty"`
	Flags       Optional[AttachmentFlags] `json:"flags,omitempty"`
}

type AttachmentFlags int

const (
	AttachmentFlagsIsRemix AttachmentFlags = 1 << 2
)

type MessageInteraction struct {
	ID     Snowflake        `json:"id"`
	Type   InteractionType  `json:"type"`
	Name   string           `json:"name"`
	User   User             `json:"user"`
	Member Optional[Member] `json:"member,omitempty"`
}

type MessageStickerItem struct {
	ID         Snowflake         `json:"id"`
	Name       string            `json:"name"`
	FormatType StickerFormatType `json:"format_type"`
}

type RoleSubscriptionData struct {
	RoleSubscriptionListingID Snowflake `json:"role_subscription_listing_id"`
	TierName                  string    `json:"tier_name"`
	TotalMonthsSubscribed     int       `json:"total_months_subscribed"`
	IsRenewal                 bool      `json:"is_renewal"`
}

type Embed struct {
	Title       Optional[string]         `json:"title,omitempty"`
	Type        Optional[EmbedType]      `json:"type,omitempty"`
	Description Optional[string]         `json:"description,omitempty"`
	URL         Optional[string]         `json:"url,omitempty"`
	Timestamp   Optional[time.Time]      `json:"timestamp,omitempty"`
	Color       Optional[int]            `json:"color,omitempty"`
	Footer      Optional[EmbedFooter]    `json:"footer,omitempty"`
	Image       Optional[EmbedImage]     `json:"image,omitempty"`
	Thumbnail   Optional[EmbedThumbnail] `json:"thumbnail,omitempty"`
	Video       Optional[EmbedVideo]     `json:"video,omitempty"`
	Provider    Optional[EmbedProvider]  `json:"provider,omitempty"`
	Author      Optional[EmbedAuthor]    `json:"author,omitempty"`
	Fields      []EmbedField             `json:"fields,omitempty"`
}

type EmbedType string

const (
	EmbedTypeRich    EmbedType = "rich"
	EmbedTypeImage   EmbedType = "image"
	EmbedTypeVideo   EmbedType = "video"
	EmbedTypeGifv    EmbedType = "gifv"
	EmbedTypeArticle EmbedType = "article"
	EmbedTypeLink    EmbedType = "link"
)

type EmbedThumbnail struct {
	URL      string           `json:"url"`
	ProxyURL Optional[string] `json:"proxy_url,omitempty"`
	Height   Optional[int]    `json:"height,omitempty"`
	Width    Optional[int]    `json:"width,omitempty"`
}

type EmbedVideo struct {
	URL      Optional[string] `json:"url,omitempty"`
	ProxyURL Optional[string] `json:"proxy_url,omitempty"`
	Height   Optional[int]    `json:"height,omitempty"`
	Width    Optional[int]    `json:"width,omitempty"`
}

type EmbedImage struct {
	URL      string           `json:"url"`
	ProxyURL Optional[string] `json:"proxy_url,omitempty"`
	Height   Optional[int]    `json:"height,omitempty"`
	Width    Optional[int]    `json:"width,omitempty"`
}

type EmbedProvider struct {
	Name Optional[string] `json:"name,omitempty"`
	URL  Optional[string] `json:"url,omitempty"`
}

type EmbedAuthor struct {
	Name         string           `json:"name"`
	URL          Optional[string] `json:"url,omitempty"`
	IconURL      Optional[string] `json:"icon_url,omitempty"`
	ProxyIconURL Optional[string] `json:"proxy_icon_url,omitempty"`
}

type EmbedFooter struct {
	Text         string           `json:"text"`
	IconURL      Optional[string] `json:"icon_url,omitempty"`
	ProxyIconURL Optional[string] `json:"proxy_icon_url,omitempty"`
}

type EmbedField struct {
	Name   string         `json:"name"`
	Value  string         `json:"value"`
	Inline Optional[bool] `json:"inline,omitempty"`
}

type MessageCreateParams struct{}
