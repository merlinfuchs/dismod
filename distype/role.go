package distype

type Role struct {
	ID           Snowflake                  `json:"id"`
	Name         string                     `json:"name"`
	Color        int                        `json:"color"`
	Hoist        bool                       `json:"hoist"`
	Icon         Optional[Nullable[string]] `json:"icon,omitempty"`
	UnicodeEmoji Optional[Nullable[string]] `json:"unicode_emoji,omitempty"`
	Position     int                        `json:"position"`
	Permissions  Permissions                `json:"permissions"`
	Managed      bool                       `json:"managed"`
	Mentionable  bool                       `json:"mentionable"`
	Tags         Optional[RoleTags]         `json:"tags,omitempty"`
	Flags        RoleFlags                  `json:"flags"`
}

type RoleTags struct {
	BotID                 Optional[Snowflake]          `json:"bot_id,omitempty"`
	IntegrationID         Optional[Snowflake]          `json:"integration_id,omitempty"`
	PremiumSubscriber     Optional[Nullable[struct{}]] `json:"premium_subscriber,omitempty"`
	SubscriptionListingID Optional[Snowflake]          `json:"subscription_listing_id"`
	AvailableForPurchase  Optional[Nullable[struct{}]] `json:"available_for_purchase,omitempty"`
	GuildConnections      Optional[Nullable[struct{}]] `json:"guild_connections,omitempty"`
}

type RoleFlags int

const (
	RoleFlagsInPrompt RoleFlags = 1 << 0
)

type RoleCreateEvent struct {
	Role    Role      `json:"role"`
	GuildID Snowflake `json:"guild_id"`
}

type RoleUpdateEvent struct {
	Role    Role      `json:"role"`
	GuildID Snowflake `json:"guild_id"`
}

type RoleDeleteEvent struct {
	RoleID  Snowflake `json:"role_id"`
	GuildID Snowflake `json:"guild_id"`
}

type GuildRoleListRequest struct {
	GuildID Snowflake `json:"guild_id"`
}

type GuildRoleListResponse = []Role

type RoleCreateRequest struct {
	GuildID      Snowflake             `json:"guild_id"`
	Name         Optional[string]      `json:"name,omitempty"`
	Permissions  Optional[Permissions] `json:"permissions,omitempty"`
	Color        Optional[int]         `json:"color,omitempty"`
	Hoist        Optional[bool]        `json:"hoist,omitempty"`
	Icon         Optional[string]      `json:"icon,omitempty"`
	UnicodeEmoji Optional[string]      `json:"unicode_emoji,omitempty"`
	Mentionable  Optional[bool]        `json:"mentionable,omitempty"`
}

type RoleCreateResponse = Role

type RoleModifyRequest struct {
	GuildID      Snowflake             `json:"guild_id"`
	RoleID       Snowflake             `json:"role_id"`
	Name         Optional[string]      `json:"name,omitempty"`
	Permissions  Optional[Permissions] `json:"permissions,omitempty"`
	Color        Optional[int]         `json:"color,omitempty"`
	Icon         Optional[string]      `json:"icon,omitempty"`
	UnicodeEmoji Optional[string]      `json:"unicode_emoji,omitempty"`
	Mentionable  Optional[bool]        `json:"mentionable,omitempty"`
	Hoist        Optional[bool]        `json:"hoist,omitempty"`
}

type RoleModifyResponse = Role

type RoleDeleteRequest struct {
	GuildID Snowflake `json:"guild_id"`
	RoleID  Snowflake `json:"role_id"`
}

type RoleDeleteResponse struct{}

type RolePositionsModifyRequest = []RolePositionsModifyEntry

type RolePositionsModifyEntry struct {
	ID       Snowflake     `json:"id"`
	Position Optional[int] `json:"position,omitempty"`
}

type RolePositionsModifyResponse struct{}
