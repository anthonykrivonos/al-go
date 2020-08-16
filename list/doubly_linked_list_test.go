package list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoublyLinkedList(t *testing.T) {
	a := NewDoublyLinkedList(nil)

	assert.Nil(t, a.Peek())
	assert.Nil(t, a.Check())

	a = NewDoublyLinkedList(1)
	assert.Equal(t, 1, a.Length())

	assert.Equal(t, 1, a.Length())

	a.Pop()
	assert.Equal(t, 0, a.Length())

	a.Push(1)
	a.Push(2)
	a.Push(3)
	a.Push(4)
	a.Push(5)
	a.Push(6)
	a.Push(7)
	a.Push(8)
	a.Push(9)

	a.Reverse()

	assert.Equal(t, 9, a.Length())
	assert.Equal(t, 9, a.Get(0))
	assert.Equal(t, 8, a.Get(1))
	assert.Equal(t, 7, a.Get(2))
	assert.Equal(t, 6, a.Get(3))
	assert.Equal(t, 5, a.Get(4))
	assert.Equal(t, 4, a.Get(5))
	assert.Equal(t, 3, a.Get(6))
	assert.Equal(t, 2, a.Get(7))
	assert.Equal(t, 1, a.Get(8))

	last := a.Pop()
	assert.Equal(t, 8, a.Length())
	assert.Equal(t, 1, last)

	a.Push(1)
	a.Push(0)
	a.Push(-1)
	assert.Equal(t, 11, a.Length())
	assert.Equal(t, -1, a.Get(a.Length() - 1))

	a.Set(a.Length() - 1, 0)
	assert.Equal(t, 0, a.Get(a.Length() - 1))

	for i := 0; i < 11; i++ {
		a.Pop()
	}
	assert.Equal(t, 0, a.Length())
	assert.Nil(t, a.Pop())

	a.Push(1)
	a.Set(0, 2)
	a.Set(1, 3)
	assert.Equal(t, 2, a.Get(0))
	assert.Equal(t, 3, a.Get(1))
	assert.Nil(t, a.Get(-1))

	a.Insert(0, 1)
	assert.Equal(t, 1, a.Get(0))
	assert.Equal(t, 2, a.Get(1))

	a.Push(5)
	a.Insert(3, 4)
	assert.Equal(t, 4, a.Get(3))

	a.Insert(5, 6)
	assert.Equal(t, 6, a.Length())
	assert.Equal(t, 6, a.Get(5))

	a.Remove(0)
	assert.Equal(t, 2, a.Get(0))

	removed := a.Remove(4)
	assert.Equal(t, 6, removed)

	a.Remove(2)
	assert.Equal(t, 5, a.Get(2))

	assert.Nil(t, a.Remove(-1))

	assert.Equal(t, 5, a.Peek())
	assert.Equal(t, 2, a.Check())
	assert.Equal(t, 2, a.Poll())

	a.Unshift(-1)
	assert.Equal(t, -1, a.Get(0))
}
