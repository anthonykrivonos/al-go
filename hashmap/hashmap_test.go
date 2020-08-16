package hashmap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashMap(t *testing.T) {
	assert.Equal(t, hash("Three"), hash("Three"))
	assert.Equal(t, hash(3), hash(3))

	h := NewHashMapWithCapacity(5)

	h.Set("One", 1)
	h.Set("Two", 2)
	h.Set("Three", 3)
	h.Set("Four", 4)
	h.Set("Five", 5)

	assert.Equal(t, 1, h.Get("One"))
	assert.Equal(t, 2, h.Get("Two"))
	assert.Equal(t, 3, h.Get("Three"))
	assert.Equal(t, 4, h.Get("Four"))
	assert.Equal(t, 5, h.Get("Five"))

	h.Set("Three", 33)
	assert.Equal(t, 33, h.Get("Three"))

	// Resize the map
	h.Set("Three", 3)
	h.Set("Six", 6)
	h.Set("Seven", 7)
	assert.Equal(t, 1, h.Get("One"))
	assert.Equal(t, 2, h.Get("Two"))
	assert.Equal(t, 3, h.Get("Three"))
	assert.Equal(t, 4, h.Get("Four"))
	assert.Equal(t, 5, h.Get("Five"))
	assert.Equal(t, 6, h.Get("Six"))
	assert.Equal(t, 7, h.Get("Seven"))

	assert.Nil(t, h.Get("Eight"))

	h = NewHashMap()
	assert.Nil(t, h.Get("One"))
}
