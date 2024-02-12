package distype

import "time"

type Member struct {
	User                       *User             `json:"user,omitempty"`
	Nick                       *Nullable[string] `json:"nick,omitempty"`
	Avatar                     *Nullable[string] `json:"avatar,omitempty"`
	Roles                      []Snowflake       `json:"roles"`
	JoinedAt                   time.Time         `json:"joined_at"`
	PremiumSince               *time.Time        `json:"premium_since,omitempty"`
	Deaf                       bool              `json:"deaf"`
	Mute                       bool              `json:"mute"`
	Flags                      MemberFlags       `json:"flags"`
	Pending                    *bool             `json:"pending,omitempty"`
	Permissions                *Permissions      `json:"permissions,omitempty"`
	CommunicationDisabledUntil *time.Time        `json:"communication_disabled_until,omitempty"`
}

type MemberFlags int

const (
	MemberFlagsDidRejoin            MemberFlags = 1 << 0
	MemberFlagsCompletedOnboarding  MemberFlags = 1 << 1
	MemberFlagsBypassesVerification MemberFlags = 1 << 2
	MemberFlagsStartedOnboarding    MemberFlags = 1 << 3
)

type MemberAddEvent struct {
	Member
	GuildID Snowflake `json:"guild_id"`
}

type MemberRemoveEvent struct {
	User    User      `json:"user"`
	GuildID Snowflake `json:"guild_id"`
}

type MemberUpdateEvent = Member

type MemberChunkEvent struct {
	GuildID    Snowflake   `json:"guild_id"`
	Members    []Member    `json:"members"`
	ChunkIndex int         `json:"chunk_index"`
	ChunkCount int         `json:"chunk_count"`
	NotFound   []Snowflake `json:"not_found,omitempty"`
	Nonce      string      `json:"nonce,omitempty"`
}

type MemberGetRequest struct {
	GuildID Snowflake `json:"guild_id"`
	UserID  Snowflake `json:"user_id"`
}

type MemberGetResponse = Member

type GuildMemberListRequest struct {
	GuildID Snowflake  `json:"guild_id"`
	Limit   *int       `json:"limit,omitempty"`
	After   *Snowflake `json:"after,omitempty"`
}

type GuildMemberListResponse = []Member

type GuildMemberSearchRequest struct {
	GuildID Snowflake `json:"guild_id"`
	Query   string    `json:"query"`
	Limit   *int      `json:"limit,omitempty"`
}

type GuildMemberSearchResponse = []Member

type MemberAddRequest struct {
	GuildID     Snowflake   `json:"guild_id"`
	UserID      Snowflake   `json:"user_id"`
	AccessToken string      `json:"access_token"`
	Nick        *string     `json:"nick,omitempty"`
	Roles       []Snowflake `json:"roles,omitempty"`
	Mute        *bool       `json:"mute,omitempty"`
	Deaf        *bool       `json:"deaf,omitempty"`
}

type MemberModifyRequest struct {
	GuildID                    Snowflake    `json:"guild_id"`
	UserID                     Snowflake    `json:"user_id"`
	Nick                       *string      `json:"nick,omitempty"`
	Roles                      []Snowflake  `json:"roles,omitempty"`
	Mute                       *bool        `json:"mute,omitempty"`
	Deaf                       *bool        `json:"deaf,omitempty"`
	ChannelID                  *Snowflake   `json:"channel_id,omitempty"`
	CommunicationDisabledUntil *time.Time   `json:"communication_disabled_until,omitempty"`
	Flags                      *MemberFlags `json:"flags,omitempty"`
}

type MemberModifyResponse = Member

type MemberModifyCurrentRequest struct {
	GuildID Snowflake `json:"guild_id"`
	Nick    *string   `json:"nick,omitempty"`
}

type MemberModifyCurrentResponse struct{}

type MemberRoleAddRequest struct {
	GuildID Snowflake `json:"guild_id"`
	UserID  Snowflake `json:"user_id"`
	RoleID  Snowflake `json:"role_id"`
}

type MemberRoleAddResponse struct{}

type MemberRoleRemoveRequest struct {
	GuildID Snowflake `json:"guild_id"`
	UserID  Snowflake `json:"user_id"`
	RoleID  Snowflake `json:"role_id"`
}

type MemberRoleRemoveResponse struct{}

type MemberRemoveRequest struct {
	GuildID Snowflake `json:"guild_id"`
	UserID  Snowflake `json:"user_id"`
}

type MemberRemoveResponse struct{}

type MemberPruneCountRequest struct {
	GuildID      Snowflake   `json:"guild_id"`
	Days         int         `json:"days"`
	IncludeRoles []Snowflake `json:"include_roles,omitempty"`
}

type MemberPruneCountResponse struct {
	Pruned int `json:"pruned"`
}

type MemberPruneRequest struct {
	GuildID           Snowflake   `json:"guild_id"`
	Days              int         `json:"days"`
	ComputePruneCount *bool       `json:"compute_prune_count,omitempty"`
	IncludeRoles      []Snowflake `json:"include_roles,omitempty"`
}

type MemberPruneResponse struct {
	Pruned Nullable[int] `json:"pruned"`
}
