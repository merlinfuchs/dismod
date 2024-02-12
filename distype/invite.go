package distype

import "time"

type Invite struct {
	Code                     string            `json:"code"`
	Guild                    *Guild            `json:"guild,omitempty"`
	Channel                  Nullable[Channel] `json:"channel"`
	Inviter                  *User             `json:"inviter,omitempty"`
	TargetType               *InviteTargetType `json:"target_type,omitempty"`
	TargetUser               *User             `json:"target_user,omitempty"`
	TargetApplication        *Application      `json:"target_application,omitempty"`
	ApproximatePresenceCount *int              `json:"approximate_presence_count,omitempty"`
	ApproximateMemberCount   *int              `json:"approximate_member_count,omitempty"`
	ExpiresAt                *Nullable[string] `json:"expires_at,omitempty"`
	StageInstance            *StageInstance    `json:"stage_instance,omitempty"`
	GuildScheduledEvent      *ScheduledEvent   `json:"guild_scheduled_event,omitempty"`

	// Extra metadata for the invite, like the number of times it's been used
	Uses      *int       `json:"uses,omitempty"`
	MaxUses   *int       `json:"max_uses,omitempty"`
	MaxAge    *int       `json:"max_age,omitempty"`
	Temporary *bool      `json:"temporary,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

type InviteTargetType int

const (
	InviteTargetTypeStream              InviteTargetType = 1
	InviteTargetTypeEmbeddedApplication InviteTargetType = 2
)

type InviteCreateEvent struct {
	ChannelID         Snowflake         `json:"channel_id"`
	Code              string            `json:"code"`
	CreatedAt         time.Time         `json:"created_at"`
	GuildID           *Snowflake        `json:"guild_id,omitempty"`
	Inviter           *User             `json:"inviter,omitempty"`
	MaxAge            int               `json:"max_age"`
	MaxUses           int               `json:"max_uses"`
	TargetType        *InviteTargetType `json:"target_type,omitempty"`
	TargetUser        *User             `json:"target_user,omitempty"`
	TargetApplication *Application      `json:"target_application,omitempty"`
	Temporary         bool              `json:"temporary"`
	Uses              int               `json:"uses"`
}

type InviteDeleteEvent struct {
	ChannelID Snowflake  `json:"channel_id"`
	GuildID   *Snowflake `json:"guild_id,omitempty"`
	Code      string     `json:"code"`
}

type InviteGetRequest struct {
	Code                  string     `json:"code"`
	WithCounts            *bool      `json:"with_counts,omitempty"`
	WithExpiration        *bool      `json:"with_expiration,omitempty"`
	GuildScheduledEventID *Snowflake `json:"guild_scheduled_event_id,omitempty"`
}

type InviteGetResponse = Invite

type InviteDeleteRequest struct {
	Code string `json:"code"`
}

type InviteDeleteResponse = Invite

type GuildInviteListRequest struct {
	GuildID Snowflake `json:"guild_id"`
}

type GuildInviteListResponse = []Invite

type ChannelInviteListRequest struct {
	ChannelID Snowflake `json:"channel_id"`
}

type ChannelInviteListResponse = []Invite

type ChannelInviteCreateRequest struct {
	ChannelID           Snowflake         `json:"channel_id"`
	MaxAge              *int              `json:"max_age,omitempty"`
	MaxUses             *int              `json:"max_uses,omitempty"`
	Temporary           *bool             `json:"temporary,omitempty"`
	Unique              *bool             `json:"unique,omitempty"`
	TargetType          *InviteTargetType `json:"target_type,omitempty"`
	TargetUserID        *Snowflake        `json:"target_user_id,omitempty"`
	TargetApplicationID *Snowflake        `json:"target_application_id,omitempty"`
}

type ChannelInviteCreateResponse = Invite
