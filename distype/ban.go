package distype

type Ban struct {
	Reason Nullable[string] `json:"reason"`
	User   User             `json:"user"`
}

type BanAddEvent struct {
	GuildID Snowflake `json:"guild_id"`
	User    User      `json:"user"`
}

type BanRemoveEvent struct {
	GuildID Snowflake `json:"guild_id"`
	User    User      `json:"user"`
}
