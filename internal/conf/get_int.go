package conf

import "strconv"

func GetInt(key string) int {
	if v, err := strconv.Atoi(Get(key)); err == nil {
		return v
	}
	return -99999
}
