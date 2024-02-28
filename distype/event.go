package distype

import (
	"fmt"
)

type EventType string

const (
	EventTypeAll                                 EventType = "*"
	EventTypeDisconnect                          EventType = "__DISCONNECT__"
	EventTypeConnect                             EventType = "__CONNECT__"
	EventTypeHello                               EventType = "HELLO"
	EventTypeReady                               EventType = "READY"
	EventTypeResumed                             EventType = "RESUMED"
	EventTypeReconnect                           EventType = "RECONNECT"
	EventTypeInvalidSession                      EventType = "INVALID_SESSION"
	EventTypeApplicationCommandPermissionsUpdate EventType = "APPLICATION_COMMAND_PERMISSIONS_UPDATE"
	EventTypeAutoModerationRuleCreate            EventType = "AUTO_MODERATION_RULE_CREATE"
	EventTypeAutoModerationRuleUpdate            EventType = "AUTO_MODERATION_RULE_UPDATE"
	EventTypeAutoModerationRuleDelete            EventType = "AUTO_MODERATION_RULE_DELETE"
	EventTypeAutoModerationActionExecution       EventType = "AUTO_MODERATION_ACTION_EXECUTION"
	EventTypeChannelCreate                       EventType = "CHANNEL_CREATE"
	EventTypeChannelUpdate                       EventType = "CHANNEL_UPDATE"
	EventTypeChannelDelete                       EventType = "CHANNEL_DELETE"
	EventTypeChannelPinsUpdate                   EventType = "CHANNEL_PINS_UPDATE"
	EventTypeThreadCreate                        EventType = "THREAD_CREATE"
	EventTypeThreadUpdate                        EventType = "THREAD_UPDATE"
	EventTypeThreadDelete                        EventType = "THREAD_DELETE"
	EventTypeThreadListSync                      EventType = "THREAD_LIST_SYNC"
	EventTypeThreadMemberUpdate                  EventType = "THREAD_MEMBER_UPDATE"
	EventTypeThreadMembersUpdate                 EventType = "THREAD_MEMBERS_UPDATE"
	EventTypeEntitlementCreate                   EventType = "ENTITLEMENT_CREATE"
	EventTypeEntitlementUpdate                   EventType = "ENTITLEMENT_UPDATE"
	EventTypeEntitlementDelete                   EventType = "ENTITLEMENT_DELETE"
	EventTypeGuildCreate                         EventType = "GUILD_CREATE"
	EventTypeGuildUpdate                         EventType = "GUILD_UPDATE"
	EventTypeGuildDelete                         EventType = "GUILD_DELETE"
	EventTypeGuildAuditLogEntryCreate            EventType = "GUILD_AUDIT_LOG_ENTRY_CREATE"
	EventTypeGuildBanAdd                         EventType = "GUILD_BAN_ADD"
	EventTypeGuildBanRemove                      EventType = "GUILD_BAN_REMOVE"
	EventTypeGuildEmojisUpdate                   EventType = "GUILD_EMOJIS_UPDATE"
	EventTypeGuildStickersUpdate                 EventType = "GUILD_STICKERS_UPDATE"
	EventTypeGuildIntegrationsUpdate             EventType = "GUILD_INTEGRATIONS_UPDATE"
	EventTypeGuildMemberAdd                      EventType = "GUILD_MEMBER_ADD"
	EventTypeGuildMemberRemove                   EventType = "GUILD_MEMBER_REMOVE"
	EventTypeGuildMemberUpdate                   EventType = "GUILD_MEMBER_UPDATE"
	EventTypeGuildMembersChunk                   EventType = "GUILD_MEMBERS_CHUNK"
	EventTypeGuildRoleCreate                     EventType = "GUILD_ROLE_CREATE"
	EventTypeGuildRoleUpdate                     EventType = "GUILD_ROLE_UPDATE"
	EventTypeGuildRoleDelete                     EventType = "GUILD_ROLE_DELETE"
	EventTypeGuildScheduledEventCreate           EventType = "GUILD_SCHEDULED_EVENT_CREATE"
	EventTypeGuildScheduledEventUpdate           EventType = "GUILD_SCHEDULED_EVENT_UPDATE"
	EventTypeGuildScheduledEventDelete           EventType = "GUILD_SCHEDULED_EVENT_DELETE"
	EventTypeGuildScheduledEventUserAdd          EventType = "GUILD_SCHEDULED_EVENT_USER_ADD"
	EventTypeGuildScheduledEventUserRemove       EventType = "GUILD_SCHEDULED_EVENT_USER_REMOVE"
	EventTypeIntegrationCreate                   EventType = "INTEGRATION_CREATE"
	EventTypeIntegrationUpdate                   EventType = "INTEGRATION_UPDATE"
	EventTypeIntegrationDelete                   EventType = "INTEGRATION_DELETE"
	EventTypeInteractionCreate                   EventType = "INTERACTION_CREATE"
	EventTypeInviteCreate                        EventType = "INVITE_CREATE"
	EventTypeInviteDelete                        EventType = "INVITE_DELETE"
	EventTypeMessageCreate                       EventType = "MESSAGE_CREATE"
	EventTypeMessageUpdate                       EventType = "MESSAGE_UPDATE"
	EventTypeMessageDelete                       EventType = "MESSAGE_DELETE"
	EventTypeMessageDeleteBulk                   EventType = "MESSAGE_DELETE_BULK"
	EventTypeMessageReactionAdd                  EventType = "MESSAGE_REACTION_ADD"
	EventTypeMessageReactionRemove               EventType = "MESSAGE_REACTION_REMOVE"
	EventTypeMessageReactionRemoveAll            EventType = "MESSAGE_REACTION_REMOVE_ALL"
	EventTypeMessageReactionRemoveEmoji          EventType = "MESSAGE_REACTION_REMOVE_EMOJI"
	EventTypePresenceUpdate                      EventType = "PRESENCE_UPDATE"
	EventTypeStageInstanceCreate                 EventType = "STAGE_INSTANCE_CREATE"
	EventTypeStageInstanceUpdate                 EventType = "STAGE_INSTANCE_UPDATE"
	EventTypeStageInstanceDelete                 EventType = "STAGE_INSTANCE_DELETE"
	EventTypeTypingStart                         EventType = "TYPING_START"
	EventTypeUserUpdate                          EventType = "USER_UPDATE"
	EventTypeVoiceStateUpdate                    EventType = "VOICE_STATE_UPDATE"
	EventTypeVoiceServerUpdate                   EventType = "VOICE_SERVER_UPDATE"
	EventTypeWebhooksUpdate                      EventType = "WEBHOOKS_UPDATE"
)

