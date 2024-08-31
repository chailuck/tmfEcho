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

func ConvertCommaStringToMap(s string) map[string]bool {
	m := make(map[string]bool)
	if strings.Trim(s, " ") != "" {
		for _, p := range strings.Split(s, ",") {
			m[p] = true
		}
	}
	return m
}
