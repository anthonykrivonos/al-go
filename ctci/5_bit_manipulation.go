package ctci

import (
	"fmt"
	bits2 "github.com/anthonykrivonos/al-go/bits"
)

// 5.1 Insertion
func Insertion(m, n, i, j int) int {
	// Shift m to the starting position j
	m = m << i
	// Create a mask that clears from the left
	clear := -1 & ^(-1 << (j + i - 1))
	// Clear the mask from the right too, then negate
	mask := ^(clear ^ (-1 & ^(-1 << (i))))
	// Now, simply clear the necessary bits and or them with m
	insertion := n & mask | m
	return insertion
}

// 5.3 Flip Bit to Win
func FlipBitToWin(x int) int {
	bits := []rune(bits2.ToBitString(x))

	flipped := -1
	count := 0
	maxCount := 0
	for i := 0; i < len(bits); i++ {
		bit := bits[i]
		if bit == '0' {
			if flipped == -1 {
				flipped = i
				count++
			} else {
				i = flipped
				flipped = -1
				if count > maxCount {
					maxCount = count
				}
			}
		} else {
			count++
		}
	}
	if count > maxCount {
		maxCount = count
	}

	// Return the longest string of ones
	return count
}

// 5.4 Next Number
func NextNumber(x int) (int, int) {
	nextLargest := x << 1
	nextSmallest := x >> 1
	return nextLargest, nextSmallest
}

// 5.6 Conversion
func Conversion(from, to int) int {
	// XOR from and to to get all bit differences
	xor := from ^ to
	// Count the bit differences
	xorStr := []rune(bits2.ToBitString(xor))
	count := 0
	for _, bit := range xorStr {
		if bit == '1' {
			count++
		}
	}
	// Returns the number of bits you need to flip to convert `from` to `to`
	return count
}

// 5.7 Pairwise Swap
func PairwiseSwap(x int) int {
	evenMask := -1431655766
	oddMask := 1431655765

	fmt.Println("even", bits2.ToBitString(evenMask))
	fmt.Println("odd ", bits2.ToBitString(oddMask))

	newOdd := (oddMask & x) << 1
	newEven := (evenMask & x) >> 1

	// Returns a new integer where even and odd bits are swapped
	return newOdd | newEven
}
