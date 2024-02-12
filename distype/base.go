package distype

import (
	"encoding/json"
	"time"
)

type Nullable[T any] struct {
	Valid bool
	Value T
}

func NullString(s string, valid bool) Nullable[string] {
	return Nullable[string]{Valid: valid, Value: s}
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

type UnixTimestamp time.Time

func (t UnixTimestamp) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(t).Unix())
}

func (t *UnixTimestamp) UnmarshalJSON(data []byte) error {
	var i int64
	err := json.Unmarshal(data, &i)
	if err != nil {
		return err
	}
	*t = UnixTimestamp(time.Unix(i, 0))
	return nil
}
