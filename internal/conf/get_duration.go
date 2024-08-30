package conf

import "time"

func GetDuration(key string, defaultDuration time.Duration) time.Duration {
	if d, err := time.ParseDuration(Get(key)); err == nil {
		return d
	}
	return defaultDuration
}
