package distype

import "time"

type Invite struct {
	Code                     string                     `json:"code"`
	Guild                    Optional[Guild]            `json:"guild,omitempty"`
	Channel                  Nullable[Channel]          `json:"channel"`
	Inviter                  Optional[User]             `json:"inviter,omitempty"`
	TargetType               Optional[InviteTargetType] `json:"target_type,omitempty"`
	TargetUser               Optional[User]             `json:"target_user,omitempty"`
	TargetApplication        Optional[Application]      `json:"target_application,omitempty"`
	ApproximatePresenceCount Optional[int]              `json:"approximate_presence_count,omitempty"`
	ApproximateMemberCount   Optional[int]              `json:"approximate_member_count,omitempty"`
	ExpiresAt                Optional[Nullable[string]] `json:"expires_at,omitempty"`
	StageInstance            Optional[StageInstance]    `json:"stage_instance,omitempty"`
	GuildScheduledEvent      Optional[ScheduledEvent]   `json:"guild_scheduled_event,omitempty"`

	// Extra metadata for the invite, like the number of times it's been used
	Uses      Optional[int]       `json:"uses"`
	MaxUses   Optional[int]       `json:"max_uses"`
	MaxAge    Optional[int]       `json:"max_age"`
	Temporary Optional[bool]      `json:"temporary"`
	CreatedAt Optional[time.Time] `json:"created_at"`
}

type InviteTargetType int

const (
	InviteTargetTypeStream              InviteTargetType = 1
	InviteTargetTypeEmbeddedApplication InviteTargetType = 2
)

type InviteCreateEvent struct {
	ChannelID         Snowflake                  `json:"channel_id"`
	Code              string                     `json:"code"`
	CreatedAt         time.Time                  `json:"created_at"`
	GuildID           Optional[Snowflake]        `json:"guild_id,omitempty"`
	Inviter           Optional[User]             `json:"inviter,omitempty"`
	MaxAge            int                        `json:"max_age"`
	MaxUses           int                        `json:"max_uses"`
	TargetType        Optional[InviteTargetType] `json:"target_type"`
	TargetUser        Optional[User]             `json:"target_user,omitempty"`
	TargetApplication Optional[Application]      `json:"target_application,omitempty"`
	Temporary         bool                       `json:"temporary"`
	Uses              int                        `json:"uses"`
}

type InviteDeleteEvent struct {
	ChannelID Snowflake           `json:"channel_id"`
	GuildID   Optional[Snowflake] `json:"guild_id,omitempty"`
	Code      string              `json:"code"`
}

type InviteGetRequest struct {
	Code                  string              `json:"code"`
	WithCounts            Optional[bool]      `json:"with_counts,omitempty"`
	WithExpiration        Optional[bool]      `json:"with_expiration,omitempty"`
	GuildScheduledEventID Optional[Snowflake] `json:"guild_scheduled_event_id,omitempty"`
}

type InviteGetResponse = Invite

type InviteDeleteRequest struct {
	Code string `json:"code"`
}

type InviteDeleteResponse = Invite

type GuildListInvitesRequest struct {
	GuildID Snowflake `json:"guild_id"`
}

type GuildListInvitesResponse = []Invite

type ChannelListInvitesRequest struct {
	ChannelID Snowflake `json:"channel_id"`
}

type ChannelListInvitesResponse = []Invite

type ChannelCreateInviteRequest struct {
	ChannelID           Snowflake                  `json:"channel_id"`
	MaxAge              Optional[int]              `json:"max_age,omitempty"`
	MaxUses             Optional[int]              `json:"max_uses,omitempty"`
	Temporary           Optional[bool]             `json:"temporary,omitempty"`
	Unique              Optional[bool]             `json:"unique,omitempty"`
	TargetType          Optional[InviteTargetType] `json:"target_type,omitempty"`
	TargetUserID        Optional[Snowflake]        `json:"target_user_id,omitempty"`
	TargetApplicationID Optional[Snowflake]        `json:"target_application_id,omitempty"`
}

type ChannelCreateInviteResponse = Invite
