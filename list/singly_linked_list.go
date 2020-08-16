package list

type SinglyLinkedList interface {
	List
	Head() *SinglyLinkedListNode
	Insert(i int, x interface{})
	Remove(i int) interface{}
}

type SinglyLinkedListNode struct {
	Value interface{}
	Next  *SinglyLinkedListNode
}

type singlyLinkedList struct {
	head *SinglyLinkedListNode
	length int
}

func NewSinglyLinkedList(x interface{}) SinglyLinkedList {
	l := &singlyLinkedList{}
	if x != nil {
		l.head = &SinglyLinkedListNode{Value: x}
		l.length = 1
	}
	return l
}

// Head is O(1)
func (l *singlyLinkedList) Head() *SinglyLinkedListNode {
	return l.head
}

// Push is O(n)
func (l *singlyLinkedList) Push(x interface{}) {
	h := &SinglyLinkedListNode{Value: x}
	if l.length == 0 {
		l.head = h
	} else {
		n := l.head
		for n.Next != nil {
			n = n.Next
		}
		n.Next = h
	}
	l.length++
}

// Unshift is O(1)
func (l *singlyLinkedList) Unshift(x interface{}) {
	n := &SinglyLinkedListNode{Value: x}
	if l.length > 0 {
		n.Next = l.head
	}
	l.head = n
}

// Pop is O(n)
func (l *singlyLinkedList) Pop() interface{} {
	var temp *SinglyLinkedListNode
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
func (l* singlyLinkedList) Poll() interface{} {
	return l.Remove(0)
}

// Peek is O(1)
func (l *singlyLinkedList) Peek() interface{} {
	if l.length == 0 {
		return nil
	}
	return l.Get(l.length - 1)
}

// Check is O(1)
func (l *singlyLinkedList) Check() interface{} {
	if l.length == 0 {
		return nil
	}
	return l.Get(0)
}

// Set is O(n)
func (l *singlyLinkedList) Set(i int, x interface{}) {
	n := &SinglyLinkedListNode{Value: x}
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
func (l *singlyLinkedList) Get(i int) interface{} {
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
func (l *singlyLinkedList) Length() int {
	return l.length
}

// Reverse is O(n)
func (l *singlyLinkedList) Reverse() {
	var prev *SinglyLinkedListNode
	var next *SinglyLinkedListNode
	curr := l.head
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	l.head = prev
}

// Insert is O(n)
func (l *singlyLinkedList) Insert(i int, x interface{}) {
	n := &SinglyLinkedListNode{Value: x}
	if i == 0 {
		n.Next = l.head
		l.head = n
		l.length++
	} else if i < l.length && i > 0 {
		h := l.head
		for j := 0; j < i - 1; j++ {
			h = h.Next
		}
		n.Next = h.Next
		h.Next = n
		l.length++
	} else if i == l.length {
		l.Push(x)
	}
}

// Remove is O(n)
func (l* singlyLinkedList) Remove(i int) interface{} {
	if i < 0 || i >= l.length {
		return nil
	}
	var x *SinglyLinkedListNode
	if i == 0 {
		x = l.head
		l.head = x.Next
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
		l.length--
		return x.Value
	} else {
		return l.Pop()
	}
}

var _ SinglyLinkedList = &singlyLinkedList{}
