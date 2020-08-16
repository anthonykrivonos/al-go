package ctci

import (
	list "github.com/anthonykrivonos/al-go/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 2.2 Return Kth to Last - O(n)
func TestReturnKthToLast(t *testing.T) {
	ll := list.NewSinglyLinkedList("this")
	ll.Push("is")
	ll.Push("sparta")
	ll.Push("xd")
	assert.Equal(t, "is", ReturnKthToLast(3, ll))
}

// 2.3 Delete Middle Node - O(n)
func TestDeleteMiddleNode(t *testing.T) {
	ll := list.NewSinglyLinkedList("this")
	ll.Push("is")
	ll.Push("sparta")
	ll.Push("xd")
	middle := ll.Head().Next.Next // "sparta"
	DeleteMiddleNode(middle)
	assert.Equal(t, "xd", ll.Get(2))
}

// 2.4 Partition - O(n^2)
func TestPartition(t *testing.T) {
	ll := list.NewSinglyLinkedList(3)
	ll.Push(5)
	ll.Push(8)
	ll.Push(5)
	ll.Push(10)
	ll.Push(2)
	ll.Push(1)
	Partition(5, ll)
}

// 2.5 Sum Lists - O(n)
func TestSumLists(t *testing.T) {
	// 110
	a1 := list.NewSinglyLinkedList(0)
	a1.Push(1)
	a1.Push(1)
	// 96
	a2 := list.NewSinglyLinkedList(6)
	a2.Push(9)
	// = 206
	a3 := SumLists(a1, a2)
	assert.Equal(t, 2, a3.Get(2))
	assert.Equal(t, 0, a3.Get(1))
	assert.Equal(t, 6, a3.Get(0))
}

// 2.5 Sum Lists: Suppose the digits are in forward order. - O(n)
func TestSumListsChallenge(t *testing.T) {
	// 110
	a1 := list.NewSinglyLinkedList(1)
	a1.Push(1)
	a1.Push(0)
	// 96
	a2 := list.NewSinglyLinkedList(9)
	a2.Push(6)
	// = 206
	a3 := SumListsChallenge(a1, a2)
	assert.Equal(t, 2, a3.Get(0))
	assert.Equal(t, 0, a3.Get(1))
	assert.Equal(t, 6, a3.Get(2))
}

// 2.6 Palindrome - O(n)
func TestPalindrome(t *testing.T) {
	p := list.NewSinglyLinkedList(nil)
	p.Push('r')
	p.Push('a')
	p.Push('c')
	p.Push('e')
	p.Push('c')
	p.Push('a')
	p.Push('r')

	f := list.NewSinglyLinkedList(nil)
	f.Push('r')
	f.Push('a')
	f.Push('c')
	f.Push('e')
	f.Push('c')
	f.Push('a')

	assert.True(t, Palindrome(p))
	assert.False(t, Palindrome(f))
}

// 2.7 Intersection - O(n)
func TestIntersection(t *testing.T) {
	a := list.NewSinglyLinkedList(0)
	a.Push(1)
	a.Push(2)
	a.Push(3)
	a.Push(4)

	b := list.NewSinglyLinkedList(3)

	assert.False(t, Intersection(a, b))

	b.Head().Next = a.Head().Next.Next.Next.Next

	assert.True(t, Intersection(a, b))
}

// 2.8 Loop Detection - O(n)
func TestLoopDetection(t *testing.T) {
	a := list.NewSinglyLinkedList(0)
	a.Push(1)
	a.Push(2)

	assert.False(t, LoopDetection(a))

	a.Head().Next.Next.Next = a.Head()

	assert.True(t, LoopDetection(a))
}
