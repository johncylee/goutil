package goutil

import (
	"encoding/json"
)

// Check if data is valid JSON string
func IsJSONString(data []byte) bool {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return false
	}
	return true
}
