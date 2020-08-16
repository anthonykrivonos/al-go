# Binary Search Strategy

A lot of problems can be generalized to binary search, where the
problem being solved is as follows:

> Minimize k such that Condition(k) is true. [1]

The following is the most generalized form of a binary search
problem solution.

```
func Condition(value int) bool {
    ...
}

func BinarySearchSolution(array []int) {
    searchSpace := // space we are searching within
    left := min(searchSpace)
    right := max(searchSpace)
    
    while left < right {
        mid := (left + right) / 2
        if Condition(mid) {
            right = mid
        } else {
            left = mid + 1
        }
    }

    return left
}
```

To use this template, we only need to concern ourselves with:
1. Corerctly initializing the bounds `left` and `right` to specify
the search space. We only have to **initalize the bounds to include
all elements**.
2. Decide the return value to be either `left` or `left - 1`. Note that
`left` satisfies `Condition` when exiting the `while` loop.
3. Design the `Condition` function.

- [1] https://leetcode.com/discuss/general-discussion/786126/python-powerful-ultimate-binary-search-template-solved-many-problems
