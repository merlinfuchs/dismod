package distype

import (
	"encoding/base64"
	"encoding/json"
	"time"
)

type Message struct {
	ID                   Snowflake             `json:"id"`
	ChannelID            Snowflake             `json:"channel_id"`
	Author               User                  `json:"author"`
	Content              string                `json:"content"`
	Timestamp            time.Time             `json:"timestamp"`
	EditedTimestamp      Nullable[time.Time]   `json:"edited_timestamp"`
	TTS                  bool                  `json:"tts"`
	MentionEveryone      bool                  `json:"mention_everyone"`
	Mentions             []User                `json:"mentions"`
	MentionRoles         []Snowflake           `json:"mention_roles"`
	MentionChannels      []ChannelMention      `json:"mention_channels,omitempty"`
	Attachments          []Attachment          `json:"attachments"`
	Embeds               []Embed               `json:"embeds"`
	Reactions            []Reaction            `json:"reactions,omitempty"`
	Nonce                *IntOrString          `json:"nonce,omitempty"`
	Pinned               bool                  `json:"pinned"`
	WebhookID            *Snowflake            `json:"webhook_id,omitempty"`
	Type                 MessageType           `json:"type"`
	Activity             *MessageActivity      `json:"activity,omitempty"`
	Application          *Application          `json:"application,omitempty"`
	ApplicationID        *Snowflake            `json:"application_id,omitempty"`
	MessageReference     *MessageReference     `json:"message_reference,omitempty"`
	Flags                *MessageFlags         `json:"flags,omitempty"`
	ReferencedMessage    *Nullable[Message]    `json:"referenced_message,omitempty"`
	Interaction          *MessageInteraction   `json:"interaction,omitempty"`
	Thread               *Channel              `json:"thread,omitempty"`
	Components           []MessageComponent    `json:"components,omitempty"`
	StickerItems         []MessageStickerItem  `json:"sticker_items,omitempty"`
	Stickers             []Sticker             `json:"stickers,omitempty"`
	Position             *int                  `json:"position,omitempty"`
	RoleSubscriptionData *RoleSubscriptionData `json:"role_subscription_data,omitempty"`
	Resolved             *ResolvedData         `json:"resolved,omitempty"`
	// Only set for message events
	GuildID *Snowflake `json:"guild_id,omitempty"`
	Member  *Member    `json:"member,omitempty"`
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
	MessageID       *Snowflake `json:"message_id,omitempty"`
	ChannelID       *Snowflake `json:"channel_id,omitempty"`
	GuildID         *Snowflake `json:"guild_id,omitempty"`
	FailIfNotExists *bool      `json:"fail_if_not_exists,omitempty"`
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
	ID          Snowflake        `json:"id"`
	Filename    string           `json:"filename"`
	Description *string          `json:"description,omitempty"`
	ContentType *string          `json:"content_type,omitempty"`
	Size        int              `json:"size"`
	URL         string           `json:"url"`
	ProxyURL    string           `json:"proxy_url"`
	Height      *int             `json:"height,omitempty"`
	Width       *int             `json:"width,omitempty"`
	Ephemeral   *bool            `json:"ephemeral,omitempty"`
	DurationSec *float64         `json:"duration_secs,omitempty"`
	Waveform    *string          `json:"waveform,omitempty"`
	Flags       *AttachmentFlags `json:"flags,omitempty"`
}

type AttachmentFlags int

const (
	AttachmentFlagsIsRemix AttachmentFlags = 1 << 2
)

type MessageInteraction struct {
	ID     Snowflake       `json:"id"`
	Type   InteractionType `json:"type"`
	Name   string          `json:"name"`
	User   User            `json:"user"`
	Member *Member         `json:"member,omitempty"`
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
	Title       *string         `json:"title,omitempty"`
	Type        *EmbedType      `json:"type,omitempty"`
	Description *string         `json:"description,omitempty"`
	URL         *string         `json:"url,omitempty"`
	Timestamp   *time.Time      `json:"timestamp,omitempty"`
	Color       *int            `json:"color,omitempty"`
	Footer      *EmbedFooter    `json:"footer,omitempty"`
	Image       *EmbedImage     `json:"image,omitempty"`
	Thumbnail   *EmbedThumbnail `json:"thumbnail,omitempty"`
	Video       *EmbedVideo     `json:"video,omitempty"`
	Provider    *EmbedProvider  `json:"provider,omitempty"`
	Author      *EmbedAuthor    `json:"author,omitempty"`
	Fields      []EmbedField    `json:"fields,omitempty"`
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
	URL      string  `json:"url"`
	ProxyURL *string `json:"proxy_url,omitempty"`
	Height   *int    `json:"height,omitempty"`
	Width    *int    `json:"width,omitempty"`
}

