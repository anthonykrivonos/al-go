package list

type DoublyLinkedList interface {
	List
	Head() *DoublyLinkedListNode
	Insert(i int, x interface{})
	Remove(i int) interface{}
}

type DoublyLinkedListNode struct {
	Value interface{}
	Prev  *DoublyLinkedListNode
	Next  *DoublyLinkedListNode
}

type doublyLinkedList struct {
	head *DoublyLinkedListNode
	length int
}

func NewDoublyLinkedList(x interface{}) DoublyLinkedList {
	l := &doublyLinkedList{}
	if x != nil {
		l.head = &DoublyLinkedListNode{Value: x}
		l.length = 1
	}
	return l
}

// Head is O(1)
func (l *doublyLinkedList) Head() *DoublyLinkedListNode {
	return l.head
}

// Push is O(n)
func (l *doublyLinkedList) Push(x interface{}) {
	h := &DoublyLinkedListNode{Value: x}
	if l.length == 0 {
		l.head = h
	} else {
		n := l.head
		for n.Next != nil {
			n = n.Next
		}
		h.Prev = n
		n.Next = h
	}
	l.length++
}

// Unshift is O(1)
func (l *doublyLinkedList) Unshift(x interface{}) {
	n := &DoublyLinkedListNode{Value: x}
	if l.length > 0 {
		l.head.Prev = n
		n.Next = l.head
	}
	l.head = n
}

// Pop is O(n)
func (l *doublyLinkedList) Pop() interface{} {
	var temp *DoublyLinkedListNode
	if l.length == 1 {
		temp = l.head
		l.head = nil
		l.length--
		return temp.Value
	} else if l.length > 1 {
		n := l.head
		for n.Next.Next != nil {
			n = n.Next
		}
		temp = n.Next
		n.Next = nil
		l.length--
		return temp.Value
	}
	return nil
}

// Poll is O(1)
func (l* doublyLinkedList) Poll() interface{} {
	return l.Remove(0)
}

// Peek is O(1)
func (l *doublyLinkedList) Peek() interface{} {
	if l.length == 0 {
		return nil
	}
	return l.Get(l.length - 1)
}

// Check is O(1)
func (l *doublyLinkedList) Check() interface{} {
	if l.length == 0 {
		return nil
	}
	return l.Get(0)
}

// Set is O(n)
func (l *doublyLinkedList) Set(i int, x interface{}) {
	n := &DoublyLinkedListNode{Value: x}
	if i == 0 {
		n.Next = l.head.Next
		l.head = n
	} else if i < l.length && i > 0 {
		j := 0
		h := l.head
		for j < i {
			h = h.Next
			j++
		}
		h.Value = x
	} else if i == l.length {
		l.Push(x)
	}
}

// Get is O(n)
func (l *doublyLinkedList) Get(i int) interface{} {
	if i >= l.length || i < 0 {
		return nil
	}
	n := l.head
	for j := 0; j < i; j++ {
		n = n.Next
	}
	return n.Value
}

// Length is O(1)
func (l *doublyLinkedList) Length() int {
	return l.length
}

// Reverse is O(n)
func (l *doublyLinkedList) Reverse() {
	var prev *DoublyLinkedListNode
	curr := l.head
	for curr != nil {
		prev = curr.Prev
		curr.Prev = curr.Next
		curr.Next = prev
		curr = curr.Prev
	}
	if prev != nil && prev.Prev != nil {
		l.head = prev.Prev
	}
}

// Insert is O(n)
func (l *doublyLinkedList) Insert(i int, x interface{}) {
	n := &DoublyLinkedListNode{Value: x}
	if i == 0 {
		n.Next = l.head
		l.head.Prev = n
		l.head = n
		l.length++
	} else if i < l.length && i > 0 {
		h := l.head
		for j := 0; j < i - 1; j++ {
			h = h.Next
		}
		n.Next = h.Next
		n.Prev = h
		h.Next = n
		l.length++
	} else if i == l.length {
		l.Push(x)
	}
}

// Remove is O(n)
func (l* doublyLinkedList) Remove(i int) interface{} {
	if i < 0 || i >= l.length {
		return nil
	}
	var x *DoublyLinkedListNode
	if i == 0 {
		x = l.head
		l.head = x.Next
		l.head.Prev = nil
		l.length--
		return x.Value
	} else if i < l.length - 1 {
		j := 0
		n := l.head
		for j < i - 1 {
			n = n.Next
			j++
		}
		x := n.Next
		n.Next = n.Next.Next
		n.Next.Prev = x
		l.length--
		return x.Value
	} else {
		return l.Pop()
	}
}

var _ DoublyLinkedList = &doublyLinkedList{}
