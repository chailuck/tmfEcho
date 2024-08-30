package conf

import "strings"

func Of(key string) map[string]string {
	list := make(map[string]string)

	mx.RLock()
	defer mx.RUnlock()

	for k, v := range confMap {
		if strings.HasPrefix(k, key) {
			newKey := strings.Replace(k, key+".", "", 1)
			list[newKey] = v
		}
	}
	return list
}
