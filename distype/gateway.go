package distype

import (
	"encoding/json"
)

type GatewayEvent struct {
	Operation GatewayOpcode `json:"op"`
	Sequence  int64         `json:"s"`
	Type      EventType     `json:"t"`
	Data      interface{}   `json:"d"`
}

func (e *GatewayEvent) UnmarshalJSON(raw []byte) error {
	type rawGatewayEvent struct {
		Operation GatewayOpcode   `json:"op"`
		Sequence  int64           `json:"s"`
		Type      EventType       `json:"t"`
		Data      json.RawMessage `json:"d"`
	}

	var tmp rawGatewayEvent
	err := json.Unmarshal(raw, &tmp)
	if err != nil {
		return err
	}

	*e = GatewayEvent{
		Operation: tmp.Operation,
		Sequence:  tmp.Sequence,
		Type:      tmp.Type,
	}

	switch tmp.Operation {
	case GatewayOpcodeDispatch:
		e.Data, err = UnmarshalEvent(tmp.Type, tmp.Data)
	case GatewayOpcodeHeartbeat:
		e.Data, err = decodeT[*HeartbeatData](tmp.Data)
	case GatewayOpcodeIdentify:
		e.Data, err = decodeT[*IdentifyData](tmp.Data)
	case GatewayOpcodePresenceUpdate:
		e.Data, err = decodeT[*PresenceUpdateData](tmp.Data)
	case GatewayOpcodeVoiceStateUpdate:
		e.Data, err = decodeT[*VoiceStateUpdataData](tmp.Data)
	case GatewayOpcodeResume:
		e.Data, err = decodeT[*ResumeData](tmp.Data)
	case GatewayOpcodeReconnect:
	case GatewayOpcodeRequestGuildMembers:
		e.Data, err = decodeT[*RequestGuildMembersData](tmp.Data)
	case GatewayOpcodeInvalidSession:
		e.Data, err = decodeT[*SessionInvalidData](tmp.Data)
	case GatewayOpcodeHello:
		e.Data, err = decodeT[*HelloData](tmp.Data)
	case GatewayOpcodeHeartbeatAck:
	}

	return err
}

type GatewayOpcode int

const (
	GatewayOpcodeDispatch            GatewayOpcode = 0
	GatewayOpcodeHeartbeat           GatewayOpcode = 1
	GatewayOpcodeIdentify            GatewayOpcode = 2
	GatewayOpcodePresenceUpdate      GatewayOpcode = 3
	GatewayOpcodeVoiceStateUpdate    GatewayOpcode = 4
	GatewayOpcodeResume              GatewayOpcode = 5
	GatewayOpcodeReconnect           GatewayOpcode = 6
	GatewayOpcodeRequestGuildMembers GatewayOpcode = 8
	GatewayOpcodeInvalidSession      GatewayOpcode = 9
	GatewayOpcodeHello               GatewayOpcode = 10
	GatewayOpcodeHeartbeatAck        GatewayOpcode = 11
)

type HelloData struct {
	HeartbeatInterval int `json:"heartbeat_interval"`
}

type HeartbeatData int64

type IdentifyData struct {
	Token          string               `json:"token"`
	Properties     ConnectionProperties `json:"properties"`
	Compress       *bool                `json:"compress,omitempty"`
	LargeThreshold *int                 `json:"large_threshold,omitempty"`
	Shard          *[2]int              `json:"shard,omitempty"`
	Presence       *PresenceUpdateData  `json:"presence,omitempty"`
	Intents        int                  `json:"intents"`
}

type ConnectionProperties struct {
	OS      string `json:"$os"`
	Browser string `json:"$browser"`
	Device  string `json:"$device"`
}

type ResumeData struct {
	Token     string `json:"token"`
	SessionID string `json:"session_id"`
	Sequence  int64  `json:"seq"`
}

type PresenceUpdateData struct {
	Since      Nullable[UnixTimestamp] `json:"since"`
	Activities []Activity              `json:"activities"`
	Status     Status                  `json:"status"`
	AFK        bool                    `json:"afk"`
}

type RequestGuildMembersData struct {
	GuildID   Snowflake `json:"guild_id"`
	Query     *string   `json:"query,omitempty"`
	Limit     int       `json:"limit"`
	Presences *bool     `json:"presences,omitempty"`
	UserIDs   []string  `json:"user_ids,omitempty"`
	Nonce     *string   `json:"nonce,omitempty"`
}

type SessionInvalidData bool

type VoiceStateUpdataData struct {
	GuildID   Snowflake           `json:"guild_id"`
	ChannelID Nullable[Snowflake] `json:"channel_id"`
	SelfMute  bool                `json:"self_mute"`
	SelfDeaf  bool                `json:"self_deaf"`
}

type GatewayBotGetRequest struct{}

type GatewayBotGetResponse struct {
	URL    string `json:"url"`
	Shards int    `json:"shards"`
	SessionStartLimit
}

type SessionStartLimit struct {
	Total          int `json:"total"`
	Remaining      int `json:"remaining"`
	ResetAfter     int `json:"reset_after"`
	MaxConcurrency int `json:"max_concurrency"`
}
