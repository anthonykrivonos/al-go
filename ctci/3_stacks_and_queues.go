package ctci

import "fmt"

const (
	defaultThreeStackSize = 12
)


// 3.1 Three in One - O(n)
type ThreeInOne interface {
	Push(stack int, x interface{})
	Pop(stack int) interface{}
	Peek(stack int) interface{}
}
type threeInOne struct {
	capacity   int
	array      []interface{}
	botIndices []int
	curIndices []int
}
func NewThreeInOne() ThreeInOne {
	t := &threeInOne{}
	t.capacity = defaultThreeStackSize
	t.array = make([]interface{}, t.capacity)
	inc := t.capacity / 3
	t.botIndices = []int{0, inc, inc*2}
	t.curIndices = []int{0, inc, inc*2}
	return t
}
func (t *threeInOne) Push(stack int, x interface{}) {
	if stack < 1 || stack > 3 {
		panic("stack must be between 1 and 3, inclusive")
	}
	cur := t.curIndices[stack - 1]

	// Increase array size on overflow
	if (stack == 1 && cur == t.botIndices[1]) || (stack == 2 && cur == t.botIndices[2]) || (stack == 3 && cur == len(t.array)) {
		t.capacity *= 2
		newArray := make([]interface{}, t.capacity)
		inc := t.capacity / 3
		newBotIndices := []int{0, inc, inc*2}
		fmt.Printf("%d %d %d\n", newBotIndices[0], newBotIndices[1], newBotIndices[2])
		newCurIndices := []int{-1, -1, -1}
		// Copy stack 1 elements
		for i := 0; i < t.botIndices[1]; i++ {
			if t.array[i] == nil {
				newCurIndices[0] = i
				break
			}
			newArray[i] = t.array[i]
		}
		if newCurIndices[0] == -1 {
			newCurIndices[0] = newBotIndices[1]
		}
		// Copy stack 2 elements
		for i := t.botIndices[1]; i < t.botIndices[2]; i++ {
			if t.array[i] == nil {
				newCurIndices[1] = i - t.botIndices[1] + newBotIndices[1]
				break
			}
			newArray[i - t.botIndices[1] + newBotIndices[1]] = t.array[i]
		}
		if newCurIndices[1] == -1 {
			newCurIndices[1] = newBotIndices[2]
		}
		// Copy stack 3 elements
		for i := t.botIndices[2]; i < len(t.array); i++ {
			if t.array[i] == nil {
				newCurIndices[2] = i - t.botIndices[2] + newBotIndices[2]
				break
			}
			newArray[i - t.botIndices[2] + newBotIndices[2]] = t.array[i]
		}
		if newCurIndices[2] == -1 {
			newCurIndices[2] = len(t.array)
		}
		t.array = newArray
		t.botIndices = newBotIndices
		t.curIndices = newCurIndices
		for i := 0; i < len(t.array); i++ {
			fmt.Printf("%s ", t.array[i])
		}
	}
	t.array[cur] = x
	t.curIndices[stack - 1] += 1
}
func (t *threeInOne) Pop(stack int) interface{} {
	if stack < 1 || stack > 3 {
		panic("stack must be between 1 and 3, inclusive")
	}
	cur := t.curIndices[stack - 1]
	top := t.botIndices[stack - 1]
	if cur == top {
		return nil
	}
	t.curIndices[stack - 1] -= 1
	cur = t.curIndices[stack - 1]
	tmp := t.array[cur]
	t.array[cur] = nil
	return tmp
}
func (t *threeInOne) Peek(stack int) interface{} {
	if stack < 1 || stack > 3 {
		panic("stack must be between 1 and 3, inclusive")
	}
	cur := t.curIndices[stack - 1]
	top := t.botIndices[stack - 1]
	if cur == top {
		return nil
	}
	return t.array[cur - 1]
}

