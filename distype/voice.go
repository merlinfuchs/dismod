package distype

type VoiceState struct {
	// TODO
}

type VoiceStateUpdateEvent = VoiceState

type VoiceServerUpdateEvent struct {
	Token    string           `json:"token"`
	GuildID  Snowflake        `json:"guild_id"`
	Endpoint Nullable[string] `json:"endpoint"`
}
