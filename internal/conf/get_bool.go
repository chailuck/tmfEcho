package conf

import "strconv"

func GetBool(key string) bool {
	if v, err := strconv.ParseBool(Get(key)); err == nil {
		return v
	}
	return false
}
