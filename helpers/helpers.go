package helpers

import (
	"strings"
)

func MapValueExists(value any, theMap map[any]any) bool {

	for _, val := range theMap {
		if _, ok := val.(string); ok {
			if strings.EqualFold(value.(string), val.(string)) {
				return true
			}
		} else {
			if value == val {
				return true
			}
		}
	}
	return false
}
