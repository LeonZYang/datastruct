package skiplist_test

import (
	"testing"

	"github.com/LeonZYang/datastruct/skiplist"
	"github.com/stretchr/testify/assert"
)

func TestSkipList(t *testing.T) {
	assert := assert.New(t)

	sl := skiplist.New(5)

	sl.Insert(1, 1)
	assert.Equal(1, sl.Len())
	val, ok := sl.Search(1)
	assert.True(ok)
	assert.Equal(1, val)

	val, ok = sl.Search(2)
	assert.False(ok)
	sl.Insert(2, 2)
	assert.Equal(2, sl.Len())
	sl.Insert(3, 3)
	sl.Remove(2)
	assert.Equal(2, sl.Len())
	val, ok = sl.Search(2)
	assert.False(ok)
	assert.Equal(0, val)
}

func TestSkipListUpdate(t *testing.T) {
	assert := assert.New(t)

	sl := skiplist.New(5)

	sl.Insert(1, 1)
	assert.Equal(1, sl.Len())
	val, ok := sl.Search(1)
	assert.True(ok)
	assert.Equal(1, val)

	sl.Insert(1, 1)
	assert.Equal(1, sl.Len())
	val, ok = sl.Search(1)
	assert.True(ok)
	assert.Equal(1, val)

	sl.Insert(1, 2)
	assert.Equal(2, sl.Len())
	val, ok = sl.Search(1)
	assert.True(ok)
	assert.Equal(2, val)
}
