package utils

import "encoding/json"

func FromJson(dest interface{}, data string) error {
	return json.Unmarshal([]byte(data), dest)
}

func ToJson(value interface{}) (string, error) {
	v, err := json.Marshal(value)
	if err != nil {
		return "", err
	}

	return string(v), err
}

func ToJsonStr(value interface{}) string {
	v, err := json.Marshal(value)
	if err != nil {
		return ""
	}

	return string(v)
}
