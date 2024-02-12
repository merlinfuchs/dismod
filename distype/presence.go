package distype

type Status string

const (
	StatusOnline    Status = "online"
	StatusDND       Status = "dnd"
	StatusIdle      Status = "idle"
	StatusInvisible Status = "invisible"
	StatusOffline   Status = "offline"
)

type ClientStatus struct {
	Desktop *Status `json:"desktop,omitempty"`
	Mobile  *Status `json:"mobile,omitempty"`
	Web     *Status `json:"web,omitempty"`
}

type PresenceUpdateEvent struct {
	User         User         `json:"user"`
	GuildID      Snowflake    `json:"guild_id"`
	Status       Status       `json:"status"`
	Activities   []Activity   `json:"activities"`
	ClientStatus ClientStatus `json:"client_status"`
}
