package restlib

import (
	"strings"
)

func IsEmptyString(value string) bool {
	return len(value) == 0
}

func TrimAndCheckForEmptyString(value *string) bool {
	if value == nil {
		return false
	}
	*value = strings.TrimSpace(*value)
	return len(*value) == 0
}