type EmbedVideo struct {
	URL      *string `json:"url,omitempty"`
	ProxyURL *string `json:"proxy_url,omitempty"`
	Height   *int    `json:"height,omitempty"`
	Width    *int    `json:"width,omitempty"`
}

type EmbedImage struct {
	URL      string  `json:"url"`
	ProxyURL *string `json:"proxy_url,omitempty"`
	Height   *int    `json:"height,omitempty"`
	Width    *int    `json:"width,omitempty"`
}

type EmbedProvider struct {
	Name *string `json:"name,omitempty"`
	URL  *string `json:"url,omitempty"`
}

type EmbedAuthor struct {
	Name         string  `json:"name"`
	URL          *string `json:"url,omitempty"`
	IconURL      *string `json:"icon_url,omitempty"`
	ProxyIconURL *string `json:"proxy_icon_url,omitempty"`
}

type EmbedFooter struct {
	Text         string  `json:"text"`
	IconURL      *string `json:"icon_url,omitempty"`
	ProxyIconURL *string `json:"proxy_icon_url,omitempty"`
}

type EmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline *bool  `json:"inline,omitempty"`
}

type AllowedMentions struct {
	Parse       []AllowedMentionType `json:"parse"`
	Roles       []Snowflake          `json:"roles"`
	Users       []Snowflake          `json:"users"`
	RepliedUser bool                 `json:"replied_user"`
}

type AllowedMentionType string

const (
	AllowedMentionTypeRoleMentions    AllowedMentionType = "roles"
	AllowedMentionTypeUserMentions    AllowedMentionType = "users"
	AllowedMentionTypeEveryoneMention AllowedMentionType = "everyone"
)

type TypingStartEvent struct {
	ChannelID Snowflake     `json:"channel_id"`
	GuildID   *Snowflake    `json:"guild_id,omitempty"`
	UserID    Snowflake     `json:"user_id"`
	Timestamp UnixTimestamp `json:"timestamp"`
	Member    *Member       `json:"member,omitempty"`
}

type MessageCreateEvent = Message

type MessageUpdateEvent = Message

type MessageDeleteEvent struct {
	ID        Snowflake  `json:"id"`
	ChannelID Snowflake  `json:"channel_id"`
	GuildID   *Snowflake `json:"guild_id,omitempty"`
}

type MessageDeleteBulkEvent struct {
	IDs       []Snowflake `json:"ids"`
	ChannelID Snowflake   `json:"channel_id"`
	GuildID   *Snowflake  `json:"guild_id,omitempty"`
}

type MessageReactionAddEvent struct {
	UserID          Snowflake  `json:"user_id"`
	ChannelID       Snowflake  `json:"channel_id"`
	MessageID       Snowflake  `json:"message_id"`
	GuildID         *Snowflake `json:"guild_id,omitempty"`
	Member          *Member    `json:"member,omitempty"`
	Emoji           Emoji      `json:"emoji"`
	MessageAuthorID *Snowflake `json:"message_author_id,omitempty"`
}

type MessageReactionRemoveEvent struct {
	UserID    Snowflake  `json:"user_id"`
	ChannelID Snowflake  `json:"channel_id"`
	MessageID Snowflake  `json:"message_id"`
	GuildID   *Snowflake `json:"guild_id,omitempty"`
	Emoji     Emoji      `json:"emoji"`
}

type MessageReactionRemoveAllEvent struct {
	ChannelID Snowflake  `json:"channel_id"`
	MessageID Snowflake  `json:"message_id"`
	GuildID   *Snowflake `json:"guild_id,omitempty"`
}

type MessageReactionRemoveEmojiEvent struct {
	ChannelID Snowflake  `json:"channel_id"`
	MessageID Snowflake  `json:"message_id"`
	GuildID   *Snowflake `json:"guild_id,omitempty"`
	Emoji     Emoji      `json:"emoji"`
}

type MessageCreateParams struct {
	Content          *string            `json:"content,omitempty"`
	Nonce            *string            `json:"nonce,omitempty"`
	TTS              *bool              `json:"tts,omitempty"`
	Embeds           []Embed            `json:"embeds,omitempty"`
	AllowedMentions  *AllowedMentions   `json:"allowed_mentions,omitempty"`
	MessageReference *MessageReference  `json:"message_reference,omitempty"`
	Components       []MessageComponent `json:"components,omitempty"`
	StickerIDs       []Snowflake        `json:"sticker_ids,omitempty"`
	Attachments      []Attachment       `json:"attachments,omitempty"`
	Files            []File             `json:"files,omitempty"`
	Flags            *MessageFlags      `json:"flags,omitempty"`
}

type MessageEditParams struct {
	Content          *string            `json:"content,omitempty"`
	Embeds           []Embed            `json:"embeds,omitempty"`
	AllowedMentions  *AllowedMentions   `json:"allowed_mentions,omitempty"`
	MessageReference *MessageReference  `json:"message_reference,omitempty"`
	Components       []MessageComponent `json:"components,omitempty"`
	StickerIDs       []Snowflake        `json:"sticker_ids,omitempty"`
	Attachments      []Attachment       `json:"attachments,omitempty"`
	Files            []File             `json:"files,omitempty"`
	Flags            *MessageFlags      `json:"flags,omitempty"`
}

