package json

import (
	"encoding/json"
)

func JsonDecode[T interface{}](jsonString string) (T, error) {
	emptyType := new(T)
	if err := json.Unmarshal([]byte(jsonString), &emptyType); err != nil {
		return *emptyType, err
	}
	return *emptyType, nil
}

func JsonEncode(obj *interface{}) (string, error) {
	result, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(result), nil
}
