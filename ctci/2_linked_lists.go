package ctci

import (
	"fmt"
	"github.com/anthonykrivonos/al-go/list"
	"github.com/anthonykrivonos/al-go/queue"
	"github.com/anthonykrivonos/al-go/set"
	"github.com/anthonykrivonos/al-go/stack"
)

// 2.2 Return Kth to Last - O(n)
func ReturnKthToLast(k int, input list.SinglyLinkedList) interface{} {
	var q queue.Queue = list.NewArrayList()
	n := input.Head()
	for n != nil {
		q.Unshift(n)
		if q.Length() >= k {
			q.Poll()
		}
		n = n.Next
	}
	kThNode := q.Poll().(*list.SinglyLinkedListNode)
	return kThNode.Value
}

// 2.3 Delete Middle Node - O(n)
func DeleteMiddleNode(middle *list.SinglyLinkedListNode) {
	n := middle
	for n.Next != nil {
		n.Value = n.Next.Value
		n = n.Next
	}
	n = nil
}

// 2.4 Partition - O(n^2)
func Partition(x int, input list.SinglyLinkedList) {
	var less *list.SinglyLinkedListNode
	if input.Head().Value.(int) < x {
		less = input.Head()
	}
	n := input.Head()
	for n.Next != nil {
		if n.Next.Value.(int) < x {
			if less == nil {
				input.Insert(0, n.Next.Value)
				n.Next = n.Next.Next
				less = input.Head()
			} else {
				tmp := n.Next
				n.Next = n.Next.Next
				tmp.Next = less.Next
				less.Next = tmp
			}
		} else {
			n = n.Next
		}
	}
}

// 2.5 Sum Lists - O(n)
func SumLists(a, b list.SinglyLinkedList) list.SinglyLinkedList {
	sum := list.NewSinglyLinkedList(nil)
	carry := 0
	aHead := a.Head()
	bHead := b.Head()
	for aHead != nil && bHead != nil {
		num := aHead.Value.(int) + bHead.Value.(int) + carry
		if num > 9 {
			carry = 1
			sum.Push(num % 10)
		} else {
			carry = 0
			sum.Push(num)
		}
		aHead = aHead.Next
		bHead = bHead.Next
	}
	for aHead != nil {
		num := aHead.Value.(int) + carry
		if num > 9 {
			carry = 1
			sum.Push(num % 10)
		} else {
			carry = 0
			sum.Push(num)
		}
		aHead = aHead.Next
	}
	for bHead != nil {
		num := bHead.Value.(int) + carry
		if num > 9 {
			carry = 1
			sum.Push(num % 10)
		} else {
			carry = 0
			sum.Push(num)
		}
		bHead = bHead.Next
	}
	return sum
}

// 2.5 Sum Lists: Suppose the digits are in forward order. - O(n)
func SumListsChallenge(a, b list.SinglyLinkedList) list.SinglyLinkedList {
	sum := list.NewSinglyLinkedList(nil)

	var aStack stack.Stack = list.NewArrayList()
	var bStack stack.Stack = list.NewArrayList()

	// Push values into stack
	nA := a.Head()
	nB := b.Head()
	for nA != nil {
		aStack.Push(nA.Value)
		nA = nA.Next
	}
	for nB != nil {
		bStack.Push(nB.Value)
		nB = nB.Next
	}

	carry := 0
	for aStack.Length() > 0 && bStack.Length() > 0 {
		num := aStack.Pop().(int) + bStack.Pop().(int) + carry
		if num > 9 {
			carry = 1
			sum.Insert(0, num % 10)
		} else {
			carry = 0
			sum.Insert(0, num)
		}
	}

	for aStack.Length() > 0 {
		num := aStack.Pop().(int) + carry
		if num > 9 {
			carry = 1
			sum.Insert(0, num % 10)
		} else {
			carry = 0
			sum.Insert(0, num)
		}
	}
	for bStack.Length() > 0 {
		num := bStack.Pop().(int) + carry
		if num > 9 {
			carry = 1
			sum.Insert(0, num % 10)
		} else {
			carry = 0
			sum.Insert(0, num)
		}
	}
	return sum
}

// 2.6 Palindrome - O(n)
func Palindrome(input list.SinglyLinkedList) bool {
	reversed := list.NewSinglyLinkedList(nil)
	n := input.Head()
	for n != nil {
		reversed.Insert(0, n.Value)
		n = n.Next
	}

	n = reversed.Head()
	p := input.Head()
	for n != nil && p != nil && n.Value == p.Value {
		n = n.Next
		p = p.Next
	}

	return n == nil && p == nil
	//var prev, curr, next *list.SinglyLinkedListNode
	//curr = input.Head()
	//for curr != nil {
	//	next = curr.Next
	//	curr.Next = prev
	//	prev = curr
	//	curr = next
	//}
}

// 2.7 Intersection - O(n)
func Intersection(a, b list.SinglyLinkedList) bool {
	visited := set.NewSet()

	nA := a.Head()
	nB := b.Head()
	for nA != nil {
		visited.Insert(nA)
		nA = nA.Next
	}
	for nB != nil {
		if visited.Has(nB) {
			return true
		}
		nB = nB.Next
	}

	return false
}

// 2.8 Loop Detection - O(n)
func LoopDetection(input list.SinglyLinkedList) bool {
	visited := set.NewSet()

	n := input.Head()
	for n != nil {
		key := fmt.Sprintf("%p", n)
		if visited.Has(key) {
			return true
		}
		visited.Insert(key)
		n = n.Next
	}
	return false
}
