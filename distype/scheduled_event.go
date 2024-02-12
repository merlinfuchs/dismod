package distype

import "time"

type ScheduledEvent struct {
	ID                 Snowflake                              `json:"id"`
	GuildID            Snowflake                              `json:"guild_id"`
	ChannelID          Nullable[Snowflake]                    `json:"channel_id"`
	CreatorID          *Nullable[Snowflake]                   `json:"creator_id,omitempty"`
	Name               string                                 `json:"name"`
	Desc               *string                                `json:"description,omitempty"`
	ScheduledStartTime time.Time                              `json:"scheduled_start_time"`
	ScheduledEndTime   Nullable[time.Time]                    `json:"scheduled_end_time"`
	PrivacyLevel       PrivacyLevel                           `json:"privacy_level"`
	Status             ScheduledEventStatus                   `json:"status"`
	EntityType         ScheduledEventEntityType               `json:"entity_type"`
	EntityID           Nullable[Snowflake]                    `json:"entity_id"`
	EntityMetadata     Nullable[ScheduledEventEntityMetadata] `json:"entity_metadata"`
	Creator            *User                                  `json:"creator,omitempty"`
	UserCount          *int                                   `json:"user_count,omitempty"`
	Image              *Nullable[string]                      `json:"image,omitempty"`
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
	Location *string `json:"location,omitempty"`
}

type ScheduledEventUser struct {
	GuildScheduledEventID Snowflake `json:"guild_scheduled_event_id"`
	User                  User      `json:"user"`
	Member                *Member   `json:"member,omitempty"`
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
	GuildID       Snowflake `json:"guild_id"`
	WithUserCount *bool     `json:"with_user_count,omitempty"`
}

type GuildScheduledEventListResponse = []ScheduledEvent

type ScheduledEventCreateRequest struct {
	GuildID            Snowflake                     `json:"guild_id"`
	ChannelID          *Snowflake                    `json:"channel_id,omitempty"`
	Name               string                        `json:"name"`
	EntityMetadata     *ScheduledEventEntityMetadata `json:"entity_metadata,omitempty"`
	PrivacyLevel       *PrivacyLevel                 `json:"privacy_level,omitempty"`
	ScheduledStartTime time.Time                     `json:"scheduled_start_time"`
	ScheduledEndTime   *time.Time                    `json:"scheduled_end_time,omitempty"`
	Description        *string                       `json:"description,omitempty"`
	EntityType         ScheduledEventEntityType      `json:"entity_type"`
	Image              *Nullable[string]             `json:"image,omitempty"`
}

type ScheduledEventCreateResponse = ScheduledEvent

type ScheduledEventModifyRequest struct {
	GuildID               Snowflake                     `json:"guild_id"`
	GuildScheduledEventID Snowflake                     `json:"guild_scheduled_event_id"`
	ChannelID             *Snowflake                    `json:"channel_id,omitempty"`
	Name                  *string                       `json:"name,omitempty"`
	EntityMetadata        *ScheduledEventEntityMetadata `json:"entity_metadata,omitempty"`
	PrivacyLevel          *PrivacyLevel                 `json:"privacy_level,omitempty"`
	ScheduledStartTime    *time.Time                    `json:"scheduled_start_time,omitempty"`
	ScheduledEndTime      *time.Time                    `json:"scheduled_end_time,omitempty"`
	Description           *string                       `json:"description,omitempty"`
	Image                 *Nullable[string]             `json:"image,omitempty"`
	Status                *ScheduledEventStatus         `json:"status,omitempty"`
	EntityType            *ScheduledEventEntityType     `json:"entity_type,omitempty"`
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
	GuildID               Snowflake  `json:"guild_id"`
	GuildScheduledEventID Snowflake  `json:"guild_scheduled_event_id"`
	Limit                 *int       `json:"limit,omitempty"`
	WithMember            *bool      `json:"with_member,omitempty"`
	Before                *Snowflake `json:"before,omitempty"`
	After                 *Snowflake `json:"after,omitempty"`
}

type ScheduledEventUserListResponse = []ScheduledEventUser
