package distype

import "encoding/json"

func decodeT[T any](raw []byte) (T, error) {
	var data T
	err := json.Unmarshal(raw, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

func Optional[T any](v T) *T {
	return &v
}

func NullString(s string, valid bool) Nullable[string] {
	return Nullable[string]{Valid: valid, Value: s}
}
