package distype

type Emoji struct {
	ID             Nullable[Snowflake] `json:"id"`
	Name           Nullable[string]    `json:"name"`
	Roles          []Snowflake         `json:"roles,omitempty"`
	User           Optional[User]      `json:"user,omitempty"`
	RequiredColons Optional[bool]      `json:"required_colons,omitempty"`
	Managed        Optional[bool]      `json:"managed,omitempty"`
	Animated       Optional[bool]      `json:"animated,omitempty"`
	Available      Optional[bool]      `json:"available,omitempty"`
}
