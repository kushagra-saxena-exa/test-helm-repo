package utils

import "encoding/json"

func ConvertToJson(inp interface{}) ([]byte, error) {
	b, err := json.Marshal(inp)
	if err != nil {
		return nil, err
	}
	return b, err
}
