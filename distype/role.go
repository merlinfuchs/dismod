package distype

type Role struct {
	ID           Snowflake                  `json:"id"`
	Name         string                     `json:"name"`
	Color        int                        `json:"color"`
	Hoist        bool                       `json:"hoist"`
	Icon         Optional[Nullable[string]] `json:"icon,omitempty"`
	UnicodeEmoji Optional[Nullable[string]] `json:"unicode_emoji,omitempty"`
	Position     int                        `json:"position"`
	Permissions  string                     `json:"permissions"`
	Managed      bool                       `json:"managed"`
	Mentionable  bool                       `json:"mentionable"`
	Tags         Optional[RoleTags]         `json:"tags,omitempty"`
	Flags        RoleFlags                  `json:"flags"`
}

type RoleTags struct {
	BotID                 Optional[Snowflake]          `json:"bot_id"`
	IntegrationID         Optional[Snowflake]          `json:"integration_id"`
	PremiumSubscriber     Optional[Nullable[struct{}]] `json:"premium_subscriber,omitempty"`
	SubscriptionListingID Optional[Snowflake]          `json:"subscription_listing_id"`
	AvailableForPurchase  Optional[Nullable[struct{}]] `json:"available_for_purchase,omitempty"`
	GuildConnections      Optional[Nullable[struct{}]] `json:"guild_connections,omitempty"`
}

type RoleFlags int

const (
	RoleFlagsInPrompt RoleFlags = 1 << 0
)
