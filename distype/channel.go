package distype

import "time"

type Channel struct {
	ID                            Snowflake                  `json:"id"`
	Type                          ChannelType                `json:"type"`
	GuildID                       Optional[Snowflake]        `json:"guild_id,omitempty"`
	Position                      Optional[int]              `json:"position,omitempty"`
	PermissionOverwrites          []PermissionOverwrite      `json:"permission_overwrites,omitempty"`
	Name                          Optional[string]           `json:"name,omitempty"`
	Topic                         Optional[string]           `json:"topic,omitempty"`
	NSFW                          Optional[bool]             `json:"nsfw,omitempty"`
	LastMessageID                 Optional[Snowflake]        `json:"last_message_id,omitempty"`
	Bitrate                       Optional[int]              `json:"bitrate,omitempty"`
	UserLimit                     Optional[int]              `json:"user_limit,omitempty"`
	RateLimitPerUser              Optional[int]              `json:"rate_limit_per_user,omitempty"`
	Recipients                    []User                     `json:"recipients,omitempty"`
	Icon                          Optional[Nullable[string]] `json:"icon,omitempty"`
	OwnerID                       Optional[Snowflake]        `json:"owner_id,omitempty"`
	ApplicationID                 Optional[Snowflake]        `json:"application_id,omitempty"`
	Managed                       Optional[bool]             `json:"managed,omitempty"`
	ParentID                      Optional[Snowflake]        `json:"parent_id,omitempty"`
	LastPinTimestamp              Optional[time.Time]        `json:"last_pin_timestamp,omitempty"`
	RTCRegion                     Optional[string]           `json:"rtc_region,omitempty"`
	VideoQualityMode              Optional[VideoQualityMode] `json:"video_quality_mode,omitempty"`
	MessageCount                  Optional[int]              `json:"message_count,omitempty"`
	MemberCount                   Optional[int]              `json:"member_count,omitempty"`
	ThreadMetadata                Optional[ThreadMetadata]   `json:"thread_metadata,omitempty"`
	Member                        Optional[ThreadMember]     `json:"member,omitempty"`
	DefaultAutoArchive            Optional[int]              `json:"default_auto_archive_duration,omitempty"`
	Permissions                   Optional[Permissions]      `json:"permissions,omitempty"`
	Flags                         Optional[ChannelFlags]     `json:"flags,omitempty"`
	TotalMessagesSent             Optional[int]              `json:"total_message_sent,omitempty"`
	AvailableTags                 []ForumTag                 `json:"available_tags,omitempty"`
	AppliedTags                   []Snowflake                `json:"applied_tags,omitempty"`
	DefaultReactionEmoji          Optional[DefaultReaction]  `json:"default_reaction_emoji,omitempty"`
	DefaultThreadRateLimitPerUser Optional[int]              `json:"default_thread_rate_limit_per_user,omitempty"`
	DefaultSortOrder              Optional[SortOrderType]    `json:"default_sort_order,omitempty"`
	DefaultForumLayout            Optional[ForumLayoutType]  `json:"default_forum_layout,omitempty"`
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
	Archived            bool                `json:"archived"`
	AutoArchiveDuration Optional[int]       `json:"auto_archive_duration,omitempty"`
	ArchiveTimestamp    Optional[time.Time] `json:"archive_timestamp,omitempty"`
	Locked              bool                `json:"locked"`
	Invitable           Optional[bool]      `json:"invitable,omitempty"`
	CreateTimestamp     Optional[time.Time] `json:"create_timestamp,omitempty"`
}

type ThreadMember struct {
	ID            Optional[Snowflake] `json:"id,omitempty"`
	UserID        Optional[Snowflake] `json:"user_id,omitempty"`
	GuildID       Optional[Snowflake] `json:"guild_id,omitempty"`
	JoinTimestamp time.Time           `json:"join_timestamp"`
	Flags         int                 `json:"flags"`
	Member        Optional[Member]    `json:"member,omitempty"`
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
	GuildID          Optional[Snowflake] `json:"guild_id,omitempty"`
	ChannelID        Snowflake           `json:"channel_id"`
	LastPinTimestamp Optional[time.Time] `json:"last_pin_timestamp,omitempty"`
}

type ChannelGetRequest struct {
	ChannelID Snowflake `json:"channel_id"`
}

type ChannelGetResponse = Channel

type ChannelModifyRequest struct{} // TODO

type ChannelModifyResponse = Channel

type ChannelDeleteRequest struct {
	ChannelID Snowflake `json:"channel_id"`
}

