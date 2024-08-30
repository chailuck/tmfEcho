package util

import (
	"strings"
)

func IsNotEmptyString(s ...string) bool {
	for _, item := range s {
		if len(strings.TrimSpace(item)) == 0 {
			return false
		}
	}
	return true
}
