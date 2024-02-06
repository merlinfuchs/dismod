package distype

type StageInstance struct {
	ID                    Snowflake    `json:"id"`
	GuildID               Snowflake    `json:"guild_id"`
	ChannelID             Snowflake    `json:"channel_id"`
	Topic                 string       `json:"topic"`
	PrivacyLevel          PrivacyLevel `json:"privacy_level"`
	DiscoverableDisabled  bool         `json:"discoverable_disabled"`
	GuildScheduledEventID Snowflake    `json:"guild_scheduled_event_id"`
}

type PrivacyLevel int

const (
	PrivacyLevelPublic    PrivacyLevel = 1
	PrivacyLevelGuildOnly PrivacyLevel = 2
)

type StageInstanceCreateEvent = StageInstance

type StageInstanceUpdateEvent = StageInstance

type StageInstanceDeleteEvent = StageInstance

type StageInstanceCreateRequest struct {
	ChannelID             Snowflake              `json:"channel_id"`
	Topic                 string                 `json:"topic"`
	PrivacyLevel          Optional[PrivacyLevel] `json:"privacy_level,omitempty"`
	SendStartNotification Optional[bool]         `json:"send_start_notification,omitempty"`
	GuildScheduledEventID Optional[Snowflake]    `json:"guild_scheduled_event_id,omitempty"`
}

type StageInstanceCreateResponse = StageInstance

type StageInstanceModifyRequest struct {
	ChannelID    Snowflake              `json:"channel_id"`
	Topic        Optional[string]       `json:"topic,omitempty"`
	PrivacyLevel Optional[PrivacyLevel] `json:"privacy_level,omitempty"`
}

type StageInstanceModifyResponse = StageInstance

type StageInstanceDeleteRequest struct {
	ChannelID Snowflake `json:"channel_id"`
}

type StageInstanceDeleteResponse struct{}

type StageInstanceGetRequest struct {
	ChannelID Snowflake `json:"channel_id"`
}

type StageInstanceGetResponse = StageInstance
