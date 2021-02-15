package lru_test

import (
	"testing"

	"github.com/LeonZYang/datastruct/cache/lru"
	"github.com/stretchr/testify/assert"
)

func TestLRU(t *testing.T) {
	assert := assert.New(t)

	l := lru.New(2)
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
}
