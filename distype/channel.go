package distype

import "time"

type Channel struct {
	ID                            Snowflake             `json:"id"`
	Type                          ChannelType           `json:"type"`
	GuildID                       *Snowflake            `json:"guild_id,omitempty"`
	Position                      *int                  `json:"position,omitempty"`
	PermissionOverwrites          []PermissionOverwrite `json:"permission_overwrites,omitempty"`
	Name                          *string               `json:"name,omitempty"`
	Topic                         *string               `json:"topic,omitempty"`
	NSFW                          *bool                 `json:"nsfw,omitempty"`
	LastMessageID                 *Snowflake            `json:"last_message_id,omitempty"`
	Bitrate                       *int                  `json:"bitrate,omitempty"`
	UserLimit                     *int                  `json:"user_limit,omitempty"`
	RateLimitPerUser              *int                  `json:"rate_limit_per_user,omitempty"`
	Recipients                    []User                `json:"recipients,omitempty"`
	Icon                          *Nullable[string]     `json:"icon,omitempty"`
	OwnerID                       *Snowflake            `json:"owner_id,omitempty"`
	ApplicationID                 *Snowflake            `json:"application_id,omitempty"`
	Managed                       *bool                 `json:"managed,omitempty"`
	ParentID                      *Snowflake            `json:"parent_id,omitempty"`
	LastPinTimestamp              *time.Time            `json:"last_pin_timestamp,omitempty"`
	RTCRegion                     *string               `json:"rtc_region,omitempty"`
	VideoQualityMode              *VideoQualityMode     `json:"video_quality_mode,omitempty"`
	MessageCount                  *int                  `json:"message_count,omitempty"`
	MemberCount                   *int                  `json:"member_count,omitempty"`
	ThreadMetadata                *ThreadMetadata       `json:"thread_metadata,omitempty"`
	Member                        *ThreadMember         `json:"member,omitempty"`
	DefaultAutoArchiveDuration    *int                  `json:"default_auto_archive_duration,omitempty"`
	Permissions                   *Permissions          `json:"permissions,omitempty"`
	Flags                         *ChannelFlags         `json:"flags,omitempty"`
	TotalMessagesSent             *int                  `json:"total_message_sent,omitempty"`
	AvailableTags                 []ForumTag            `json:"available_tags,omitempty"`
	AppliedTags                   []Snowflake           `json:"applied_tags,omitempty"`
	DefaultReactionEmoji          *DefaultReaction      `json:"default_reaction_emoji,omitempty"`
	DefaultThreadRateLimitPerUser *int                  `json:"default_thread_rate_limit_per_user,omitempty"`
	DefaultSortOrder              *SortOrderType        `json:"default_sort_order,omitempty"`
	DefaultForumLayout            *ForumLayoutType      `json:"default_forum_layout,omitempty"`
}

type ChannelType int

const (
	ChannelTypeGuildText          ChannelType = 0
	ChannelTypeDM                 ChannelType = 1
	ChannelTypeGuildVoice         ChannelType = 2
	ChannelTypeGroupDM            ChannelType = 3
	ChannelTypeGuildCategory      ChannelType = 4
	ChannelTypeGuildAnnouncement  ChannelType = 5
	ChannelTypeAnnouncementThread ChannelType = 10
	ChannelTypePublicThread       ChannelType = 11
	ChannelTypePrivateThread      ChannelType = 12
	ChannelTypeGuildStageVoice    ChannelType = 13
	ChannelTypeGuildDirectory     ChannelType = 14
	ChannelTypeGuildForum         ChannelType = 15
	ChannelTypeGuildMedia         ChannelType = 16
)

type VideoQualityMode int

const (
	VideoQualityModeAuto VideoQualityMode = 1
	VideoQualityModeFull VideoQualityMode = 2
)

type ChannelFlags int

const (
	ChannelFlagsPinned                   ChannelFlags = 1 << 1
	ChannelFlagsRequireTag               ChannelFlags = 1 << 4
	ChannelFlagsHideMediaDownloadOptions ChannelFlags = 1 << 15
)

type SortOrderType int

