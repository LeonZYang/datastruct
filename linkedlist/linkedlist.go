package linkedlist

type Element struct {
	Value      interface{}
	Next, Prev *Element
}

func NewElement(val interface{}) *Element {
	return &Element{Value: val}
}

type LinkList struct {
	head   *Element
	tail   *Element
	length int
}

func NewLinkList() *LinkList {
	ll := &LinkList{
		head: NewElement(nil),
		tail: NewElement(nil),
	}

	ll.head.Next = ll.tail
	ll.tail.Prev = ll.head

	return ll
}

func (ll *LinkList) PushFront(value interface{}) {
	ele := NewElement(value)
	ll.PushFrontElement(ele)
}

func (ll *LinkList) PushFrontElement(ele *Element) {
	ll.head.Next.Prev = ele
	ele.Next = ll.head.Next
	ele.Prev = ll.head
	ll.head.Next = ele
	ll.length++
}

func (ll *LinkList) Front() *Element {
	if ll.length == 0 {
		return nil
	}
	ele := ll.head.Next
	return ele
}

func (ll *LinkList) Back() *Element {
	if ll.length == 0 {
		return nil
	}
	ele := ll.tail.Prev
	return ele
}

func (ll *LinkList) PopFront() *Element {
	if ll.length == 0 {
		return nil
	}
	ele := ll.head
	ll.head = ll.head.Next
	return ele
}

func (ll *LinkList) Remove(n *Element) {
	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev
	ll.length--
}

func (ll *LinkList) MoveToFront(n *Element) {
	ll.Remove(n)
	ll.PushFrontElement(n)
}

func (ll *LinkList) Len() int {
	return ll.length
}

func (ll *LinkList) IsEmpty() bool {
	return ll.Len() == 0
}
