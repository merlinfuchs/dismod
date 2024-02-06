package distype

import "time"

type ScheduledEvent struct {
	ID                 Snowflake                              `json:"id"`
	GuildID            Snowflake                              `json:"guild_id"`
	ChannelID          Nullable[Snowflake]                    `json:"channel_id"`
	CreatorID          Optional[Nullable[Snowflake]]          `json:"creator_id,omitempty"`
	Name               string                                 `json:"name"`
	Desc               Optional[string]                       `json:"description,omitempty"`
	ScheduledStartTime time.Time                              `json:"scheduled_start_time"`
	ScheduledEndTime   Nullable[time.Time]                    `json:"scheduled_end_time"`
	PrivacyLevel       PrivacyLevel                           `json:"privacy_level"`
	Status             ScheduledEventStatus                   `json:"status"`
	EntityType         ScheduledEventEntityType               `json:"entity_type"`
	EntityID           Nullable[Snowflake]                    `json:"entity_id"`
	EntityMetadata     Nullable[ScheduledEventEntityMetadata] `json:"entity_metadata"`
	Creator            Optional[User]                         `json:"creator,omitempty"`
	UserCount          Optional[int]                          `json:"user_count,omitempty"`
	Image              Optional[Nullable[string]]             `json:"image,omitempty"`
}

type ScheduledEventStatus int

const (
	ScheduledEventStatusScheduled ScheduledEventStatus = 1
	ScheduledEventStatusActive    ScheduledEventStatus = 2
	ScheduledEventStatusCompleted ScheduledEventStatus = 3
	ScheduledEventStatusCanceled  ScheduledEventStatus = 4
)

type ScheduledEventEntityType int

const (
	ScheduledEventEntityTypeStageInstance ScheduledEventEntityType = 1
	ScheduledEventEntityTypeVoice         ScheduledEventEntityType = 2
	ScheduledEventEntityTypeExternal      ScheduledEventEntityType = 3
)

type ScheduledEventEntityMetadata struct {
	Location Optional[string] `json:"location,omitempty"`
}

type ScheduledEventUser struct {
	GuildScheduledEventID Snowflake        `json:"guild_scheduled_event_id"`
	User                  User             `json:"user"`
	Member                Optional[Member] `json:"member,omitempty"`
}

type ScheduledEventCreateEvent = ScheduledEvent

type ScheduledEventUpdateEvent = ScheduledEvent

type ScheduledEventDeleteEvent = ScheduledEvent

type ScheduledEventUserAddEvent struct {
	GuildScheduledEventID Snowflake `json:"guild_scheduled_event_id"`
	UserID                Snowflake `json:"user_id"`
	GuildID               Snowflake `json:"guild_id"`
}

type ScheduledEventUserRemoveEvent struct {
	GuildScheduledEventID Snowflake `json:"guild_scheduled_event_id"`
	UserID                Snowflake `json:"user_id"`
	GuildID               Snowflake `json:"guild_id"`
}

type GuildScheduledEventListRequest struct {
	GuildID         Snowflake      `json:"guild_id"`
	WithUserCountrs Optional[bool] `json:"with_user_counts,omitempty"`
}

type GuildScheduledEventListResponse = []ScheduledEvent

type ScheduledEventCreateRequest struct {
	GuildID            Snowflake                              `json:"guild_id"`
	ChannelID          Optional[Snowflake]                    `json:"channel_id,omitempty"`
	Name               string                                 `json:"name"`
	EntityMetadata     Optional[ScheduledEventEntityMetadata] `json:"entity_metadata,omitempty"`
	PrivacyLevel       Optional[PrivacyLevel]                 `json:"privacy_level,omitempty"`
	ScheduledStartTime time.Time                              `json:"scheduled_start_time"`
	ScheduledEndTime   Optional[time.Time]                    `json:"scheduled_end_time,omitempty"`
	Description        Optional[string]                       `json:"description,omitempty"`
	EntityType         ScheduledEventEntityType               `json:"entity_type"`
	Image              Optional[Nullable[string]]             `json:"image,omitempty"`
}

type ScheduledEventCreateResponse = ScheduledEvent

type ScheduledEventModifyRequest struct {
	GuildID               Snowflake                              `json:"guild_id"`
	GuildScheduledEventID Snowflake                              `json:"guild_scheduled_event_id"`
	ChannelID             Optional[Snowflake]                    `json:"channel_id,omitempty"`
	Name                  Optional[string]                       `json:"name,omitempty"`
	EntityMetadata        Optional[ScheduledEventEntityMetadata] `json:"entity_metadata,omitempty"`
	PrivacyLevel          Optional[PrivacyLevel]                 `json:"privacy_level,omitempty"`
	ScheduledStartTime    Optional[time.Time]                    `json:"scheduled_start_time,omitempty"`
	ScheduledEndTime      Optional[time.Time]                    `json:"scheduled_end_time,omitempty"`
	Description           Optional[string]                       `json:"description,omitempty"`
	Image                 Optional[Nullable[string]]             `json:"image,omitempty"`
	Status                Optional[ScheduledEventStatus]         `json:"status,omitempty"`
	EntityType            Optional[ScheduledEventEntityType]     `json:"entity_type,omitempty"`
}

type ScheduledEventModifyResponse = ScheduledEvent

type ScheduledEventDeleteRequest struct {
	GuildID               Snowflake `json:"guild_id"`
	GuildScheduledEventID Snowflake `json:"guild_scheduled_event_id"`
}

type ScheduledEventDeleteResponse struct{}

type ScheduledEventGetRequest struct {
	GuildID               Snowflake `json:"guild_id"`
	GuildScheduledEventID Snowflake `json:"guild_scheduled_event_id"`
}

type ScheduledEventGetResponse = ScheduledEvent

type ScheduledEventUserListRequest struct {
	GuildID               Snowflake           `json:"guild_id"`
	GuildScheduledEventID Snowflake           `json:"guild_scheduled_event_id"`
	Limit                 Optional[int]       `json:"limit,omitempty"`
	WithMember            Optional[bool]      `json:"with_member,omitempty"`
	Before                Optional[Snowflake] `json:"before,omitempty"`
	After                 Optional[Snowflake] `json:"after,omitempty"`
}

type ScheduledEventUserListResponse = []ScheduledEventUser
