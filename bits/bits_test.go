package bits

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToBitString(t *testing.T) {
	x := 256
	assert.Equal(t, "00000000000000000000000100000000", ToBitString(x))
}
