package distype

type Guild struct {
	ID                          Snowflake                       `json:"id"`
	Name                        string                          `json:"name"`
	Icon                        Nullable[string]                `json:"icon"`
	IconHash                    Optional[Nullable[string]]      `json:"icon_hash,omitempty"`
	Splash                      Nullable[string]                `json:"splash"`
	DiscoverySplash             Nullable[string]                `json:"discovery_splash"`
	Owner                       Optional[bool]                  `json:"owner,omitempty"`
	OwnerID                     Snowflake                       `json:"owner_id"`
	Permissions                 Optional[Permissions]           `json:"permissions,omitempty"`
	Region                      Optional[string]                `json:"region,omitempty"`
	AFKChannelID                Nullable[Snowflake]             `json:"afk_channel_id"`
	AFKTimeout                  int                             `json:"afk_timeout"`
	WidgetEnabled               Optional[bool]                  `json:"widget_enabled,omitempty"`
	WidgetChannelID             Optional[Nullable[Snowflake]]   `json:"widget_channel_id,omitempty"`
	VerificationLevel           VerificationLevel               `json:"verification_level"`
	DefaultMessageNotifications DefaultMessageNotificationlevel `json:"default_message_notifications"`
	ExplicitContentFilter       ExplicitContentFilterLevel      `json:"explicit_content_filter"`
	Roles                       []Role                          `json:"roles"`
	Emojis                      []Emoji                         `json:"emojis"`
	Features                    []string                        `json:"features"`
	MFALevel                    MFALevel                        `json:"mfa_level"`
	ApplicationID               Nullable[Snowflake]             `json:"application_id"`
	SystemChannelID             Nullable[Snowflake]             `json:"system_channel_id"`
	SystemChannelFlags          SystemChannelFlags              `json:"system_channel_flags"`
	MaxPresences                Optional[Nullable[int]]         `json:"max_presences,omitempty"`
	MaxMembers                  Optional[int]                   `json:"max_members,omitempty"`
	VanityURLCode               Nullable[string]                `json:"vanity_url_code"`
	Description                 Nullable[string]                `json:"description"`
	Banner                      Nullable[string]                `json:"banner"`
	PremiumTier                 PremiumTier                     `json:"premium_tier"`
	PremiumSubscriptionCount    Optional[int]                   `json:"premium_subscription_count,omitempty"`
	PreferredLocale             string                          `json:"preferred_locale"`
	PublicUpdatesChannelID      Nullable[Snowflake]             `json:"public_updates_channel_id"`
	MaxVideoChannelUsers        Optional[int]                   `json:"max_video_channel_users,omitempty"`
	MaxStageVideoChannelUsers   Optional[int]                   `json:"max_stage_video_channel_users,omitempty"`
	ApproximateMemberCount      Optional[int]                   `json:"approximate_member_count,omitempty"`
	ApproximatePresenceCount    Optional[int]                   `json:"approximate_presence_count,omitempty"`
	WelcomeScreen               Optional[WelcomeScreen]         `json:"welcome_screen,omitempty"`
	NSFWLevel                   GuildNSFWLevel                  `json:"nsfw_level"`
	Stickers                    []Sticker                       `json:"stickers,omitempty"`
	PremiumProgressBarEnabled   bool                            `json:"premium_progress_bar_enabled"`
	SafetyAlertsChannelID       Nullable[Snowflake]             `json:"safety_alerts_channel_id"`
}

type UnavailableGuild struct {
	ID          Snowflake `json:"id"`
	Unavailable bool      `json:"unavailable"`
}

type DefaultMessageNotificationlevel int

const (
	DefaultMessageNotificationLevelAllMessages  DefaultMessageNotificationlevel = 0
	DefaultMessageNotificationLevelOnlyMentions DefaultMessageNotificationlevel = 1
)

type ExplicitContentFilterLevel int

const (
	ExplicitContentFilterLevelDisabled            ExplicitContentFilterLevel = 0
	ExplicitContentFilterLevelMembersWithoutRoles ExplicitContentFilterLevel = 1
	ExplicitContentFilterLevelAllMembers          ExplicitContentFilterLevel = 2
)

type MFALevel int

const (
	MFALevelNone     MFALevel = 0
	MFALevelElevated MFALevel = 1
)

type VerificationLevel int

