package lfu_test

import (
	"testing"

	"github.com/LeonZYang/datastruct/cache/lfu"
	"github.com/stretchr/testify/assert"
)

func TestLFU(t *testing.T) {
	assert := assert.New(t)

	l := lfu.New(2)
	assert.NotNil(l)
	assert.Equal(0, l.Len())
	assert.Equal(2, l.Cap())

	l.Put("a", 1)
	l.Put("b", 2)
	assert.Equal(2, l.Len())
	assert.Equal(1, l.Get("a"))
	assert.Equal(2, l.Get("b"))

	l.Put("c", 3)
	assert.Equal(nil, l.Get("a"))

	assert.Equal(3, l.Get("c"))
	assert.Equal(2, l.Get("b"))

	l.Put("d", 4)

	assert.Equal(nil, l.Get("c"))

	l.Get("d")
	l.Get("d")
	l.Put("e", 5)
	assert.Equal(nil, l.Get("b"))
}

func TestLFUUpdate(t *testing.T) {
	assert := assert.New(t)

	l := lfu.New(2)
	assert.NotNil(l)
	assert.Equal(0, l.Len())
	assert.Equal(2, l.Cap())

	l.Put("a", 1)
	l.Put("b", 2)
	assert.Equal(2, l.Len())
	assert.Equal(1, l.Get("a"))
	assert.Equal(2, l.Get("b"))

	l.Put("b", 3)

	assert.Equal(1, l.Get("a"))
	assert.Equal(3, l.Get("b"))
}