func UnmarshalEvent(t EventType, raw []byte) (interface{}, error) {
	switch t {
	case EventTypeDisconnect:
		return &DisconnectEvent{}, nil
	case EventTypeConnect:
		return &ConnectEvent{}, nil
	case EventTypeHello:
		return decodeT[*HelloData](raw)
	case EventTypeReady:
		return decodeT[*ReadyEvent](raw)
	case EventTypeResumed:
		return &ResumedEvent{}, nil
	case EventTypeApplicationCommandPermissionsUpdate:
		return decodeT[*ApplicationCommandPermissionsUpdateEvent](raw)
	case EventTypeAutoModerationRuleCreate:
		return decodeT[*AutoModerationRuleCreateEvent](raw)
	case EventTypeAutoModerationRuleUpdate:
		return decodeT[*AutoModerationRuleUpdateEvent](raw)
	case EventTypeAutoModerationRuleDelete:
		return decodeT[*AutoModerationRuleDeleteEvent](raw)
	case EventTypeAutoModerationActionExecution:
		return decodeT[*AutoModerationActionExecutionEvent](raw)
	case EventTypeChannelCreate:
		return decodeT[*ChannelCreateEvent](raw)
	case EventTypeChannelUpdate:
		return decodeT[*ChannelUpdateEvent](raw)
	case EventTypeChannelDelete:
		return decodeT[*ChannelDeleteEvent](raw)
	case EventTypeChannelPinsUpdate:
		return decodeT[*ChannelPinsUpdateEvent](raw)
	case EventTypeThreadCreate:
		return decodeT[*ThreadCreateEvent](raw)
	case EventTypeThreadUpdate:
		return decodeT[*ThreadUpdateEvent](raw)
	case EventTypeThreadDelete:
		return decodeT[*ThreadDeleteEvent](raw)
	case EventTypeThreadListSync:
		return decodeT[*ThreadListSyncEvent](raw)
	case EventTypeThreadMemberUpdate:
		return decodeT[*ThreadMemberUpdateEvent](raw)
	case EventTypeThreadMembersUpdate:
		return decodeT[*ThreadMembersUpdateEvent](raw)
	case EventTypeEntitlementCreate:
		return decodeT[*EntitlementCreateEvent](raw)
	case EventTypeEntitlementUpdate:
		return decodeT[*EntitlementUpdateEvent](raw)
	case EventTypeEntitlementDelete:
		return decodeT[*EntitlementDeleteEvent](raw)
	case EventTypeGuildCreate:
		return decodeT[*GuildCreateEvent](raw)
	case EventTypeGuildUpdate:
		return decodeT[*GuildUpdateEvent](raw)
	case EventTypeGuildDelete:
		return decodeT[*GuildDeleteEvent](raw)
	case EventTypeGuildAuditLogEntryCreate:
		return decodeT[*AuditLogEntryCreateEvent](raw)
	case EventTypeGuildBanAdd:
		return decodeT[*BanAddEvent](raw)
	case EventTypeGuildBanRemove:
		return decodeT[*BanRemoveEvent](raw)
	case EventTypeGuildEmojisUpdate:
		return decodeT[*GuildEmojisUpdateEvent](raw)
	case EventTypeGuildStickersUpdate:
		return decodeT[*GuildStickersUpdateEvent](raw)
	case EventTypeGuildIntegrationsUpdate:
		return decodeT[*GuildIntegrationsUpdateEvent](raw)
	case EventTypeGuildMemberAdd:
		return decodeT[*MemberAddEvent](raw)
	case EventTypeGuildMemberRemove:
		return decodeT[*MemberRemoveEvent](raw)
	case EventTypeGuildMemberUpdate:
		return decodeT[*MemberUpdateEvent](raw)
	case EventTypeGuildMembersChunk:
		return decodeT[*MemberChunkEvent](raw)
	case EventTypeGuildRoleCreate:
		return decodeT[*RoleCreateEvent](raw)
	case EventTypeGuildRoleUpdate:
		return decodeT[*RoleUpdateEvent](raw)
	case EventTypeGuildRoleDelete:
		return decodeT[*RoleDeleteEvent](raw)
	case EventTypeGuildScheduledEventCreate:
		return decodeT[*ScheduledEventCreateEvent](raw)
	case EventTypeGuildScheduledEventUpdate:
		return decodeT[*ScheduledEventUpdateEvent](raw)
	case EventTypeGuildScheduledEventDelete:
		return decodeT[*ScheduledEventDeleteEvent](raw)
	case EventTypeGuildScheduledEventUserAdd:
		return decodeT[*ScheduledEventUserAddEvent](raw)
	case EventTypeGuildScheduledEventUserRemove:
		return decodeT[*ScheduledEventUserRemoveEvent](raw)
	case EventTypeIntegrationCreate:
		return decodeT[*IntegrationCreateEvent](raw)
	case EventTypeIntegrationUpdate:
		return decodeT[*IntegrationUpdateEvent](raw)
	case EventTypeIntegrationDelete:
		return decodeT[*IntegrationDeleteEvent](raw)
	case EventTypeInteractionCreate:
		return decodeT[*InteractionCreateEvent](raw)
	case EventTypeInviteCreate:
		return decodeT[*InviteCreateEvent](raw)
	case EventTypeInviteDelete:
		return decodeT[*InviteDeleteEvent](raw)
	case EventTypeMessageCreate:
		return decodeT[*MessageCreateEvent](raw)
	case EventTypeMessageUpdate:
		return decodeT[*MessageUpdateEvent](raw)
	case EventTypeMessageDelete:
		return decodeT[*MessageDeleteEvent](raw)
	case EventTypeMessageDeleteBulk:
		return decodeT[*MessageDeleteBulkEvent](raw)
	case EventTypeMessageReactionAdd:
		return decodeT[*MessageReactionAddEvent](raw)
	case EventTypeMessageReactionRemove:
		return decodeT[*MessageReactionRemoveEvent](raw)
	case EventTypeMessageReactionRemoveAll:
		return decodeT[*MessageReactionRemoveAllEvent](raw)
	case EventTypeMessageReactionRemoveEmoji:
		return decodeT[*MessageReactionRemoveEmojiEvent](raw)
	case EventTypePresenceUpdate:
		return decodeT[*PresenceUpdateEvent](raw)
	case EventTypeStageInstanceCreate:
		return decodeT[*StageInstanceCreateEvent](raw)
	case EventTypeStageInstanceUpdate:
		return decodeT[*StageInstanceUpdateEvent](raw)
	case EventTypeStageInstanceDelete:
		return decodeT[*StageInstanceDeleteEvent](raw)
	case EventTypeTypingStart:
		return decodeT[*TypingStartEvent](raw)
	case EventTypeUserUpdate:
		return decodeT[*UserUpdateEvent](raw)
	case EventTypeVoiceStateUpdate:
		return decodeT[*VoiceStateUpdateEvent](raw)
	case EventTypeVoiceServerUpdate:
		return decodeT[*VoiceServerUpdateEvent](raw)
	case EventTypeWebhooksUpdate:
		return decodeT[*WebhooksUpdateEvent](raw)
	default:
		return nil, fmt.Errorf("unknown event type: %s", t)
	}
}

type DisconnectEvent struct{}

type ConnectEvent struct{}

type ReadyEvent struct {
	V                int                `json:"v"`
	User             User               `json:"user"`
	Guilds           []UnavailableGuild `json:"guilds"`
	SessionID        string             `json:"session_id"`
	ResumeGatewayURL string             `json:"resume_gateway_url"`
	Shard            [2]int             `json:"shard"`
	Application      Application        `json:"application"`
}

type ResumedEvent struct{}
