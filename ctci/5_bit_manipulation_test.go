package ctci

import (
	"fmt"
	"github.com/anthonykrivonos/al-go/bits"
	"github.com/stretchr/testify/assert"
	"testing"
)

//  Insertion
func TestInsertion(t *testing.T) {
	m := 12
	n := 315

	i := 2
	j := 5

	assert.Equal(t, "1100", bits.ToBitString(Insertion(m, n, i, j))[26:30])
}

// 5.7 Pairwise Swap
func TestPairwiseSwap(t *testing.T) {
	x := 1221341234

	fmt.Println(bits.ToBitString(x))
	assert.Equal(t, "10000100110011000011000000110001", bits.ToBitString(PairwiseSwap(x)))
}
