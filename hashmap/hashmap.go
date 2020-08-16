package hashmap

import (
	"bytes"
	"encoding/gob"
)

const (
	defaultCapacity = 1217
	loadFactorThreshold = 1.1
)

type HashMap interface {
	Set(i interface{}, x interface {})
	Get(i interface{}) interface{}
	Has(i interface{}) bool
	Remove(i interface{})
}

type hashMap struct {
	buckets []*hashNode
	capacity int
	length int
}

type hashNode struct {
	key interface{}
	value interface{}
	next *hashNode
}

func NewHashMapWithCapacity(capacity int) HashMap {
	h := &hashMap{}
	h.buckets = make([]*hashNode, capacity)
	h.capacity = capacity
	h.length = 0
	return h
}

func NewHashMap() HashMap {
	return NewHashMapWithCapacity(defaultCapacity)
}

// Set is O(n).
func (h* hashMap) Set(i interface{}, x interface {}) {
	hashKey := hash(i) % h.capacity
	if h.buckets[hashKey] == nil {
		// Add the value to the map
		h.buckets[hashKey] = &hashNode{
			key:   i,
			value: x,
		}
		h.length++
	} else {
		// Attempt to find the next open space
		n := h.buckets[hashKey]
		if n.key == i {
			n.value = x
		} else {
			for n.next != nil && n.next.key != i {
				n = n.next
			}
			if n.next == nil {
				// Create a new key/value pair
				n.next = &hashNode{
					key:   i,
					value: x,
				}
				h.length++
			} else if n.next.key != i {
				// Overwrite existing key
				n.next.value = x
			}
		}
	}
	// Increase table size and rehash if necessary
	loadFactor := float64(h.length) / float64(h.capacity)
	if loadFactor > loadFactorThreshold {
		// Increase size of map
		// Add all nodes to a universal list
		allNodes := make([]*hashNode, 0)
		for _, b := range h.buckets {
			if b == nil {
				continue
			}
			n := b
			for n != nil {
				allNodes = append(allNodes, n)
				n = n.next
			}
		}
		// Resize the table
		h.capacity = eratosthenesLargestPrime(h.capacity * 2)
		h.buckets = make([]*hashNode, h.capacity)
		// Rehash the keys
		for _, n := range allNodes {
			h.Set(n.key, n.value)
		}
	}
}

// Get is O(n).
func (h* hashMap) Get(i interface{}) interface{} {
	hashKey := hash(i) % h.capacity
	n := h.buckets[hashKey]
	for n != nil && n.key != i {
		n = n.next
	}
	if n == nil {
		return nil
	}
	return n.value
}

// Get is O(n).
func (h* hashMap) Has(i interface{}) bool {
	return h.Get(i) != nil
}

// Remove is O(n).
func (h* hashMap) Remove(i interface{}) {
	hashKey := hash(i) % h.capacity
	n := h.buckets[hashKey]
	if n.key == i {
		h.buckets[hashKey] = n.next
	} else {
		for n.next != nil && n.next.key != i {
			n = n.next
		}
		if n.next != nil {
			n.next = n.next.next
		}
	}
}

func hash(i interface{}) int {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(i); err != nil {
		return 0
	}
	x := 0
	fac := 1
	for _, b := range buf.Bytes() {
		x += int(b) * fac
		fac *= 2
	}
	if x < 0 {
		x *= -1
	}
	return x
}

func eratosthenesLargestPrime(n int) int {
	primes := make([]int, 0)
	b := make([]bool, n)
	for i := 2; i < n; i++ {
		if b[i] == true { continue }
		primes = append(primes, i)
		for k := i * i; k < n; k += i {
			b[k] = true
		}
	}
	return primes[len(primes) - 1]
}
