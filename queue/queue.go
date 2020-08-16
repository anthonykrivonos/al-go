package queue

import "github.com/anthonykrivonos/al-go/list"

type Queue interface {
	Unshift(x interface{})
	Poll() interface{}
	Check() interface{}
	Length() int
}

var _ Queue = list.NewArrayList()
var _ Queue = list.NewSinglyLinkedList(nil)
var _ Queue = list.NewDoublyLinkedList(nil)