const (
	SortOrderTypeLatestActivity SortOrderType = 0
	SortOrderTypeCreationDate   SortOrderType = 1
)

type ForumLayoutType int

const (
	ForumLayoutTypeNotSet      ForumLayoutType = 0
	ForumLayoutTypeListView    ForumLayoutType = 1
	ForumLayoutTypeGalleryView ForumLayoutType = 2
)

type ForumTag struct {
	ID        Snowflake           `json:"id"`
	Name      string              `json:"name"`
	Moderated bool                `json:"moderated"`
	EmojiID   Nullable[Snowflake] `json:"emoji_id"`
	EmojiName Nullable[string]    `json:"emoji_name"`
}

type ThreadMetadata struct {
	Archived            bool       `json:"archived"`
	AutoArchiveDuration *int       `json:"auto_archive_duration,omitempty"`
	ArchiveTimestamp    *time.Time `json:"archive_timestamp,omitempty"`
	Locked              bool       `json:"locked"`
	Invitable           *bool      `json:"invitable,omitempty"`
	CreateTimestamp     *time.Time `json:"create_timestamp,omitempty"`
}

type ThreadMember struct {
	ID            *Snowflake `json:"id,omitempty"`
	UserID        *Snowflake `json:"user_id,omitempty"`
	GuildID       *Snowflake `json:"guild_id,omitempty"`
	JoinTimestamp time.Time  `json:"join_timestamp"`
	Flags         int        `json:"flags"`
	Member        *Member    `json:"member,omitempty"`
}

type DefaultReaction struct {
	EmojiID   Nullable[Snowflake] `json:"emoji_id"`
	EmojiName Nullable[string]    `json:"emoji_name"`
}

type PermissionOverwrite struct {
	ID    Snowflake               `json:"id"`
	Type  PermissionOverwriteType `json:"type"`
	Allow Permissions             `json:"allow"`
	Deny  Permissions             `json:"deny"`
}

type PermissionOverwriteType int

const (
	PermissionOverwriteTypeRole   PermissionOverwriteType = 0
	PermissionOverwriteTypeMember PermissionOverwriteType = 1
)

type ChannelCreateEvent = Channel

type ChannelUpdateEvent = Channel

type ChannelDeleteEvent = Channel

type ThreadCreateEvent = Channel

type ThreadUpdateEvent = Channel

type ThreadDeleteEvent = Channel

type ThreadListSyncEvent struct {
	GuildID    Snowflake      `json:"guild_id"`
	ChannelIDs []Snowflake    `json:"channel_ids,omitempty"`
	Threads    []Channel      `json:"threads"`
	Members    []ThreadMember `json:"members"`
}

type ThreadMemberUpdateEvent = ThreadMember

type ThreadMembersUpdateEvent struct {
	ID               Snowflake      `json:"id"`
	GuildID          Snowflake      `json:"guild_id"`
	MemberCount      int            `json:"member_count"`
	AddedMembers     []ThreadMember `json:"added_members,omitempty"`
	RemovedMemberIDs []Snowflake    `json:"removed_member_ids,omitempty"`
}

type ChannelPinsUpdateEvent struct {
	GuildID          *Snowflake `json:"guild_id,omitempty"`
	ChannelID        Snowflake  `json:"channel_id"`
	LastPinTimestamp *time.Time `json:"last_pin_timestamp,omitempty"`
}

type ChannelGetRequest struct {
	ChannelID Snowflake `json:"channel_id"`
}

type ChannelGetResponse = Channel

