package conf

func SetFake(key string, value string) {
	mx.Lock()
	defer mx.Unlock()

	confMap[key] = value
}

func SetFakeDebug(key string, value string) {
	debugMx.Lock()
	defer debugMx.Unlock()

	debugConfMap[key] = value
}
