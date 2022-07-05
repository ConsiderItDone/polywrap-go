package json

import "encoding/json"

func Decode(value string) (*map[string]interface{}, error) {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(value), &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func Encode(data *map[string]interface{}) (string, error) {
	out, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return string(out), nil
}
