package distype

import "time"

type Member struct {
	User                       Optional[User]             `json:"user,omitempty"`
	Nick                       Optional[Nullable[string]] `json:"nick,omitempty"`
	Avatar                     Optional[Nullable[string]] `json:"avatar,omitempty"`
	Roles                      []Snowflake                `json:"roles"`
	JoinedAt                   time.Time                  `json:"joined_at"`
	PremiumSince               Optional[time.Time]        `json:"premium_since,omitempty"`
	Deaf                       bool                       `json:"deaf"`
	Mute                       bool                       `json:"mute"`
	Flags                      MemberFlags                `json:"flags"`
	Pending                    Optional[bool]             `json:"pending,omitempty"`
	Permissions                Optional[Permissions]      `json:"permissions,omitempty"`
	CommunicationDisabledUntil Optional[time.Time]        `json:"communication_disabled_until,omitempty"`
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
	GuildID Snowflake           `json:"guild_id"`
	Limit   Optional[int]       `json:"limit,omitempty"`
	After   Optional[Snowflake] `json:"after,omitempty"`
}

type GuildMemberListResponse = []Member

type GuildMemberSearchRequest struct {
	GuildID Snowflake     `json:"guild_id"`
	Query   string        `json:"query"`
	Limit   Optional[int] `json:"limit,omitempty"`
}

type GuildMemberSearchResponse = []Member

type MemberAddRequest struct {
	GuildID     Snowflake        `json:"guild_id"`
	UserID      Snowflake        `json:"user_id"`
	AccessToken string           `json:"access_token"`
	Nick        Optional[string] `json:"nick,omitempty"`
	Roles       []Snowflake      `json:"roles,omitempty"`
	Mute        Optional[bool]   `json:"mute,omitempty"`
	Deaf        Optional[bool]   `json:"deaf,omitempty"`
}

type MemberModifyRequest struct {
	GuildID                    Snowflake             `json:"guild_id"`
	UserID                     Snowflake             `json:"user_id"`
	Nick                       Optional[string]      `json:"nick,omitempty"`
	Roles                      []Snowflake           `json:"roles,omitempty"`
	Mute                       Optional[bool]        `json:"mute,omitempty"`
	Deaf                       Optional[bool]        `json:"deaf,omitempty"`
	ChannelID                  Optional[Snowflake]   `json:"channel_id,omitempty"`
	CommunicationDisabledUntil Optional[time.Time]   `json:"communication_disabled_until,omitempty"`
	Flags                      Optional[MemberFlags] `json:"flags,omitempty"`
}

type MemberModifyResponse = Member

type MemberModifyCurrentRequest struct {
	GuildID Snowflake        `json:"guild_id"`
	Nick    Optional[string] `json:"nick,omitempty"`
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
