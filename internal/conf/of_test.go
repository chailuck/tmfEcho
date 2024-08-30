package conf_test

import (
	"testing"
	"tmfEcho/internal/conf"

	"github.com/stretchr/testify/assert"
)

func TestOf(t *testing.T) {
	conf.SetFake("key.item.1", "item1")
	conf.SetFake("key.item.2", "item2")

	x := map[string]string{
		"1": "item1",
		"2": "item2",
	}
	assert.Equal(t, x, conf.Of("key.item"))
}
