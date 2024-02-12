package distype

type Sticker struct {
	ID          Snowflake         `json:"id"`
	PackID      *string           `json:"pack_id,omitempty"`
	Name        string            `json:"name"`
	Description *string           `json:"description,omitempty"`
	Tags        string            `json:"tags"`
	Asset       *string           `json:"asset,omitempty"`
	Type        StickerType       `json:"type"`
	FormatType  StickerFormatType `json:"format_type"`
	Available   *bool             `json:"available,omitempty"`
	GuildID     *Snowflake        `json:"guild_id,omitempty"`
	User        *User             `json:"user,omitempty"`
	SortValue   *int              `json:"sort_value,omitempty"`
}

type StickerType int

const (
	StickerTypeStandard StickerType = 1
	StickerTypeGuild    StickerType = 2
)

type StickerFormatType int

const (
	StickerFormatTypePNG    StickerFormatType = 1
	StickerFormatTypeAPNG   StickerFormatType = 2
	StickerFormatTypeLOTTIE StickerFormatType = 3
	StickerFormatTypeGIF    StickerFormatType = 4
)

type GuildStickersUpdateEvent struct {
	GuildID  Snowflake `json:"guild_id"`
	Stickers []Sticker `json:"stickers"`
}

type GuildStickerListRequest struct {
	GuildID Snowflake `json:"guild_id"`
}

type GuildStickerListResponse = []Sticker

type StickerGetRequest struct {
	GuildID   Snowflake `json:"guild_id"`
	StickerID Snowflake `json:"sticker_id"`
}

type StickerGetResponse = Sticker

type StickerCreateRequest struct {
	GuildID     Snowflake `json:"guild_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Tags        string    `json:"tags"`
	File        string    `json:"file"`
}

type StickerCreateResponse = Sticker

type StickerModifyRequest struct {
	GuildID     Snowflake `json:"guild_id"`
	StickerID   Snowflake `json:"sticker_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Tags        string    `json:"tags"`
}

type StickerModifyResponse = Sticker

type StickerDeleteRequest struct {
	GuildID   Snowflake `json:"guild_id"`
	StickerID Snowflake `json:"sticker_id"`
}

type StickerDeleteResponse struct{}
