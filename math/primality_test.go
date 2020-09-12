package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPrimeNaive(t *testing.T) {
	assert.Equal(t, false, IsPrimeNaive(0))
	assert.Equal(t, false, IsPrimeNaive(1))
	assert.Equal(t, true, IsPrimeNaive(2))
	assert.Equal(t, true, IsPrimeNaive(7))
	assert.Equal(t, true, IsPrimeNaive(113))
	assert.Equal(t, false, IsPrimeNaive(114))
	assert.Equal(t, false, IsPrimeNaive(12315))
}

func TestIsPrimeSieveOfEratosthenes(t *testing.T) {
	assert.Equal(t, false, IsPrimeSieveOfEratosthenes(0))
	assert.Equal(t, false, IsPrimeSieveOfEratosthenes(1))
	assert.Equal(t, true, IsPrimeSieveOfEratosthenes(2))
	assert.Equal(t, true, IsPrimeSieveOfEratosthenes(7))
	assert.Equal(t, true, IsPrimeSieveOfEratosthenes(113))
	assert.Equal(t, false, IsPrimeSieveOfEratosthenes(114))
	assert.Equal(t, false, IsPrimeSieveOfEratosthenes(12315))
}
