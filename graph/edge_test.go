package graph

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEdge(t *testing.T) {
	n1 := NewNode(1)
	n2 := NewNode(2)
	e := NewEdge(n1, n2, 0)

	assert.Equal(t, n1, e.From())
	assert.Equal(t, n2, e.To())
	assert.Equal(t, float64(0), e.Weight())
	assert.Equal(t, "1 => 2 (0.000000)", e.String())
}
