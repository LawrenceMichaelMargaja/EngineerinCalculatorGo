package utils

import "encoding/json"

func EncloseString(s string, d string) string {
	return `` + d + `` + s + d + ``
}

func ToJSON(i interface{}) string {
	_json, err := json.Marshal(i)
	if err != nil {
		return "error marshaling"
	} else {
		return string(_json)
	}
}