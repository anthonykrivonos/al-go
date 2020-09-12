package bits

import (
	"strings"
)

// DivideInHalf uses an arithmetic right shift to move all bits right by one and then fill in the sign bit on the left.
func DivideInHalf(x int) int {
	return x >> 1
}

// MultiplyByTwo uses an arithmetic left shift to move all bits left by one and then fill in the sign bit on the left.
func MultiplyByTwo(x int) int {
	return x << 1
}

// GetBit returns the bit at the given place, starting from the LSB.
func GetBit(x, place int) bool {
	return (x & (1 << place)) != 0
}

// SetBit1 sets the bit at the given place, starting from the LSB. It uses bit shifting only.
func SetBit1(x, place int, bit bool) int {
	if bit {
		return x | (1 << place)
	}
	return x & ^(1 << place)
}

// SetBit2 sets the bit at the given place, starting from the LSB. It uses masking and bit shifting.
func SetBit2(x, place int, bit bool) int {
	val := 0
	if bit {
		val = 1
	}
	mask := ^(1 << place)
	return (x & mask) | (val << place)
}

// ClearAll sets all bits to 0, returning 0.
func ClearAll(x int) int {
	return x & 0
}

// ClearRight clears the ith bit from the right up until the LSB, inclusive.
func ClearRight(x, place int) int {
	mask := (-1 << place) - 1
	return mask & x
}

// ClearLeft clears the ith bit from the left up until the MSB, inclusive.
func ClearLeft(x, place int) int {
	mask := -1 << (place + 1)
	return mask & x
}

// ToBitString returns the integer as a binary bit string.
func ToBitString(m int) string {
	str := strings.Builder{}
	for i := 31; i >= 0; i-- {
		isZero := m & (1 << i) == 0
		if isZero {
			str.WriteRune('0')
		} else {
			str.WriteRune('1')
		}
	}
	return str.String()
}
