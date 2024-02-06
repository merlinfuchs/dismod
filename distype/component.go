package distype

import (
	"encoding/json"
	"fmt"
)

type MessageComponent interface {
	json.Marshaler
	Type() MessageComponentType
}

type unmarshalableMessageComponent struct {
	MessageComponent
}

// UnmarshalJSON is a helper function to unmarshal MessageComponent object.
func (umc *unmarshalableMessageComponent) UnmarshalJSON(src []byte) error {
	var v struct {
		Type MessageComponentType `json:"type"`
	}
	err := json.Unmarshal(src, &v)
	if err != nil {
		return err
	}

	switch v.Type {
	case MessageComponentTypeActionRow:
		umc.MessageComponent = &ActionRow{}
	case MessageComponentTypeButton:
		umc.MessageComponent = &Button{}
	case MessageComponentTypeStringSelect, MessageComponentTypeChannelSelect, MessageComponentTypeUserSelect,
		MessageComponentTypeRoleSelect, MessageComponentTypeMentionableSelect:
		umc.MessageComponent = &SelectMenu{}
	case MessageComponentTypeTextInput:
		umc.MessageComponent = &TextInput{}
	default:
		return fmt.Errorf("unknown component type: %d", v.Type)
	}
	return json.Unmarshal(src, umc.MessageComponent)
}

type MessageComponentType int

const (
	MessageComponentTypeActionRow         MessageComponentType = 1
	MessageComponentTypeButton            MessageComponentType = 2
	MessageComponentTypeStringSelect      MessageComponentType = 3
	MessageComponentTypeTextInput         MessageComponentType = 4
	MessageComponentTypeUserSelect        MessageComponentType = 5
	MessageComponentTypeRoleSelect        MessageComponentType = 6
	MessageComponentTypeMentionableSelect MessageComponentType = 7
	MessageComponentTypeChannelSelect     MessageComponentType = 8
)

type ActionRow struct {
	Components []MessageComponent `json:"components"`
}

func (ActionRow) Type() MessageComponentType {
	return MessageComponentTypeActionRow
}

func (r ActionRow) MarshalJSON() ([]byte, error) {
	type actionsRow ActionRow

	return json.Marshal(struct {
		actionsRow
		Type MessageComponentType `json:"type"`
	}{
		actionsRow: actionsRow(r),
		Type:       r.Type(),
	})
}

func (r *ActionRow) UnmarshalJSON(data []byte) error {
	var v struct {
		RawComponents []unmarshalableMessageComponent `json:"components"`
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	r.Components = make([]MessageComponent, len(v.RawComponents))
	for i, v := range v.RawComponents {
		r.Components[i] = v.MessageComponent
	}

	return err
}

type Button struct {
	Style    ButtonStyle      `json:"style"`
	Label    Optional[string] `json:"label,omitempty"`
	Emoji    Optional[Emoji]  `json:"emoji,omitempty"`
	CustomID Optional[string] `json:"custom_id,omitempty"`
	URL      Optional[string] `json:"url,omitempty"`
	Disabled Optional[bool]   `json:"disabled,omitempty"`
}

func (Button) Type() MessageComponentType {
	return MessageComponentTypeButton
}

func (b Button) MarshalJSON() ([]byte, error) {
	type button Button

	if b.Style == 0 {
		b.Style = ButtonStylePrimary
	}

	return json.Marshal(struct {
		button
		Type MessageComponentType `json:"type"`
	}{
		button: button(b),
		Type:   b.Type(),
	})
}

type ButtonStyle int

const (
	ButtonStylePrimary   ButtonStyle = 1
	ButtonStyleSecondary ButtonStyle = 2
	ButtonStyleSuccess   ButtonStyle = 3
	ButtonStyleDanger    ButtonStyle = 4
	ButtonStyleLink      ButtonStyle = 5
)

type SelectMenu struct {
	MenuType      SelectMenuType       `json:"type"`
	CustomID      string               `json:"custom_id"`
	Options       []SelectOption       `json:"options,omitempty"`
	ChannelTypes  []ChannelType        `json:"channel_types,omitempty"`
	Placeholder   Optional[string]     `json:"placeholder,omitempty"`
	DefaultValues []SelectDefaultValue `json:"default_values,omitempty"`
	MinValues     Optional[int]        `json:"min_values,omitempty"`
	MaxValues     Optional[int]        `json:"max_values,omitempty"`
	Disabled      Optional[bool]       `json:"disabled,omitempty"`
}

func (s SelectMenu) Type() MessageComponentType {
	return MessageComponentType(s.MenuType)
}

func (s SelectMenu) MarshalJSON() ([]byte, error) {
	type selectMenu SelectMenu

	return json.Marshal(struct {
		selectMenu
		Type MessageComponentType `json:"type"`
	}{
		selectMenu: selectMenu(s),
		Type:       s.Type(),
	})
}

type SelectMenuType int

const (
	SelectMenuTypeString      SelectMenuType = SelectMenuType(MessageComponentTypeStringSelect)
	SelectMenuTypeChannel     SelectMenuType = SelectMenuType(MessageComponentTypeChannelSelect)
	SelectMenuTypeUser        SelectMenuType = SelectMenuType(MessageComponentTypeUserSelect)
	SelectMenuTypeRole        SelectMenuType = SelectMenuType(MessageComponentTypeRoleSelect)
	SelectMenuTypeMentionable SelectMenuType = SelectMenuType(MessageComponentTypeMentionableSelect)
)

type SelectOption struct {
	Label       string           `json:"label"`
	Value       string           `json:"value"`
	Description Optional[string] `json:"description,omitempty"`
	Emoji       Optional[Emoji]  `json:"emoji,omitempty"`
	Default     Optional[bool]   `json:"default,omitempty"`
}

type SelectDefaultValue struct {
	ID   string                 `json:"id"`
	Type SelectDefaultValueType `json:"type"`
}

type SelectDefaultValueType string

const (
	SelectDefaultValueTypeUser    SelectDefaultValueType = "user"
	SelectDefaultValueTypeRole    SelectDefaultValueType = "role"
	SelectDefaultValueTypeChannel SelectDefaultValueType = "channel"
)

type TextInput struct {
	CustomID    string           `json:"custom_id"`
	Style       TextInputStyle   `json:"style"`
	Label       string           `json:"label"`
	MinLength   Optional[int]    `json:"min_length,omitempty"`
	MaxLength   Optional[int]    `json:"max_length,omitempty"`
	Required    Optional[bool]   `json:"required,omitempty"`
	Value       Optional[string] `json:"value,omitempty"`
	Placeholder Optional[string] `json:"placeholder,omitempty"`
}

func (TextInput) Type() MessageComponentType {
	return MessageComponentTypeTextInput
}

func (t TextInput) MarshalJSON() ([]byte, error) {
	type textInput TextInput

	return json.Marshal(struct {
		textInput
		Type MessageComponentType `json:"type"`
	}{
		textInput: textInput(t),
		Type:      t.Type(),
	})
}

type TextInputStyle int

const (
	TextInputStyleShort     TextInputStyle = 1
	TextInputStyleParagraph TextInputStyle = 2
)
