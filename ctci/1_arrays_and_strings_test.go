package ctci

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 1.1 Is Unique - O(n)
func TestIsUnique(t *testing.T) {
	assert.True(t, IsUnique("abcdefghijklmnopqrstuvwxyz"))
	assert.False(t, IsUnique("abcabcabcabc"))
	assert.True(t, IsUnique(""))
}

// 1.1 Is Unique: What if you cannot use additional data structures? - O(n^2)
func TestIsUniqueChallenge(t *testing.T) {
	assert.True(t, IsUniqueChallenge("abcdefghijklmnopqrstuvwxyz"))
	assert.False(t, IsUniqueChallenge("abcabcabcabc"))
	assert.True(t, IsUniqueChallenge(""))
}

// 1.2 Check Permutation - O(n)
func TestCheckPermutation(t *testing.T) {
	passCases := [][]string{{"abcd", "dabc"}, {"abccd", "cdabc"}, {" ", " "}, {"", ""}}
	failCases := [][]string{{"abcd", "dabcg"}, {"abcd", "abce"}, {" ", ""}}
	for _, passCase := range passCases {
		assert.True(t, ChackPermutation(passCase[0], passCase[1]))
	}
	for _, failCase := range failCases {
		assert.False(t, ChackPermutation(failCase[0], failCase[1]))
	}
}

// 1.3 URLify - O(n)
func TestURLify(t *testing.T) {
	assert.Equal(t, "Mr%20John%20Smith", URLify("Mr John Smith    "))
	assert.Equal(t, "", URLify(""))
	assert.Equal(t, "", URLify("             "))
}

// 1.4 Palindrome Permutation - O(n)
// Uses only lowercase inputs.
func TestPalindromePermutation(t *testing.T) {
	assert.True(t, PalindromePermutation("tact coa"))
	assert.True(t, PalindromePermutation("patpat"))
	assert.True(t, PalindromePermutation("aaa"))
	assert.True(t, PalindromePermutation(" "))
	assert.True(t, PalindromePermutation(""))
	assert.False(t, PalindromePermutation("pattern pat"))
}

// 1.5 One Away - O(n)
func TestOneAway(t *testing.T) {
	assert.True(t, OneAway("palp", "plp"))
	assert.True(t, OneAway("palps", "palp"))
	assert.True(t, OneAway("palp", "balp"))
	assert.False(t, OneAway("palps", "balp"))
	assert.False(t, OneAway("palpop", "balp"))
	assert.False(t, OneAway("palpo", "balpso"))
}


// 1.6 String Compression - O(n)
func TestStringCompression(t *testing.T) {
	assert.Equal(t, "a2b1c5a3", StringCompression("aabcccccaaa"))
	assert.Equal(t, "abcdd", StringCompression("abcdd"))
	assert.Equal(t, "aabbccddeeffgg", StringCompression("aabbccddeeffgg"))
	assert.Equal(t, "", StringCompression(""))
}

// 1.7 RotateMatrix - O(n)
func TestRotateMatrix(t *testing.T) {
	assert.Equal(t, [][]int{{}}, RotateMatrix([][]int{{}}))
	assert.Equal(t, [][]int{{1}}, RotateMatrix([][]int{{1}}))
	assert.Equal(t, [][]int{{3, 1}, {4, 2}}, RotateMatrix([][]int{{1, 2}, {3, 4}}))
	assert.Equal(t, [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}, RotateMatrix([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}

// 1.7 Rotate Matrix: Can you do this in place? - O(n)
func TestRotateMatrixChallenge(t *testing.T) {
	emptyTest := [][]int{{}}
	RotateMatrixChallenge(&emptyTest)
	assert.Equal(t, [][]int{{}}, emptyTest)

	unaryTest := [][]int{{1}}
	RotateMatrixChallenge(&unaryTest)
	assert.Equal(t, [][]int{{1}}, unaryTest)

	twoByTwoTest := [][]int{{1, 2}, {3, 4}}
	RotateMatrixChallenge(&twoByTwoTest)
	assert.Equal(t, [][]int{{3, 1}, {4, 2}}, twoByTwoTest)

	threeByThreeTest := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	RotateMatrixChallenge(&threeByThreeTest)
	assert.Equal(t, [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}, threeByThreeTest)
}

// 1.8 Zero Matrix - O(n)
func TestZeroMatrix(t *testing.T) {

}
