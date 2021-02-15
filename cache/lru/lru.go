package lru

import (
	"github.com/LeonZYang/datastruct/list"
)

type LRU struct {
	data map[string]*list.Element
	list *list.LinkList
	cap  int
}

func New(cap int) *LRU {
	return &LRU{
		data: make(map[string]*list.Element),
		list: list.NewLinkList(),
		cap:  cap,
	}
}

func (l *LRU) Put(key string, value interface{}) {
	if node, ok := l.data[key]; ok {
		node.Value = value
		l.list.MoveToFront(node)
		return
	}

	if l.Len() >= l.cap {
		head := l.list.Back()
		delete(l.data, head.Key)
		l.list.Remove(head)
	}

	node := list.NewElement(key, value)
	l.data[key] = node
	l.list.PushFrontElement(node)
	return
}

func (l *LRU) Get(key string) interface{} {
	node, ok := l.data[key]
	if !ok {
		return nil
	}
	l.list.MoveToFront(node)
	return node.Value
}

func (l *LRU) Len() int {
	return len(l.data)
}

func (l *LRU) Cap() int {
	return l.cap
}
