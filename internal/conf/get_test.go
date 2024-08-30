package conf_test

import (
	"strconv"
	"sync"
	"testing"
	"tmfEcho/internal/conf"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	conf.SetFake("api.port", "9999")
	assert.Equal(t, "9999", conf.Get("api.port"))

	conf.SetFake("not.exist", "")
}

func TestGetConcurrent(t *testing.T) {
	conf.SetFake("some.value", "abc")

	n := 100
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(a int) {
			conf.SetFake("some.value", strconv.Itoa(a))
			assert.NotEmpty(t, conf.Get("some.value"))
			wg.Done()
		}(i)
	}
	wg.Wait()
}
