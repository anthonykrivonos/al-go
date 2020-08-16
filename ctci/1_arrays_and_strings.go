package ctci

import (
	"fmt"
	"github.com/anthonykrivonos/al-go/hashmap"
	"github.com/anthonykrivonos/al-go/list"
	"github.com/anthonykrivonos/al-go/set"
	"strings"
)

// 1.1 Is Unique - O(n)
func IsUnique(input string) bool {
	charMap := hashmap.NewHashMap()
	for _, c := range input {
		if charMap.Has(c) {
			return false
		}
		charMap.Set(c, 1)
	}
	return true
}

// 1.1 Is Unique: What if you cannot use additional data structures? - O(n^2)
func IsUniqueChallenge(input string) bool {
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			if input[i] == input[j] {
				return false
			}
		}
	}
	return true
}

// 1.2 Check Permutation - O(n)
func ChackPermutation(base, perm string) bool {
	if len(base) != len(perm) {
		return false
	}

	// Store base rune counts in map
	baseMap := hashmap.NewHashMap()
	for _, c := range base {
		if count := baseMap.Get(c); count != nil {
			newCount := count.(int) + 1
			baseMap.Set(c, newCount)
		} else {
			baseMap.Set(c, 1)
		}
	}

	// Store perm rune counts in map
	permList := list.NewArrayList()
	permMap := hashmap.NewHashMap()
	for _, c := range perm {
		permList.Push(c)
		if count := permMap.Get(c); count != nil {
			newCount := count.(int) + 1
			permMap.Set(c, newCount)
		} else {
			permMap.Set(c, 1)
		}
	}

	// Check that rune counts equal
	for i := 0; i < permList.Length(); i++ {
		c := permList.Get(i).(rune)
		if baseMap.Get(c) != permMap.Get(c) {
			return false
		}
	}
	return true
}

// 1.3 URLify - O(n)
func URLify(input string) string {
	// Create the URL string into an ArrayList
	urlList := list.NewArrayList()
	didEncounterSpace := false
	for _, c := range input {
		if c != ' ' {
			if didEncounterSpace {
				urlList.Push('%')
				urlList.Push('2')
				urlList.Push('0')
				didEncounterSpace = false
			}
			urlList.Push(c)
		} else if !didEncounterSpace {
			didEncounterSpace = true
		}
	}
	// Join the string back together
	s := strings.Builder{}
	for i := 0; i < urlList.Length(); i++ {
		s.WriteRune(urlList.Get(i).(rune))
	}
	return s.String()
}

// 1.4 Palindrome Permutation - O(n)
// Uses only lowercase inputs.
func PalindromePermutation(input string) bool {
	charList := list.NewArrayList()
	charMap := hashmap.NewHashMap()
	for _, c := range input {
		if !charMap.Has(c) {
			charList.Push(c)
			charMap.Set(c, 1)
		} else {
			newCount := charMap.Get(c).(int) + 1
			charMap.Set(c, newCount)
		}
	}
	oddCount := 0
	for i := 0; i < charList.Length(); i++ {
		c := charList.Get(i).(rune)
		cCount := charMap.Get(c).(int)
		if c != ' ' && cCount % 2 != 0 {
			oddCount++
			if oddCount > 1 {
				return false
			}
		}
	}
	return true
}

// 1.5 One Away - O(n)
func OneAway(base, modified string) bool {
	// Ensure sizes are at least 1 away
	if len(base) - len(modified) > 1 {
		return false
	}

	// Counters
	i := 0
	j := 0

	// Traverse each string
	didEdit := false
	for i < len(base) && j < len(modified) {
		if base[i] != modified[j] {
			if didEdit {
				return false
			}
			didEdit = true
			if len(base) > len(modified) {
				j--
			} else if len(modified) > len(base) {
				i--
			}
		}
		i++
		j++
	}
	return true
}

// 1.6 String Compression - O(n)
func StringCompression(input string) string {
	// Ensure an input string is provided
	if len(input) == 0 {
		return input
	}

	// Count the last char and the number of times it repeats
	lastChar := []rune(input)[0]
	repCount := 1

	// Build a new compressed string
	compressed := strings.Builder{}

	// Compress the string
	for i := 1; i < len(input); i++ {
		c := []rune(input)[i]
		if lastChar != c {
			compressed.WriteRune(lastChar)
			compressed.WriteString(fmt.Sprint(repCount))
			lastChar = c
			repCount = 1
		} else {
			repCount++
		}
	}

	// Write the last char to the compressed string
	compressed.WriteRune(lastChar)
	compressed.WriteString(fmt.Sprint(repCount))

	// "If the "compressed" string would not become smaller than the original string, your
	//  method should return the original string."
	if len(input) <= compressed.Len() {
		return input
	}

	return compressed.String()
}

// 1.7 Rotate Matrix - O(n)
func RotateMatrix(image [][]int) [][]int {
	if len(image) == 0 {
		return image
	}

	// Store length and width
	r := len(image)
	c := len(image[0])

	// Create empty array
	rotated := make([][]int, r)
	for i := range rotated {
		rotated[i] = make([]int, c)
	}

	// Perform rotation
	for i := r - 1; i >= 0; i-- {
		for j := 0; j < c; j++ {
			iNew := j
			jNew := r - 1 - i
			rotated[iNew][jNew] = image[i][j]
		}
	}

	return rotated
}

// 1.7 Rotate Matrix: Can you do this in place? - O(n)
func RotateMatrixChallenge(image *[][]int) {
	if image == nil || len(*image) == 0 {
		return
	}

	// Store length and width
	r := len(*image)
	c := len((*image)[0])

	// Transpose the matrix
	for i := 0; i < r; i++ {
		for j := 0; j < i; j++ {
			if i != j {
				// Swap across diagonal
				temp := (*image)[i][j]
				(*image)[i][j] = (*image)[j][i]
				(*image)[j][i] = temp
			}
		}
	}

	// Reverse the columns
	for i := 0; i < r; i++ {
		for j := 0; j < c / 2; j++ {
			temp := (*image)[i][j]
			(*image)[i][j] = (*image)[i][c - j - 1]
			(*image)[i][c - j - 1] = temp
		}
	}
}

// 1.8 Zero Matrix - O(n)
func ZeroMatrix(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return matrix
	}

	// Store length and width
	r := len(matrix)
	c := len((matrix)[0])

	// Create empty array
	zero := make([][]int, r)
	for i := range zero {
		zero[i] = make([]int, c)
	}

	// Store zero positions
	zeroRow := set.NewSet()
	zeroCol := set.NewSet()
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if matrix[i][j] == 0 {
				zeroRow.Insert(i)
				zeroCol.Insert(j)
			}
		}
	}

	// Write nonzero values
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if !zeroRow.Has(i) && !zeroCol.Has(j) {
				zero[i][j] = matrix[i][j]
			}
		}
	}

	return zero
}
