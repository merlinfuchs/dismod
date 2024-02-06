package distype

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
}

type InviteTargetType int

const (
	InviteTargetTypeStream              InviteTargetType = 1
	InviteTargetTypeEmbeddedApplication InviteTargetType = 2
)