const (
	VerificationLevelNone     VerificationLevel = 0
	VerificationLevelLow      VerificationLevel = 1
	VerificationLevelMedium   VerificationLevel = 2
	VerificationLevelHigh     VerificationLevel = 3
	VerificationLevelVeryHigh VerificationLevel = 4
)

type GuildNSFWLevel int

const (
	GuildNSFWLevelDefault       GuildNSFWLevel = 0
	GuildNSFWLevelExplicit      GuildNSFWLevel = 1
	GuildNSFWLevelSafe          GuildNSFWLevel = 2
	GuildNSFWLevelAgeRestricted GuildNSFWLevel = 3
)

type PremiumTier int

const (
	PremiumTierNone PremiumTier = 0
	PremiumTier1    PremiumTier = 1
	PremiumTier2    PremiumTier = 2
	PremiumTier3    PremiumTier = 3
)

type SystemChannelFlags int

const (
	SystemChannelFlagsSuppressJoinNotifications                           SystemChannelFlags = 1 << 0
	SystemChannelFlagsSuppressPremiumSubscriptions                        SystemChannelFlags = 1 << 1
	SystemChannelFlagsSuppressGuildReminderNotifications                  SystemChannelFlags = 1 << 2
	SystemChannelFlagsSuppressJoinNotificationReplies                     SystemChannelFlags = 1 << 3
	SystemChannelFlagsSuppressRoleSubscriptionPurchaseNotifications       SystemChannelFlags = 1 << 4
	SystemChannelFlagsSuppressRoleSubscriptionPurchaseNotificationReplies SystemChannelFlags = 1 << 5
)

type WelcomeScreen struct {
	Description     Nullable[string]       `json:"description"`
	WelcomeChannels []WelcomeScreenChannel `json:"welcome_channels"`
}

type WelcomeScreenChannel struct {
	ChannelID   Snowflake           `json:"channel_id"`
	Description string              `json:"description"`
	EmojiID     Nullable[Snowflake] `json:"emoji_id"`
	EmojiName   Nullable[string]    `json:"emoji_name"`
}

type GuildCreateEvent = Guild

type GuildUpdateEvent = Guild

type GuildDeleteEvent = UnavailableGuild

type GuildGetRequest struct {
	GuildID Snowflake `json:"guild_id"`
}

type GuildGetResponse = Guild

type GuildUpdateRequest struct {
	GuildID                     Snowflake                                 `json:"guild_id"`
	Name                        Optional[string]                          `json:"name,omitempty"`
	Region                      Optional[Nullable[string]]                `json:"region,omitempty"`
	VerificationLevel           Optional[VerificationLevel]               `json:"verification_level,omitempty"`
	DefaultMessageNotifications Optional[DefaultMessageNotificationlevel] `json:"default_message_notifications,omitempty"`
	ExplicitContentFilter       Optional[ExplicitContentFilterLevel]      `json:"explicit_content_filter,omitempty"`
	AFKChannelID                Optional[Nullable[Snowflake]]             `json:"afk_channel_id,omitempty"`
	AFKTimeout                  Optional[int]                             `json:"afk_timeout,omitempty"`
	Icon                        Optional[Nullable[string]]                `json:"icon,omitempty"`
	OwnerID                     Optional[Snowflake]                       `json:"owner_id,omitempty"`
	Splash                      Optional[Nullable[string]]                `json:"splash,omitempty"`
	DiscoverySplash             Optional[Nullable[string]]                `json:"discovery_splash,omitempty"`
	Banner                      Optional[Nullable[string]]                `json:"banner,omitempty"`
	SystemChannelID             Optional[Nullable[Snowflake]]             `json:"system_channel_id,omitempty"`
	SystemChannelFlags          Optional[SystemChannelFlags]              `json:"system_channel_flags,omitempty"`
	RulesChannelID              Optional[Nullable[Snowflake]]             `json:"rules_channel_id,omitempty"`
	PublicUpdatesChannelID      Optional[Nullable[Snowflake]]             `json:"public_updates_channel_id,omitempty"`
	PreferredLocale             Optional[Nullable[string]]                `json:"preferred_locale,omitempty"`
	Features                    []string                                  `json:"features,omitempty"`
	Description                 Optional[Nullable[string]]                `json:"description,omitempty"`
	PremiumProgressBarEnabled   Optional[bool]                            `json:"premium_progress_bar_enabled,omitempty"`
	SafetyAlertsChannelID       Optional[Nullable[Snowflake]]             `json:"safety_alerts_channel_id,omitempty"`
}

type GuildUpdateResponse = Guild
