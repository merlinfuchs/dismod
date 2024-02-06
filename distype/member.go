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
	Permissions                Optional[string]           `json:"permissions,omitempty"`
	CommunicationDisabledUntil Optional[time.Time]        `json:"communication_disabled_until,omitempty"`
}

type MemberFlags int

const (
	MemberFlagsDidRejoin            MemberFlags = 1 << 0
	MemberFlagsCompletedOnboarding  MemberFlags = 1 << 1
	MemberFlagsBypassesVerification MemberFlags = 1 << 2
	MemberFlagsStartedOnboarding    MemberFlags = 1 << 3
)
