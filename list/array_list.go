package list

const (
	defaultCapacity = 10
)

type ArrayList interface {
	List
}

type arrayList struct {
	items []interface{}
	length int
	capacity int
}

func NewArrayList() ArrayList {
	l := &arrayList{}
	l.items = make([]interface{}, defaultCapacity)
	l.length = 0
	l.capacity = defaultCapacity
	return l
}

// Push is amortized O(1), since it only runs in O(n) time when the ArrayList is doubled in size.
func (l *arrayList) Push(x interface{}) {
	if l.length == l.capacity {
		// Double the ArrayList capacity
		l.capacity = 2 * l.capacity
		newItems := make([]interface{}, l.capacity)
		for idx, item := range l.items {
			newItems[idx] = item
		}
		l.items = newItems
	}
	// Push the item to the end of the ArrayList
	l.items[l.length] = x
	l.length++
}

// Unshift is amortized O(1), since it only runs in O(n) time when the ArrayList is doubled in size.
func (l *arrayList) Unshift(x interface{}) {
	if l.length == l.capacity {
		// Double the ArrayList capacity
		l.capacity = 2 * l.capacity
		newItems := make([]interface{}, l.capacity)
		for idx, item := range l.items {
			newItems[idx] = item
		}
		l.items = newItems
	}
	// Add the item to the beginning of the ArrayList
	for i := l.length; i > 0; i-- {
		l.swap(i, i - 1)
	}
	l.items[0] = x
	l.length++
}

// Poll is O(n)
func (l* arrayList) Poll() interface{} {
	if l.length == 0 {
		return nil
	}
	h := l.items[0]
	for i := 0; i < l.length - 1; i++ {
		l.swap(i, i + 1)
	}
	l.items[l.length - 1] = nil
	l.length--
	return h
}

// Pop is O(1)
func (l *arrayList) Pop() interface{} {
	item := l.items[l.length - 1]
	l.items[l.length - 1] = nil
	l.length--
	return item
}

// Peek is O(1)
func (l *arrayList) Peek() interface{} {
	if l.length == 0 {
		return nil
	}
	return l.items[l.length - 1]
}

// Check is O(1)
func (l *arrayList) Check() interface{} {
	if l.length == 0 {
		return nil
	}
	return l.items[0]
}

// Set is O(1)
func (l *arrayList) Set(i int, x interface{}) {
	l.items[i] = x
}

// Get is O(1)
func (l *arrayList) Get(i int) interface{} {
	return l.items[i]
}

// Length is O(1)
func (l *arrayList) Length() int {
	return l.length
}

// swap is O(1)
func (l *arrayList) swap(i int, j int) {
	temp := l.items[i]
	l.items[i] = l.items[j]
	l.items[j] = temp
}

// Reverse is O(n)
func (l *arrayList) Reverse() {
	i := 0
	for i < l.Length()/2 {
		l.swap(i, l.Length() - 1 - i)
		i++
	}
}

var _ ArrayList = &arrayList{}
