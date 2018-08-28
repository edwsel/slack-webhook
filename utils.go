package slack

import "encoding/json"

func toJson(data interface{}) string {
	stringInByte, err := json.Marshal(data)

	if err != nil {
		return ""
	}

	return string(stringInByte)
}
