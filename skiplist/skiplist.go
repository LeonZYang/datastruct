package skiplist

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

type Node struct {
	Key     int
	Value   int
	Forward []*Node
}

func NewNode(level int, key int, value int) *Node {
	return &Node{
		Key:     key,
		Value:   value,
		Forward: make([]*Node, level),
	}
}

type SkipList struct {
	header *Node

	level    int
	maxLevel int
	size     int
}

func New(maxLevel int) *SkipList {
	return &SkipList{
		header:   NewNode(maxLevel, -1, -1),
		level:    0,
		maxLevel: maxLevel,
	}
}

func (s *SkipList) Search(key int) (int, bool) {
	updateList := s.GetUpdateList(key)
	n, ok := s.Find(updateList, key)
	if !ok {
		return 0, false
	}
	return n.Value, true
}

func (s *SkipList) Insert(key int, value int) {
	updateList := s.GetUpdateList(key)

	// 存在
	q, ok := s.Find(updateList, key)
	if ok && q.Value == value {
		return
	}

	k := s.Random()
	if k > s.maxLevel {
		k = s.maxLevel
	}
	for lv := s.level; lv < k; lv++ {
		updateList[lv] = s.header
		s.level++
	}

	q = NewNode(k, key, value)

	for i := 0; i < k; i++ {
		q.Forward[i] = updateList[i].Forward[i]
		updateList[i].Forward[i] = q
	}
	s.size++
}

func (s *SkipList) Remove(key int) {
	updateList := s.GetUpdateList(key)
	q, ok := s.Find(updateList, key)
	if !ok {
		return
	}

	for i := 0; i < s.level; i++ {
		if updateList[i].Forward[i] == q {
			updateList[i].Forward[i] = q.Forward[i]
		}
	}
	for i := s.level - 1; i >= 0; i-- {
		if s.header.Forward[i] == nil {
			s.level--
		}
		i--
	}
	s.size--
}

func (s *SkipList) GetUpdateList(key int) (updateList []*Node) {
	updateList = make([]*Node, s.maxLevel)
	i := s.level - 1
	for i >= 0 {
		q := s.header
		for q.Forward[i] != nil && q.Forward[i].Key < key {
			q = q.Forward[i]
		}
		updateList[i] = q
		i--
	}

	return updateList
}

func (s *SkipList) Find(updateList []*Node, key int) (*Node, bool) {
	if len(updateList) == 0 {
		updateList = s.GetUpdateList(key)
	}

	if len(updateList) > 0 {
		if updateList[0] != nil && len(updateList[0].Forward) > 0 && updateList[0].Forward[0] != nil && updateList[0].Forward[0].Key == key {
			return updateList[0].Forward[0], true
		}
	}

	return nil, false
}
func (s *SkipList) Random() int {
	level := 1
	for rand.Intn(10) < 5 && s.maxLevel > level {
		level++
	}

	return level
}

func (s *SkipList) Len() int {
	return s.size
}

func (s *SkipList) Traval() {
	for i := s.level - 1; i >= 0; i-- {
		var ss []string
		for q := s.header; q != nil; q = q.Forward[i] {
			ss = append(ss, strconv.Itoa(q.Key))
		}
		fmt.Println("层数: ", i, strings.Join(ss, "->"))
	}
}
