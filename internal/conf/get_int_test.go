package conf_test

import (
	"testing"
	"tmfEcho/internal/conf"

	"github.com/stretchr/testify/assert"
)

func TestGetInt(t *testing.T) {
	conf.SetFake("key.item.1", "234")

	assert.Equal(t, 234, conf.GetInt("key.item.1"))
	assert.Equal(t, -99999, conf.GetInt("key.item.2"))
}
