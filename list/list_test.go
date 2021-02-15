package list_test

import (
	"testing"

	"github.com/LeonZYang/datastruct/list"
	"github.com/stretchr/testify/assert"
)

func TestLinkedList(t *testing.T) {
	assert := assert.New(t)

	ll := list.NewLinkList()

	ll.PushFront("a", 1)
	ll.PushFront("b", 2)
	ll.PushFront("c", 3)

	assert.Equal(3, ll.Len())

	assert.Equal("c", ll.Front().Key)
	assert.Equal(3, ll.Front().Value)

	assert.Equal(1, ll.Back().Value)
	assert.Equal("a", ll.Back().Key)
}

func TestLinkedListRemove(t *testing.T) {
	assert := assert.New(t)

	ll := list.NewLinkList()

	ll.PushFront("a", 1)
	ll.PushFront("b", 2)
	ll.PushFront("c", 3)

	ll.Remove(ll.Back())

	assert.Equal(2, ll.Back().Value)

	assert.Equal(3, ll.Front().Value)
	assert.Equal(2, ll.Len())
}

func TestLinkedListMoveToFront(t *testing.T) {
	assert := assert.New(t)

	ll := list.NewLinkList()

	ll.PushFront("a", 1)
	ll.PushFront("b", 2)
	ll.PushFront("c", 3)

	ll.MoveToFront(ll.Back())

	assert.Equal(2, ll.Back().Value)

	assert.Equal(1, ll.Front().Value)
	assert.Equal(3, ll.Len())
}
