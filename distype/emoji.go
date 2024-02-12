package distype

type Emoji struct {
	ID             Nullable[Snowflake] `json:"id"`
	Name           Nullable[string]    `json:"name"`
	Roles          []Snowflake         `json:"roles,omitempty"`
	User           *User               `json:"user,omitempty"`
	RequiredColons *bool               `json:"required_colons,omitempty"`
	Managed        *bool               `json:"managed,omitempty"`
	Animated       *bool               `json:"animated,omitempty"`
	Available      *bool               `json:"available,omitempty"`
}

type GuildEmojisUpdateEvent struct {
	GuildID Snowflake `json:"guild_id"`
	Emojis  []Emoji   `json:"emojis"`
}

type GuildEmojiListRequest struct {
	GuildID Snowflake `json:"guild_id"`
}

type GuildEmojiListResponse = []Emoji

type EmojiGetRequest struct {
	GuildID Snowflake `json:"guild_id"`
	EmojiID Snowflake `json:"emoji_id"`
}

type EmojiGetResponse = Emoji

type EmojiCreateRequest struct {
	GuildID Snowflake   `json:"guild_id"`
	Name    string      `json:"name"`
	Image   string      `json:"image"`
	Roles   []Snowflake `json:"roles,omitempty"`
}

type EmojiCreateResponse = Emoji

type EmojiModifyRequest struct {
	GuildID Snowflake   `json:"guild_id"`
	EmojiID Snowflake   `json:"emoji_id"`
	Name    string      `json:"name"`
	Roles   []Snowflake `json:"roles,omitempty"`
}

type EmojiModifyResponse = Emoji

type EmojiDeleteRequest struct {
	GuildID Snowflake `json:"guild_id"`
	EmojiID Snowflake `json:"emoji_id"`
}

type EmojiDeleteResponse struct{}
