# Binary Search Strategy

A lot of problems can be generalized to binary search, where the
problem being solved is as follows:

> Minimize k such that Condition(k) is true. [1]

This approach requires applying binary search to a monotonically
increasing search space. The following is the most generalized form
of a binary search problem solution.

```
func BinarySearchSolution(array []int) {
    Condition := func(value int) bool {
        ...
    }

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
1. Correctly initializing the bounds `left` and `right` to specify
the search space. We only have to **initialize the bounds to include
all elements**.
2. Design the `Condition` function.
3. Decide the return value to be either `left` or `left - 1`. Note that
`left` satisfies `Condition` when exiting the `while` loop.

## Example: 69. Sqrt(x) [Easy]

> Implement `int sqrt(int x)`.
>
> Compute and return the square root of x, where x is guaranteed to be a non-negative integer.
>
> Since the return type is an integer, the decimal digits are truncated and only the integer part of the result is returned.

The first step is to properly initialize the left and right bounds. We know that `0^2` = `0`, so
we will initialize `left := 0`. For special cases, we set `right := x + 1`.

We'll then create the `Condition` function. To turn `sqrt` into a minimization function,
we frame the problem as follows:

> Minimize n such that n^2 > x.

Thus, the `n` that satisfies this requirement is the value after the square root
we are looking for, so we return `left - 1`.

```
func Sqrt(x int) {
    Condition := func(value int) bool {
        return value * value > x
    }

    left := 0
    right := x + 1
    
    while left < right {
        mid := (left + right) / 2
        if Condition(mid) {
            right = mid
        } else {
            left = mid + 1
        }
    }

    return left - 1
}
```

## Example: 875. Koko Eating Bananas [Medium]

> Koko loves to eat bananas. There are `N` piles of bananas,
> the i-th pile has `piles[i]` bananas.  The guards have gone
> and will come back in `H` hours.
>
> Koko can decide her bananas-per-hour eating speed of `K`.
> Each hour, she chooses some pile of bananas, and eats `K`
> bananas from that pile. If the pile has less than `K` bananas,
> she eats all of them instead, and won't eat any more bananas
> during this hour.
>
> Koko likes to eat slowly, but still wants to finish eating all
> the bananas before the guards come back. 
> Return the minimum integer `K` such that she can eat all the
> bananas within `H` hours.

First, we turn this into a monotonically increasing problem. As stated in
the question, we find the minimum number of bananas `k` to be consumed by
Koko per hour, with our `Condition(k)` function returning `true` if she can
eat all the bananas within `H` hours. Thus, the first `k` we find will be
the minimum number of bananas to eat per hour, and so we return `left`.

Our `Condition` function simply returns `true` if it takes more than `H` days
to consume all bananas.

```
import (
    "math"
)

func minEatingSpeed(piles []int, H int) int {
    Condition := func(k int) bool {
        totalHours := 0
        for _, pile := range piles {
            totalHours += int(math.Ceil(float64(pile) / float64(k)))
        }
        return totalHours <= H
    }
    
    left := 1
    right := max(piles)
    
    for left < right {
        mid := (left + right) / 2
        if Condition(mid) {
            right = mid
        } else {
            left = mid + 1
        }
    }
    
    return left
}

func max(nums []int) int {
    max := nums[0]
    for _, num := range nums {
        if num > max {
            max = num
        }
    }
    return max
}
```


- [1] https://leetcode.com/discuss/general-discussion/786126/python-powerful-ultimate-binary-search-template-solved-many-problems
