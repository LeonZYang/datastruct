package datastruct

import "github.com/LeonZYang/datastruct/linkedlist"

type Stack struct {
	list *linkedlist.LinkList
}

func NewStack() *Stack {
	return &Stack{linkedlist.NewLinkList()}
}

func (s *Stack) Push(i interface{}) {
	s.list.PushFront(i)
}

func (s *Stack) Pop() interface{} {
	if s.list.IsEmpty() {
		return nil
	}
	b := s.list.Front()
	s.list.Remove(b)
	return b.Value
}

func (s *Stack) Len() int {
	return s.list.Len()
}
