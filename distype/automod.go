package distype

type AutoModerationRuleObject struct {
	ID              Snowflake                     `json:"id"`
	GuildID         Snowflake                     `json:"guild_id"`
	Name            string                        `json:"name"`
	CreatorID       Snowflake                     `json:"creator_id"`
	EventType       AutoModerationEventType       `json:"event_type"`
	TriggerType     AutoModerationTriggerType     `json:"trigger_type"`
	TriggerMetadata AutoModerationTriggerMetadata `json:"trigger_metadata"`
	// TODO
}

type AutoModerationEventType int

const (
	AutoModerationEventTypeMessageSend AutoModerationEventType = 1
)

type AutoModerationTriggerType int

const (
	AutoModerationTriggerTypeKeyword       AutoModerationTriggerType = 1
	AutoModerationTriggerTypeSpam          AutoModerationTriggerType = 3
	AutoModerationTriggerTypeKeywordPreset AutoModerationTriggerType = 4
	AutoModerationTriggerTypeMentionSpam   AutoModerationTriggerType = 5
)

type AutoModerationTriggerMetadata struct{}

type AutoModerationRuleCreateEvent = AutoModerationRuleObject

type AutoModerationRuleUpdateEvent = AutoModerationRuleObject

type AutoModerationRuleDeleteEvent = AutoModerationRuleObject

type AutoModerationActionExecutionEvent struct{}
