package conf_test

import (
	"GOKIT_v001/internal/conf"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetDuration(t *testing.T) {
	conf.SetFake("my.token.timeout", "10s")

	assert.Equal(t, 10*time.Second, conf.GetDuration("my.token.timeout", 0))
	assert.Equal(t, 5*time.Second, conf.GetDuration("no.token.timeout", 5*time.Second))
}
