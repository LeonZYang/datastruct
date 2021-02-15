package linkedlist_test

import (
	"testing"

	"github.com/LeonZYang/datastruct/linkedlist"
	"github.com/stretchr/testify/assert"
)

func TestLinkedList(t *testing.T) {
	assert := assert.New(t)

	ll := linkedlist.NewLinkList()

	ll.PushFront(1)
	ll.PushFront(2)
	ll.PushFront(3)

	assert.Equal(3, ll.Len())

	assert.Equal(3, ll.Front().Value)
	assert.Equal(1, ll.Back().Value)
}

func TestLinkedListRemove(t *testing.T) {
	assert := assert.New(t)

	ll := linkedlist.NewLinkList()

	ll.PushFront(1)
	ll.PushFront(2)
	ll.PushFront(3)
	ll.Remove(ll.Back())

	assert.Equal(2, ll.Back().Value)

	assert.Equal(3, ll.Front().Value)
	assert.Equal(2, ll.Len())
}

func TestLinkedListMoveToFront(t *testing.T) {
	assert := assert.New(t)

	ll := linkedlist.NewLinkList()

	ll.PushFront(1)
	ll.PushFront(2)
	ll.PushFront(3)

	ll.MoveToFront(ll.Back())

	assert.Equal(2, ll.Back().Value)

	assert.Equal(1, ll.Front().Value)
	assert.Equal(3, ll.Len())
}
