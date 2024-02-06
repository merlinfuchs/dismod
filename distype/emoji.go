package distype

type Emoji struct {
	ID             Nullable[Snowflake] `json:"id"`
	Name           Nullable[string]    `json:"name"`
	Roles          []Snowflake         `json:"roles,omitempty"`
	User           Optional[User]      `json:"user,omitempty"`
	RequiredColons Optional[bool]      `json:"required_colons,omitempty"`
	Managed        Optional[bool]      `json:"managed,omitempty"`
	Animated       Optional[bool]      `json:"animated,omitempty"`
	Available      Optional[bool]      `json:"available,omitempty"`
}

type GuildEmojisUpdateEvent struct {
	GuildID Snowflake `json:"guild_id"`
	Emojis  []Emoji   `json:"emojis"`
}

type GuildEmojiListRequest struct {
	GuildID Snowflake `json:"guild_id"`
}

type GuildEmojiListResponse = []Emoji

type GuildEmojiGetRequest struct {
	GuildID Snowflake `json:"guild_id"`
	EmojiID Snowflake `json:"emoji_id"`
}

type GuildEmojiGetResponse = Emoji

type GuildEmojiCreateRequest struct {
	GuildID Snowflake   `json:"guild_id"`
	Name    string      `json:"name"`
	Image   string      `json:"image"`
	Roles   []Snowflake `json:"roles,omitempty"`
}

type GuildEmojiCreateResponse = Emoji

type GuildEmojiModifyRequest struct {
	GuildID Snowflake   `json:"guild_id"`
	EmojiID Snowflake   `json:"emoji_id"`
	Name    string      `json:"name"`
	Roles   []Snowflake `json:"roles,omitempty"`
}

type GuildEmojiModifyResponse = Emoji

type GuildEmojiDeleteRequest struct {
	GuildID Snowflake `json:"guild_id"`
	EmojiID Snowflake `json:"emoji_id"`
}

type GuildEmojiDeleteResponse struct{}
