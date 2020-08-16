package heap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHeap(t *testing.T) {
	// Max heap
	maxComp := func (a, b interface{}) int {
		return b.(int) - a.(int)
	}
	maxHeap := NewHeap(maxComp)

	maxHeap.Insert(10)
	maxHeap.Insert(24)
	maxHeap.Insert(50)

	assert.Equal(t, 50, maxHeap.Pop())
	assert.Equal(t, 24, maxHeap.Pop())
	assert.Equal(t, 10, maxHeap.Pop())

	// Min heap
	minComp := func (a, b interface{}) int {
		return a.(int) - b.(int)
	}
	minHeap := NewHeap(minComp)

	minHeap.Insert(10)
	minHeap.Insert(24)
	minHeap.Insert(50)

	assert.Equal(t, 10, minHeap.Pop())
	assert.Equal(t, 24, minHeap.Pop())
	assert.Equal(t, 50, minHeap.Pop())
}
