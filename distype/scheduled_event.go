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
