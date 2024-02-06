package distype

type User struct {
	ID               Snowflake                  `json:"id"`
	Username         string                     `json:"username"`
	Discriminator    string                     `json:"discriminator"`
	GlobalName       Nullable[string]           `json:"global_name"`
	Avatar           Nullable[string]           `json:"avatar"`
	Bot              Optional[bool]             `json:"bot,omitempty"`
	System           Optional[bool]             `json:"system,omitempty"`
	MFAEnabled       Optional[bool]             `json:"mfa_enabled,omitempty"`
	Banner           Optional[Nullable[string]] `json:"banner,omitempty"`
	AccentColor      Optional[Nullable[int]]    `json:"accent_color,omitempty"`
	Locale           Optional[string]           `json:"locale,omitempty"`
	Verified         Optional[bool]             `json:"verified,omitempty"`
	Email            Optional[Nullable[string]] `json:"email,omitempty"`
	Flags            Optional[int]              `json:"flags,omitempty"`
	PremiumType      Optional[int]              `json:"premium_type,omitempty"`
	PublicFlags      Optional[int]              `json:"public_flags,omitempty"`
	AvatarDecoration Optional[Nullable[string]] `json:"avatar_decoration,omitempty"`
}

type UserFlags int

const (
	UserFlagsStaff                 UserFlags = 1 << 0
	UserFlagsPartner               UserFlags = 1 << 1
	UserFlagsHypeSquad             UserFlags = 1 << 2
	UserFlagsBugHunterLevel1       UserFlags = 1 << 3
	UserFlagsHypeSquadOnlineHouse1 UserFlags = 1 << 6
	UserFlagsHypeSquadOnlineHouse2 UserFlags = 1 << 7
	UserFlagsHypeSquadOnlineHouse3 UserFlags = 1 << 8
	UserFlagsPremiumEarlySupporter UserFlags = 1 << 9
	UserFlagsTeamPseudoUser        UserFlags = 1 << 10
	UserFlagsBugHunterLevel2       UserFlags = 1 << 14
	UserFlagsVerifiedBot           UserFlags = 1 << 16
	UserFlagsVerifiedDeveloper     UserFlags = 1 << 17
	UserFlagsCertifiedModerator    UserFlags = 1 << 18
	UserFlagsBotHTTPInteractions   UserFlags = 1 << 19
	UserFlagsActiveDeveloper       UserFlags = 1 << 22
)

type PremiumType int

const (
	PremiumTypeNone         PremiumType = 0
	PremiumTypeNitroClassic PremiumType = 1
	PremiumTypeNitro        PremiumType = 2
	PremiumTypeNitroBasic   PremiumType = 3
)
