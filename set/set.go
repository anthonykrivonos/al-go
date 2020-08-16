package set

import (
	"github.com/anthonykrivonos/al-go/hashmap"
	"github.com/anthonykrivonos/al-go/list"
)

type Set interface {
	Insert(x interface{})
	Remove(x interface{})
	Has(x interface{}) bool
	Array() []interface{}
	Clear()
}

type set struct {
	pos hashmap.HashMap
	keys list.DoublyLinkedList
}

func NewSet() Set {
	set := &set{}
	set.pos = hashmap.NewHashMap()
	set.keys = list.NewDoublyLinkedList(nil)
	return set
}

// Insert is amortized O(n) due to DoublyLinkedList.Push().
func (s *set) Insert(x interface{}) {
	if !s.pos.Has(x) {
		s.keys.Push(x)
		index := s.keys.Length() - 1
		s.pos.Set(x, index)
	}
}

// Remove is O(n).
func (s *set) Remove(x interface{}) {
	if s.pos.Has(x) {
		index := s.pos.Get(x).(int)
		s.keys.Remove(index)
		s.pos.Remove(x)
	}
}

// Has is O(n).
func (s *set) Has(x interface{}) bool {
	return s.pos.Has(x)
}

// Array is O(n^2).
func (s *set) Array() []interface{} {
	arr := make([]interface{}, 0)
	for i := 0; i < s.keys.Length(); i++ {
		key := s.keys.Get(i).(int)
		arr = append(arr, key)
	}
	return arr
}

// Clear is O(1).
func (s *set) Clear() {
	s.keys = list.NewDoublyLinkedList(nil)
	s.pos = hashmap.NewHashMap()
}

var _ Set = &set{}
