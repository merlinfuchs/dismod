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
