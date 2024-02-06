package distype

type Sticker struct {
	ID          Snowflake           `json:"id"`
	PackID      Optional[string]    `json:"pack_id,omitempty"`
	Name        string              `json:"name"`
	Description Optional[string]    `json:"description,omitempty"`
	Tags        string              `json:"tags"`
	Asset       Optional[string]    `json:"asset,omitempty"`
	Type        StickerType         `json:"type"`
	FormatType  StickerFormatType   `json:"format_type"`
	Available   Optional[bool]      `json:"available,omitempty"`
	GuildID     Optional[Snowflake] `json:"guild_id,omitempty"`
	User        Optional[User]      `json:"user,omitempty"`
	SortValue   Optional[int]       `json:"sort_value,omitempty"`
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