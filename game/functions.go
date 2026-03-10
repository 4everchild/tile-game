package game

import (
	"encoding/json"
)

func temp_add(a, b uint8) uint8 {
	return a + b
}

func GameToJson(g Game) string {
	b, err := json.Marshal(g)
	if err != nil {
		return ""
	}
	return string(b)
}
