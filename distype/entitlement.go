package distype

type Entitlement struct {
	ID            Snowflake           `json:"id"`
	SKUID         Snowflake           `json:"sku_id"`
	ApplicationID Snowflake           `json:"application_id"`
	UserID        Optional[Snowflake] `json:"user_id,omitempty"`
	Type          EntitlementType     `json:"type"`
	Deleted       bool                `json:"deleted"`
	StartsAt      Optional[string]    `json:"starts_at,omitempty"`
	EndsAt        Optional[string]    `json:"ends_at,omitempty"`
	GuildID       Optional[Snowflake] `json:"guild_id,omitempty"`
}

type EntitlementType int

const (
	EntitlementTypeApplicationSubscription EntitlementType = 8
)
