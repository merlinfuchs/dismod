package distype

type Webhook struct {
	ID            Snowflake           `json:"id"`
	Type          WebhookType         `json:"type"`
	GuildID       Nullable[Snowflake] `json:"guild_id"`
	ChannelID     Nullable[Snowflake] `json:"channel_id"`
	User          Optional[User]      `json:"user,omitempty"`
	Name          Nullable[string]    `json:"name"`
	Avatar        Nullable[string]    `json:"avatar"`
	Token         Optional[string]    `json:"token,omitempty"`
	ApplicationID Nullable[Snowflake] `json:"application_id"`
	SourceGuild   Optional[Guild]     `json:"source_guild,omitempty"`
	SoureChannel  Optional[Channel]   `json:"source_channel,omitempty"`
	URL           Optional[string]    `json:"url,omitempty"`
}

type WebhookType int

const (
	WebhookTypeIncoming        WebhookType = 1
	WebhookTypeChannelFollower WebhookType = 2
	WebhookTypeApplication     WebhookType = 3
)

type WebhooksUpdateEvent struct {
	GuildID   Snowflake `json:"guild_id"`
	ChannelID Snowflake `json:"channel_id"`
}

type WebhookCreateRequest struct {
	ChannelID Snowflake                  `json:"channel_id"`
	Name      string                     `json:"name"`
	Avatar    Optional[Nullable[string]] `json:"avatar,omitempty"`
}

type WebhookCreateResponse = Webhook

type GuildWebhookListRequest struct {
	GuildID Snowflake `json:"guild_id"`
}

type GuildWebhookListResponse = []Webhook

type ChannelWebhookListRequest struct {
	ChannelID Snowflake `json:"channel_id"`
}

type ChannelWebhookListResponse = []Webhook

type WebhookGetRequest struct {
	WebhookID Snowflake `json:"webhook_id"`
}

type WebhookGetResponse = Webhook

type WebhookGetWithTokenRequest struct {
	WebhookID    Snowflake `json:"webhook_id"`
	WebhookToken string    `json:"webhook_token"`
}

type WebhookGetWithTokenResponse = Webhook

type WebhookModifyRequest struct {
	WebhookID Snowflake                  `json:"webhook_id"`
	Name      Optional[string]           `json:"name,omitempty"`
	Avatar    Optional[Nullable[string]] `json:"avatar,omitempty"`
	ChannelID Optional[Snowflake]        `json:"channel_id,omitempty"`
}

type WebhookModifyResponse = Webhook

type WebhookModifyWithTokenRequest struct {
	WebhookID    Snowflake                  `json:"webhook_id"`
	WebhookToken string                     `json:"webhook_token"`
	Name         Optional[string]           `json:"name,omitempty"`
	Avatar       Optional[Nullable[string]] `json:"avatar,omitempty"`
	ChannelID    Optional[Snowflake]        `json:"channel_id,omitempty"`
}

type WebhookModifyWithTokenResponse = Webhook

type WebhookDeleteRequest struct {
	WebhookID Snowflake `json:"webhook_id"`
}

type WebhookDeleteResponse struct{}

type WebhookDeleteWithTokenRequest struct {
	WebhookID    Snowflake `json:"webhook_id"`
	WebhookToken string    `json:"webhook_token"`
}

type WebhookDeleteWithTokenResponse struct{}

type WebhookExecuteRequest struct {
	WebhookID    Snowflake           `json:"webhook_id"`
	WebhookToken string              `json:"webhook_token"`
	Wait         Optional[bool]      `json:"wait,omitempty"`
	ThreadID     Optional[Snowflake] `json:"thread_id,omitempty"`
	MessageCreateParams
}

type WebhookExecuteResponse = Optional[Message]

type WebhookMessageGetRequest struct {
	WebhookID    Snowflake `json:"webhook_id"`
	WebhookToken string    `json:"webhook_token"`
	MessageID    Snowflake `json:"message_id"`
}

type WebhookMessageGetResponse = Message

type WebhookMessageEditRequest struct {
	WebhookID    Snowflake `json:"webhook_id"`
	WebhookToken string    `json:"webhook_token"`
	MessageID    Snowflake `json:"message_id"`
	MessageEditParams
}

type WebhookMessageEditResponse = Message

type WebhookMessageDeleteRequest struct {
	WebhookID    Snowflake `json:"webhook_id"`
	WebhookToken string    `json:"webhook_token"`
	MessageID    Snowflake `json:"message_id"`
}

type WebhookMessageDeleteResponse = Message
