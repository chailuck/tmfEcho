package conf

import "sync"

var mx sync.RWMutex

func Get(key string) string {
	mx.RLock()
	defer mx.RUnlock()

	return confMap[key]
}
