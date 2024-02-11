package distype

type Activity struct {
	Name          string                       `json:"name"`
	Type          ActivityType                 `json:"type"`
	URL           Optional[Nullable[string]]   `json:"url,omitempty"`
	CreatedAt     UnixTimestamp                `json:"created_at"`
	Timestamps    Optional[ActivityTimestamps] `json:"timestamps,omitempty"`
	ApplicationID Optional[Snowflake]          `json:"application_id,omitempty"`
	Details       Optional[string]             `json:"details,omitempty"`
	State         Optional[string]             `json:"state,omitempty"`
	Emoji         Optional[ActivityEmoji]      `json:"emoji,omitempty"`
	Party         Optional[ActivityParty]      `json:"party,omitempty"`
	Assets        Optional[ActivityAssets]     `json:"assets,omitempty"`
	Secrets       Optional[ActivitySecrets]    `json:"secrets,omitempty"`
	Instance      Optional[bool]               `json:"instance,omitempty"`
	Flags         Optional[ActivityFlags]      `json:"flags,omitempty"`
	Buttons       []ActivityButton             `json:"buttons,omitempty"`
}

type ActivityType int

const (
	ActivityTypeGame      ActivityType = 0
	ActivityTypeStreaming ActivityType = 1
	ActivityTypeListening ActivityType = 2
	ActivityTypeWatching  ActivityType = 3
	ActivityTypeCustom    ActivityType = 4
	ActivityTypeCompeting ActivityType = 5
)

type ActivityTimestamps struct {
	Start Optional[UnixTimestamp] `json:"start,omitempty"`
	End   Optional[UnixTimestamp] `json:"end,omitempty"`
}

type ActivityEmoji struct {
	Name     string              `json:"name"`
	ID       Optional[Snowflake] `json:"id,omitempty"`
	Animated Optional[bool]      `json:"animated,omitempty"`
}

type ActivityParty struct {
	ID   Optional[Snowflake] `json:"id,omitempty"`
	Size [2]int              `json:"size"`
}

type ActivityAssets struct {
	LargeImage Optional[string] `json:"large_image,omitempty"`
	LargeText  Optional[string] `json:"large_text,omitempty"`
	SmallImage Optional[string] `json:"small_image,omitempty"`
	SmallText  Optional[string] `json:"small_text,omitempty"`
}

type ActivitySecrets struct {
	Join     Optional[string] `json:"join,omitempty"`
	Spectate Optional[string] `json:"spectate,omitempty"`
	Match    Optional[string] `json:"match,omitempty"`
}

type ActivityFlags int

const (
	ActivityFlagsInstance                 ActivityFlags = 1 << 0
	ActivityFlagsJoin                     ActivityFlags = 1 << 1
	ActivityFlagsSpectate                 ActivityFlags = 1 << 2
	ActivityFlagsJoinRequest              ActivityFlags = 1 << 3
	ActivityFlagsSync                     ActivityFlags = 1 << 4
	ActivityFlagsPlay                     ActivityFlags = 1 << 5
	ActivityFlagsPartyPrivacyFriends      ActivityFlags = 1 << 6
	ActivityFlagsPartyPrivacyVoiceChannel ActivityFlags = 1 << 7
	ActivityFlagsEmbedded                 ActivityFlags = 1 << 12
)

type ActivityButton struct {
	Label string `json:"label"`
	URL   string `json:"url"`
}
