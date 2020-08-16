package heap

const (
	defaultHeapCapacity = 1024
)

type Heap interface {
	Insert(x interface{})
	Root() interface{}
	Pop() interface{}
	List() []interface{}
	Length() int
}

type heap struct {
	list []interface{}
	capacity int
	length int
	comparator func(a, b interface{}) int
}

func NewHeap(comparator func(a, b interface{}) int) Heap {
	h := &heap{}
	h.capacity = defaultHeapCapacity
	h.list = make([]interface{}, h.capacity)
	h.length = 0
	h.comparator = comparator
	return h
}

func parentIndex(i int) int {
	return (i - 1) / 2
}

func leftChildIndex(i int) int {
	return 2 * i + 1
}

func rightChildIndex(i int) int {
	return 2 * i + 2
}

func (h *heap) Insert(x interface{}) {
	// Index of new element
	i := h.length

	h.list[i] = x
	h.length++

	// Bubble up
	for i > 0 && h.comparator(x, h.list[parentIndex(i)]) < 0 {
		h.list[i] = h.list[parentIndex(i)]
		h.list[parentIndex(i)] = x
		i = parentIndex(i)
	}

	// If necessary, expand the heap capacity
	if h.length > h.capacity / 2 {
		h.capacity *= 2
		newList := make([]interface{}, h.capacity)
		for i := 0; i < h.length; i++ {
			newList[i] = h.list[i]
		}
		h.list = newList
	}
}

func (h *heap) Pop() interface{} {
	if h.length == 0 {
		return nil
	}

	// Store the root for later
	root := h.list[0]

	// Overwrite the root with the last element
	h.list[0] = h.list[h.length - 1]
	h.length--

	// Bubble down
	h.bubbleDown(0)

	return root
}

func (h *heap) bubbleDown(i int) {
	left := leftChildIndex(i)
	right := rightChildIndex(i)
	j := i

	if left < h.length && h.comparator(h.list[j], h.list[left]) >= 0 {
		j = left
	}

	if right < h.length && h.comparator(h.list[j], h.list[right]) >= 0 {
		j = right
	}

	if j != i {
		tmp := h.list[j]
		h.list[j] = h.list[i]
		h.list[i] = tmp
		h.bubbleDown(j)
	}
}

func (h *heap) Root() interface{} {
	return h.list[0]
}

func (h *heap) Length() int {
	return h.length
}

func (h *heap) List() []interface{} {
	return h.list[:h.length]
}
