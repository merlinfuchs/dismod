package distype

type Ban struct {
	Reason Nullable[string] `json:"reason"`
	User   User             `json:"user"`
}
