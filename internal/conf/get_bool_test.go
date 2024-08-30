package conf_test

import (
	"testing"

	"tmfEcho/internal/conf"

	"github.com/stretchr/testify/assert"
)

func TestGetBool(t *testing.T) {
	conf.SetFake("key.item.1", "true")

	assert.True(t, conf.GetBool("key.item.1"))
	assert.False(t, conf.GetBool("key.item.2"))
}
