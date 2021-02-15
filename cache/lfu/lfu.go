package lfu

import (
	"github.com/LeonZYang/datastruct/list"
)

type LFU struct {
	data    map[string]*list.Element
	freqs   map[int]*list.LinkList
	cap     int
	minFreq int
}

func New(cap int) *LFU {
	return &LFU{
		data:    make(map[string]*list.Element),
		freqs:   make(map[int]*list.LinkList),
		cap:     cap,
		minFreq: 0,
	}
}

func (l *LFU) Put(key string, value interface{}) {
	if l.cap == 0 {
		return
	}
	if node, ok := l.data[key]; ok {
		node.Value = value
		l.increase(node)
		return
	}

	if l.Len() >= l.Cap() {
		last := l.freqs[l.minFreq].Back()
		l.freqs[l.minFreq].Remove(last)
		delete(l.data, last.Key)
	}

	node := list.NewElement(key, value)
	node.Freq = 1
	l.minFreq = 1
	if _, ok := l.freqs[l.minFreq]; !ok {
		l.freqs[l.minFreq] = list.NewLinkList()
	}
	l.freqs[l.minFreq].PushFrontElement(node)
	l.data[key] = node
}

func (l *LFU) Get(key string) interface{} {
	node, ok := l.data[key]
	if !ok {
		return nil
	}
	l.increase(node)
	return node.Value
}

func (l *LFU) increase(node *list.Element) {
	freq := node.Freq
	l.freqs[freq].Remove(node)

	node.Freq++
	if _, ok := l.freqs[node.Freq]; !ok {
		l.freqs[node.Freq] = list.NewLinkList()
	}

	l.freqs[node.Freq].PushFrontElement(node)
	if l.freqs[freq].Len() == 0 {
		delete(l.freqs, freq)
		if freq == l.minFreq {
			l.minFreq++
		}
	}
}

func (l *LFU) Len() int {
	return len(l.data)
}

func (l *LFU) Cap() int {
	return l.cap
}
