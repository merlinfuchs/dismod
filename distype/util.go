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
