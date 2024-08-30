package conf

import (
	"strconv"
	"sync"
)

var debugMx sync.RWMutex

func GetDebug(key string) string {
	debugMx.RLock()
	defer debugMx.RUnlock()

	return debugConfMap[key]
}

func BoolDebug(key string) bool {
	if v, err := strconv.ParseBool(GetDebug(key)); err == nil {
		return v
	}
	return false
}
