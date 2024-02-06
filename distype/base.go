package distype

import (
	"encoding/json"
)

type Snowflake string

type Nullable[T any] struct {
	Valid bool
	Value T
}

func (n Nullable[T]) MarshalJSON() ([]byte, error) {
	if n.Valid {
		return json.Marshal(n.Value)
	}
	return []byte("null"), nil
}

func (n *Nullable[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		n.Valid = false
		return nil
	}
	n.Valid = true
	return json.Unmarshal(data, &n.Value)
}

type Optional[T any] *T

type IntOrString struct {
	Int    int
	String string
}

func (i IntOrString) MarshalJSON() ([]byte, error) {
	if i.Int != 0 {
		return json.Marshal(i.Int)
	}
	return json.Marshal(i.String)
}

func (i *IntOrString) UnmarshalJSON(data []byte) error {
	if data[0] == '"' {
		i.Int = 0
		return json.Unmarshal(data, &i.String)
	}
	i.String = ""
	return json.Unmarshal(data, &i.Int)
}
