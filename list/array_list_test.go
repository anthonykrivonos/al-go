package list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayList(t *testing.T) {
	a := NewArrayList()

	assert.Nil(t, a.Poll())
	assert.Nil(t, a.Peek())
	assert.Nil(t, a.Check())

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
	assert.Equal(t, 0, a.Peek())
	assert.Equal(t, 9, a.Check())

	assert.Equal(t, 9, a.Poll())

	a.Unshift(-1)
	assert.Equal(t, -1, a.Get(0))

	b := NewArrayList()
	b.Unshift(1)
	b.Unshift(2)
	b.Unshift(3)
	b.Unshift(4)
	b.Unshift(5)
	b.Unshift(6)
	b.Unshift(7)
	b.Unshift(8)
	b.Unshift(9)
	b.Unshift(10)
	b.Unshift(10)
	b.Unshift(10)

	b.Reverse()
	assert.Equal(t, 1, b.Get(0))
}
