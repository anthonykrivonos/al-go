package algorithms

import "github.com/anthonykrivonos/al-go/list"

// BinarySearch searches a sorted list in O(log(n)) time.
func BinarySearch(x interface{}, comparator func(a, b interface{}) int, list list.List) int {
	var low = 0
	var high = list.Length() - 1
	for low < high {
		mid := (high + low) / 2
		if comparator(list.Get(mid), x) < 0 {
			high = mid + 1
		} else if comparator(list.Get(mid), x) > 0 {
			low = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

// BinarySearchRecursive searches a sorted list in O(log(n)) time, recursively.
func BinarySearchRecursive(x interface{}, comparator func(a, b interface{}) int, list list.List) int {
	return binarySearchRecursive(x, comparator, list, 0, list.Length() - 1)
}

// binarySearchRecursive is a helper method for BinarySearchRecursive.
func binarySearchRecursive(x interface{}, comparator func(a, b interface{}) int, list list.List, low, high int) int {
	if low >= high {
		return -1
	}

	mid := (high + low) / 2
	if comparator(list.Get(mid), x) < 0 {
		return binarySearchRecursive(x, comparator, list, mid + 1, high)
	} else if comparator(list.Get(mid), x) > 0 {
		return binarySearchRecursive(x, comparator, list, low, mid - 1)
	}
	return mid
}
