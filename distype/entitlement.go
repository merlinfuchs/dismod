package distype

type Entitlement struct {
	ID            Snowflake       `json:"id"`
	SKUID         Snowflake       `json:"sku_id"`
	ApplicationID Snowflake       `json:"application_id"`
	UserID        *Snowflake      `json:"user_id,omitempty"`
	Type          EntitlementType `json:"type"`
	Deleted       bool            `json:"deleted"`
	StartsAt      *string         `json:"starts_at,omitempty"`
	EndsAt        *string         `json:"ends_at,omitempty"`
	GuildID       *Snowflake      `json:"guild_id,omitempty"`
}

type EntitlementType int

const (
	EntitlementTypeApplicationSubscription EntitlementType = 8
)

type EntitlementCreateEvent = Entitlement

type EntitlementUpdateEvent = Entitlement

type EntitlementDeleteEvent = Entitlement
