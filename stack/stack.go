package stack

import "github.com/anthonykrivonos/al-go/list"

type Stack interface {
	Push(x interface{})
	Pop() interface{}
	Peek() interface{}
	Length() int
}

var _ Stack = list.NewArrayList()
var _ Stack = list.NewSinglyLinkedList(nil)
var _ Stack = list.NewDoublyLinkedList(nil)