type File struct {
	Filename    string `json:"filename"`
	ContentType string `json:"content_type"`
	Data        []byte `json:"data"`
}

type rawFile struct {
	Filename    string `json:"filename"`
	ContentType string `json:"content_type"`
	Data        string `json:"data"`
}

func (f *File) MarshalJSON() ([]byte, error) {
	base64Data := base64.StdEncoding.EncodeToString(f.Data)

	return json.Marshal(rawFile{
		Filename:    f.Filename,
		ContentType: f.ContentType,
		Data:        base64Data,
	})
}

func (f *File) UnmarshalJSON(data []byte) error {
	var rf rawFile
	err := json.Unmarshal(data, &rf)
	if err != nil {
		return err
	}

	f.Filename = rf.Filename
	f.ContentType = rf.ContentType
	f.Data, err = base64.StdEncoding.DecodeString(rf.Data)
	return err
}

type ChannelMessageListRequest struct {
	ChannelID Snowflake  `json:"channel_id"`
	Around    *Snowflake `json:"around,omitempty"`
	Before    *Snowflake `json:"before,omitempty"`
	After     *Snowflake `json:"after,omitempty"`
	Limit     *int       `json:"limit,omitempty"`
}

type ChannelMessageListResponse = []Message

type MessageGetRequest struct {
	ChannelID Snowflake `json:"channel_id"`
	MessageID Snowflake `json:"message_id"`
}

type MessageGetResponse = Message

type MessageCreateRequest struct {
	ChannelID Snowflake `json:"channel_id"`
	MessageCreateParams
}

type MessageCreateResponse = Message

type MessageCrosspostRequest struct {
	ChannelID Snowflake `json:"channel_id"`
	MessageID Snowflake `json:"message_id"`
}

type MessageCrosspostResponse = Message

type MessageEditRequest struct {
	ChannelID Snowflake `json:"channel_id"`
	MessageID Snowflake `json:"message_id"`
	MessageEditParams
}

type MessageEditResponse = Message

type MessageDeleteRequest struct {
	ChannelID Snowflake `json:"channel_id"`
	MessageID Snowflake `json:"message_id"`
}

type MessageDeleteResponse struct{}

type MessageBulkDeleteRequest struct {
	ChannelID Snowflake   `json:"channel_id"`
	Messages  []Snowflake `json:"messages"`
}

type MessageBulkDeleteResponse struct{}

type MessageReactionListRequest struct {
	ChannelID Snowflake  `json:"channel_id"`
	MessageID Snowflake  `json:"message_id"`
	Emoji     string     `json:"emoji"`
	After     *Snowflake `json:"after,omitempty"`
	Limit     *int       `json:"limit,omitempty"`
}

type MessageReactionListResponse = []Reaction

type MessageReactionCreateRequest struct {
	ChannelID Snowflake `json:"channel_id"`
	MessageID Snowflake `json:"message_id"`
	Emoji     string    `json:"emoji"`
}

type MessageReactionCreateResponse struct{}

type MessageReactionDeleteOwnRequest struct {
	ChannelID Snowflake `json:"channel_id"`
	MessageID Snowflake `json:"message_id"`
	Emoji     string    `json:"emoji"`
}

type MessageReactionDeleteOwnResponse struct{}

type MessageReactionDeleteRequest struct {
	ChannelID Snowflake `json:"channel_id"`
	MessageID Snowflake `json:"message_id"`
	Emoji     string    `json:"emoji"`
	UserID    Snowflake `json:"user_id"`
}

type MessageReactionDeleteResponse struct{}

type MessageReactionDeleteAllRequest struct {
	ChannelID Snowflake `json:"channel_id"`
	MessageID Snowflake `json:"message_id"`
}

type MessageReactionDeleteAllResponse struct{}

type MessageReactionDeleteEmojiRequest struct {
	ChannelID Snowflake `json:"channel_id"`
	MessageID Snowflake `json:"message_id"`
	Emoji     string    `json:"emoji"`
}

type MessageReactionDeleteEmojiResponse struct{}

type ChannelPinnedMessageListRequest struct {
	ChannelID Snowflake `json:"channel_id"`
}

type ChannelPinnedMessageListResponse = []Message

type MessagePinRequest struct {
	ChannelID Snowflake `json:"channel_id"`
	MessageID Snowflake `json:"message_id"`
}

type MessagePinResponse struct{}

type MessageUnpinRequest struct {
	ChannelID Snowflake `json:"channel_id"`
	MessageID Snowflake `json:"message_id"`
}

type MessageUnpinResponse struct{}
