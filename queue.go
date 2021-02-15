package datastruct

import "github.com/LeonZYang/datastruct/linkedlist"

type Queue struct {
	list *linkedlist.LinkList
}

func NewQueue() *Queue {
	return &Queue{linkedlist.NewLinkList()}
}

func (q *Queue) Enqueue(i interface{}) {
	q.list.PushFront(i)
}

func (q *Queue) Dequeue() interface{} {
	if q.list.IsEmpty() {
		return nil
	}
	b := q.list.Back()
	q.list.Remove(b)
	return b.Value
}

func (q *Queue) Len() int {
	return q.list.Len()
}
