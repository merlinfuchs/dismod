package distype

import (
	"fmt"
	"strconv"
)

type User struct {
	ID               Snowflake         `json:"id"`
	Username         string            `json:"username"`
	Discriminator    string            `json:"discriminator"`
	GlobalName       Nullable[string]  `json:"global_name"`
	Avatar           Nullable[string]  `json:"avatar"`
	Bot              *bool             `json:"bot,omitempty"`
	System           *bool             `json:"system,omitempty"`
	MFAEnabled       *bool             `json:"mfa_enabled,omitempty"`
	Banner           *Nullable[string] `json:"banner,omitempty"`
	AccentColor      *Nullable[int]    `json:"accent_color,omitempty"`
	Locale           *string           `json:"locale,omitempty"`
	Verified         *bool             `json:"verified,omitempty"`
	Email            *Nullable[string] `json:"email,omitempty"`
	Flags            *UserFlags        `json:"flags,omitempty"`
	PremiumType      *int              `json:"premium_type,omitempty"`
	PublicFlags      *UserFlags        `json:"public_flags,omitempty"`
	AvatarDecoration *Nullable[string] `json:"avatar_decoration,omitempty"`
}

func (u User) AvatarURL() string {
	if u.Avatar.Valid {
		return fmt.Sprintf("https://cdn.discordapp.com/avatars/%s/%s.png", u.ID, u.Avatar.Value)
	}

	if u.Discriminator == "0" {
		dis, _ := strconv.ParseUint(u.Discriminator, 10, 8)
		return fmt.Sprintf("https://cdn.discordapp.com/embed/avatars/%d.png", dis%5)
	}

	id, _ := strconv.ParseUint(u.Discriminator, 10, 64)
	return fmt.Sprintf("https://cdn.discordapp.com/embed/avatars/%d.png", (id>>22)%6)
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

type UserUpdateEvent = User

type UserGetRequest struct {
	UserID Snowflake `json:"user_id"`
}

type UserGetResponse = User