type ChannelDeleteResponse = Channel

type ChannelEditPermissionsRequest struct {
	ChannelID   Snowflake                       `json:"channel_id"`
	OverwriteID Snowflake                       `json:"overwrite_id"`
	Allow       Optional[Nullable[Permissions]] `json:"allow,omitempty"`
	Deny        Optional[Nullable[Permissions]] `json:"deny,omitempty"`
	Type        PermissionOverwriteType         `json:"type,omitempty"`
}

type ChannelEditPermissionsResponse struct{}

type ChannelDeletePermissionsRequest struct {
	ChannelID   Snowflake `json:"channel_id"`
	OverwriteID Snowflake `json:"overwrite_id"`
}

type ChannelDeletePermissionsResponse struct{}

type ThreadStartFromMessageRequest struct {
	ChannelID           Snowflake           `json:"channel_id"`
	MessageID           Snowflake           `json:"message_id"`
	Name                string              `json:"name"`
	AutoArchiveDuration Optional[time.Time] `json:"auto_archive_duration,omitempty"`
	RateLimitPerUser    Optional[int]       `json:"rate_limit"`
}

type ThreadStartFromMessageResponse = Channel

type ThreadStartWithoutMessageRequest struct {
	ChannelID           Snowflake             `json:"channel_id"`
	Name                string                `json:"name"`
	AutoArchiveDuration Optional[time.Time]   `json:"auto_archive_duration,omitempty"`
	Type                Optional[ChannelType] `json:"type,omitempty"`
	Invitable           Optional[bool]        `json:"invitable,omitempty"`
	RateLimitPerUser    Optional[int]         `json:"rate_limit"`
}

type ThreadStartWithoutMessageResponse = Channel

type ThreadStartInForumRequest struct {
	ChannelID           Snowflake                     `json:"channel_id"`
	Name                string                        `json:"name"`
	AutoArchiveDuration Optional[time.Time]           `json:"auto_archive_duration,omitempty"`
	RateLimitPerUser    Optional[int]                 `json:"rate_limit"`
	Message             Optional[MessageCreateParams] `json:"message,omitempty"`
	AppliedTags         []Snowflake                   `json:"applied_tags,omitempty"`
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
	ChannelID  Snowflake           `json:"channel_id"`
	WithMember Optional[bool]      `json:"with_member,omitempty"`
	After      Optional[Snowflake] `json:"after,omitempty"`
	Limit      Optional[int]       `json:"limit,omitempty"`
}

type ThreadMemberListResponse = []ThreadMember

type ThreadListPublicArchivedRequest struct {
	ChannelID Snowflake           `json:"channel_id"`
	Before    Optional[time.Time] `json:"before,omitempty"`
	Limit     Optional[int]       `json:"limit,omitempty"`
}

type ThreadListPublicArchivedResponse struct {
	Threads []Channel      `json:"threads"`
	Members []ThreadMember `json:"members"`
	HasMore bool           `json:"has_more"`
}

type ThreadListPrivateArchivedRequest struct {
	ChannelID Snowflake           `json:"channel_id"`
	Before    Optional[time.Time] `json:"before,omitempty"`
	Limit     Optional[int]       `json:"limit,omitempty"`
}

type ThreadListPrivateArchivedResponse struct {
	Threads []Channel      `json:"threads"`
	Members []ThreadMember `json:"members"`
	HasMore bool           `json:"has_more"`
}

type ThreadListJoinedPrivateArchivedRequest struct {
	ChannelID Snowflake           `json:"channel_id"`
	Before    Optional[time.Time] `json:"before,omitempty"`
	Limit     Optional[int]       `json:"limit,omitempty"`
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

type GuildChannelCreateRequest struct{} // TODO

type GuildChannelCreateResponse = Channel

type GuildChannelModifyPositionsRequest = []GuildChannelModifyPositionsEntry

type GuildChannelModifyPositionsEntry struct {
	ID              Snowflake           `json:"id"`
	Position        Optional[int]       `json:"position"`
	LockPermissions Optional[bool]      `json:"lock_permissions"`
	ParentID        Optional[Snowflake] `json:"parent_id"`
}

type GuildChannelModifyPositionsResponse struct{}

type GuildThreadListActiveRequest struct {
	GuildID Snowflake `json:"guild_id"`
}

type GuildThreadListActiveResponse struct {
	Threads []Channel      `json:"threads"`
	Members []ThreadMember `json:"members"`
}
