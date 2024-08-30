package util

import (
	"strings"
)

/*
func IsNotEmptyString(s ...string) (bool, error) {
	for _, item := range s {
		if len(strings.TrimSpace(item)) == 0 {
			return false, errors.New("Empty string found")
		}
	}
	return true, nil
}
*/

func IsNotEmptyString(s ...string) bool {
	for _, item := range s {
		//fmt.Println("item: ", item)
		if len(strings.TrimSpace(item)) == 0 {

			return false
		}
	}
	return true
}
