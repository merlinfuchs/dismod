package distype

type Activity struct {
	Name          string              `json:"name"`
	Type          ActivityType        `json:"type"`
	URL           *Nullable[string]   `json:"url,omitempty"`
	CreatedAt     UnixTimestamp       `json:"created_at"`
	Timestamps    *ActivityTimestamps `json:"timestamps,omitempty"`
	ApplicationID *Snowflake          `json:"application_id,omitempty"`
	Details       *string             `json:"details,omitempty"`
	State         *string             `json:"state,omitempty"`
	Emoji         *ActivityEmoji      `json:"emoji,omitempty"`
	Party         *ActivityParty      `json:"party,omitempty"`
	Assets        *ActivityAssets     `json:"assets,omitempty"`
	Secrets       *ActivitySecrets    `json:"secrets,omitempty"`
	Instance      *bool               `json:"instance,omitempty"`
	Flags         *ActivityFlags      `json:"flags,omitempty"`
	Buttons       []ActivityButton    `json:"buttons,omitempty"`
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
	Start *UnixTimestamp `json:"start,omitempty"`
	End   *UnixTimestamp `json:"end,omitempty"`
}

type ActivityEmoji struct {
	Name     string     `json:"name"`
	ID       *Snowflake `json:"id,omitempty"`
	Animated *bool      `json:"animated,omitempty"`
}

type ActivityParty struct {
	ID   *Snowflake `json:"id,omitempty"`
	Size [2]int     `json:"size"`
}

type ActivityAssets struct {
	LargeImage *string `json:"large_image,omitempty"`
	LargeText  *string `json:"large_text,omitempty"`
	SmallImage *string `json:"small_image,omitempty"`
	SmallText  *string `json:"small_text,omitempty"`
}

type ActivitySecrets struct {
	Join     *string `json:"join,omitempty"`
	Spectate *string `json:"spectate,omitempty"`
	Match    *string `json:"match,omitempty"`
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
