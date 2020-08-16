package ctci

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThreeInOne(t *testing.T) {
	tio := NewThreeInOne()

	assert.Nil(t, tio.Pop(1))
	assert.Nil(t, tio.Peek(1))

	tio.Push(1, "Good")
	tio.Push(1, "job")
	tio.Push(2, "Great")
	tio.Push(2, "work")
	tio.Push(3, "Excellent")
	tio.Push(3, "achievement")

	assert.Equal(t, "job", tio.Peek(1))
	assert.Equal(t, "job", tio.Pop(1))
	assert.Equal(t, "Good", tio.Pop(1))

	assert.Equal(t, "work", tio.Peek(2))
	assert.Equal(t, "work", tio.Pop(2))
	assert.Equal(t, "Great", tio.Pop(2))

	assert.Equal(t, "achievement", tio.Peek(3))
	assert.Equal(t, "achievement", tio.Pop(3))
	assert.Equal(t, "Excellent", tio.Pop(3))
}