type ChannelModifyRequest struct {
	ChannelID                     Snowflake             `json:"channel_id"`
	Name                          *string               `json:"name,omitempty"`
	Type                          *ChannelType          `json:"type,omitempty"`
	Position                      *int                  `json:"position,omitempty"`
	Topic                         *string               `json:"topic,omitempty"`
	NSFW                          *bool                 `json:"nsfw,omitempty"`
	RateLimitPerUser              *int                  `json:"rate_limit_per_user,omitempty"`
	Bitrate                       *int                  `json:"bitrate,omitempty"`
	UserLimit                     *int                  `json:"user_limit,omitempty"`
	PermissionOverwrites          []PermissionOverwrite `json:"permission_overwrites,omitempty"`
	ParentID                      *Snowflake            `json:"parent_id,omitempty"`
	RTCRegion                     *string               `json:"rtc_region,omitempty"`
	VideoQualityMode              *VideoQualityMode     `json:"video_quality_mode,omitempty"`
	DefaultAutoArchiveDuration    *int                  `json:"default_auto_archive_duration,omitempty"`
	Flags                         *ChannelFlags         `json:"flags,omitempty"`
	AvailableTags                 []ForumTag            `json:"available_tags,omitempty"`
	DefaultReactionEmoji          *DefaultReaction      `json:"default_reaction_emoji,omitempty"`
	DefaultThreadRateLimitPerUser *int                  `json:"default_thread_rate_limit_per_user,omitempty"`
	DefaultSortOrder              *SortOrderType        `json:"default_sort_order,omitempty"`
	DefaultForumLayout            *ForumLayoutType      `json:"default_forum_layout,omitempty"`
}

type ChannelModifyResponse = Channel

type ChannelDeleteRequest struct {
	ChannelID Snowflake `json:"channel_id"`
}

type ChannelDeleteResponse = Channel

type ChannelEditPermissionsRequest struct {
	ChannelID   Snowflake               `json:"channel_id"`
	OverwriteID Snowflake               `json:"overwrite_id"`
	Allow       *Nullable[Permissions]  `json:"allow,omitempty"`
	Deny        *Nullable[Permissions]  `json:"deny,omitempty"`
	Type        PermissionOverwriteType `json:"type,omitempty"`
}

type ChannelEditPermissionsResponse struct{}

type ChannelDeletePermissionsRequest struct {
	ChannelID   Snowflake `json:"channel_id"`
	OverwriteID Snowflake `json:"overwrite_id"`
}

type ChannelDeletePermissionsResponse struct{}

type ThreadStartFromMessageRequest struct {
	ChannelID           Snowflake  `json:"channel_id"`
	MessageID           Snowflake  `json:"message_id"`
	Name                string     `json:"name"`
	AutoArchiveDuration *time.Time `json:"auto_archive_duration,omitempty"`
	RateLimitPerUser    *int       `json:"rate_limit"`
}

type ThreadStartFromMessageResponse = Channel

type ThreadStartWithoutMessageRequest struct {
	ChannelID           Snowflake    `json:"channel_id"`
	Name                string       `json:"name"`
	AutoArchiveDuration *time.Time   `json:"auto_archive_duration,omitempty"`
	Type                *ChannelType `json:"type,omitempty"`
	Invitable           *bool        `json:"invitable,omitempty"`
	RateLimitPerUser    *int         `json:"rate_limit"`
}

type ThreadStartWithoutMessageResponse = Channel

type ThreadStartInForumRequest struct {
	ChannelID           Snowflake            `json:"channel_id"`
	Name                string               `json:"name"`
	AutoArchiveDuration *time.Time           `json:"auto_archive_duration,omitempty"`
	RateLimitPerUser    *int                 `json:"rate_limit"`
	Message             *MessageCreateParams `json:"message,omitempty"`
	AppliedTags         []Snowflake          `json:"applied_tags,omitempty"`
}

type ThreadStartInForumResponse = Channel

type ThreadJoinRequest struct {
	ChannelID Snowflake `json:"channel_id"`
}

type ThreadJoinResponse struct{}

type ThreadMemberAddRequest struct {
	ChannelID Snowflake `json:"channel_id"`
	UserID    Snowflake `json:"user_id"`
}

type ThreadMemberAddResponse struct{}

type ThreadLeaveRequest struct {
	ChannelID Snowflake `json:"channel_id"`
}

type ThreadLeaveResponse struct{}

type ThreadMemberRemoveRequest struct {
	ChannelID Snowflake `json:"channel_id"`
	UserID    Snowflake `json:"user_id"`
}

type ThreadMemberRemoveResponse struct{}

type ThreadMemberGetRequest struct {
	ChannelID Snowflake `json:"channel_id"`
	UserID    Snowflake `json:"user_id"`
}

type ThreadMemberGetResponse = ThreadMember

