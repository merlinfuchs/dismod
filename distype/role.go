package distype

type Role struct {
	ID           Snowflake         `json:"id"`
	Name         string            `json:"name"`
	Color        int               `json:"color"`
	Hoist        bool              `json:"hoist"`
	Icon         *Nullable[string] `json:"icon,omitempty"`
	UnicodeEmoji *Nullable[string] `json:"unicode_emoji,omitempty"`
	Position     int               `json:"position"`
	Permissions  Permissions       `json:"permissions"`
	Managed      bool              `json:"managed"`
	Mentionable  bool              `json:"mentionable"`
	Tags         *RoleTags         `json:"tags,omitempty"`
	Flags        RoleFlags         `json:"flags"`
}

type RoleTags struct {
	BotID                 *Snowflake          `json:"bot_id,omitempty"`
	IntegrationID         *Snowflake          `json:"integration_id,omitempty"`
	PremiumSubscriber     *Nullable[struct{}] `json:"premium_subscriber,omitempty"`
	SubscriptionListingID *Snowflake          `json:"subscription_listing_id,omitempty"`
	AvailableForPurchase  *Nullable[struct{}] `json:"available_for_purchase,omitempty"`
	GuildConnections      *Nullable[struct{}] `json:"guild_connections,omitempty"`
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
	GuildID      Snowflake    `json:"guild_id"`
	Name         *string      `json:"name,omitempty"`
	Permissions  *Permissions `json:"permissions,omitempty"`
	Color        *int         `json:"color,omitempty"`
	Hoist        *bool        `json:"hoist,omitempty"`
	Icon         *string      `json:"icon,omitempty"`
	UnicodeEmoji *string      `json:"unicode_emoji,omitempty"`
	Mentionable  *bool        `json:"mentionable,omitempty"`
}

type RoleCreateResponse = Role

type RoleModifyRequest struct {
	GuildID      Snowflake    `json:"guild_id"`
	RoleID       Snowflake    `json:"role_id"`
	Name         *string      `json:"name,omitempty"`
	Permissions  *Permissions `json:"permissions,omitempty"`
	Color        *int         `json:"color,omitempty"`
	Icon         *string      `json:"icon,omitempty"`
	UnicodeEmoji *string      `json:"unicode_emoji,omitempty"`
	Mentionable  *bool        `json:"mentionable,omitempty"`
	Hoist        *bool        `json:"hoist,omitempty"`
}

type RoleModifyResponse = Role

type RoleDeleteRequest struct {
	GuildID Snowflake `json:"guild_id"`
	RoleID  Snowflake `json:"role_id"`
}

type RoleDeleteResponse struct{}

type RolePositionsModifyRequest = []RolePositionsModifyEntry

type RolePositionsModifyEntry struct {
	ID       Snowflake `json:"id"`
	Position *int      `json:"position,omitempty"`
}

type RolePositionsModifyResponse struct{}
