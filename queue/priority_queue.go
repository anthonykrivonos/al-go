package queue

import "github.com/anthonykrivonos/al-go/heap"

type priorityQueue struct {
	heap heap.Heap
}

func NewPriorityQueue(comparator func(a interface{}, b interface{}) int) Queue {
	q := &priorityQueue{}
	q.heap = heap.NewHeap(comparator)
	return q
}

func (p *priorityQueue) Unshift(x interface{}) {
	p.heap.Insert(x)
}

func (p *priorityQueue) Poll() interface{} {
	return p.heap.Pop()
}

func (p *priorityQueue) Check() interface{} {
	return p.heap.Root()
}

func (p *priorityQueue) Length() int {
	return p.heap.Length()
}

var _ Queue = &priorityQueue{}