type ThreadMemberListRequest struct {
	ChannelID  Snowflake  `json:"channel_id"`
	WithMember *bool      `json:"with_member,omitempty"`
	After      *Snowflake `json:"after,omitempty"`
	Limit      *int       `json:"limit,omitempty"`
}

type ThreadMemberListResponse = []ThreadMember

type ThreadListPublicArchivedRequest struct {
	ChannelID Snowflake  `json:"channel_id"`
	Before    *time.Time `json:"before,omitempty"`
	Limit     *int       `json:"limit,omitempty"`
}

type ThreadListPublicArchivedResponse struct {
	Threads []Channel      `json:"threads"`
	Members []ThreadMember `json:"members"`
	HasMore bool           `json:"has_more"`
}

type ThreadListPrivateArchivedRequest struct {
	ChannelID Snowflake  `json:"channel_id"`
	Before    *time.Time `json:"before,omitempty"`
	Limit     *int       `json:"limit,omitempty"`
}

type ThreadListPrivateArchivedResponse struct {
	Threads []Channel      `json:"threads"`
	Members []ThreadMember `json:"members"`
	HasMore bool           `json:"has_more"`
}

type ThreadListJoinedPrivateArchivedRequest struct {
	ChannelID Snowflake  `json:"channel_id"`
	Before    *time.Time `json:"before,omitempty"`
	Limit     *int       `json:"limit,omitempty"`
}

type ThreadListJoinedPrivateArchivedResponse struct {
	Threads []Channel      `json:"threads"`
	Members []ThreadMember `json:"members"`
	HasMore bool           `json:"has_more"`
}

type GuildChannelListRequest struct {
	GuildID Snowflake `json:"guild_id"`
}

type GuildChannelListResponse = []Channel

type GuildChannelCreateRequest struct {
	GuildID                       Snowflake             `json:"guild_id"`
	Name                          string                `json:"name"`
	Type                          ChannelType           `json:"type"`
	Topic                         *string               `json:"topic,omitempty"`
	Bitrate                       *int                  `json:"bitrate,omitempty"`
	UserLimit                     *int                  `json:"user_limit,omitempty"`
	RateLimitPerUser              *int                  `json:"rate_limit_per_user,omitempty"`
	Position                      *int                  `json:"position,omitempty"`
	PermissionOverwrites          []PermissionOverwrite `json:"permission_overwrites,omitempty"`
	ParentID                      *Snowflake            `json:"parent_id,omitempty"`
	NSFW                          *bool                 `json:"nsfw,omitempty"`
	RTCRegion                     *string               `json:"rtc_region,omitempty"`
	VideoQualityMode              *VideoQualityMode     `json:"video_quality_mode,omitempty"`
	DefaultAutoArchiveDuration    *int                  `json:"default_auto_archive_duration,omitempty"`
	DefaultReactionEmoji          *DefaultReaction      `json:"default_reaction_emoji,omitempty"`
	AvailableTags                 []ForumTag            `json:"available_tags,omitempty"`
	DefaultSortOrder              *SortOrderType        `json:"default_sort_order,omitempty"`
	DefaultForumLayout            *ForumLayoutType      `json:"default_forum_layout,omitempty"`
	DefaultThreadRateLimitPerUser *int                  `json:"default_thread_rate_limit_per_user,omitempty"`
}

type GuildChannelCreateResponse = Channel

type GuildChannelModifyPositionsRequest struct {
	GuildID Snowflake                          `json:"guild_id"`
	Entries []GuildChannelModifyPositionsEntry `json:"entries"`
}

type GuildChannelModifyPositionsEntry struct {
	ID              Snowflake  `json:"id"`
	Position        *int       `json:"position,omitempty"`
	LockPermissions *bool      `json:"lock_permissions,omitempty"`
	ParentID        *Snowflake `json:"parent_id,omitempty"`
}

type GuildChannelModifyPositionsResponse struct{}

type GuildThreadListActiveRequest struct {
	GuildID Snowflake `json:"guild_id"`
}

type GuildThreadListActiveResponse struct {
	Threads []Channel      `json:"threads"`
	Members []ThreadMember `json:"members"`
}
