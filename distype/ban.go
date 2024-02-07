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

type BanListRequest struct {
	GuildID Snowflake           `json:"guild_id"`
	Limit   Optional[int]       `json:"limit,omitempty"`
	Before  Optional[Snowflake] `json:"before,omitempty"`
	After   Optional[Snowflake] `json:"after,omitempty"`
}

type BanListResponse = []Ban

type BanGetRequest struct {
	GuildID Snowflake `json:"guild_id"`
	UserID  Snowflake `json:"user_id"`
}

type BanGetResponse = Ban

type BanCreateRequest struct {
	GuildID              Snowflake     `json:"guild_id"`
	UserID               Snowflake     `json:"user_id"`
	DeleteMessageSeconds Optional[int] `json:"delete_message_seconds,omitempty"`
}

type BanCreateResponse = Ban

type BanRemoveRequest struct {
	GuildID Snowflake `json:"guild_id"`
	UserID  Snowflake `json:"user_id"`
}

type BanRemoveResponse = Ban
