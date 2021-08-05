// Package ring implements circular list of nodes with array backend.
package ring

import (
	"fmt"
	"sync"
	"unsafe"
)

// Ring represents a circular ring structure
// with nodes being maintained by an array of unsafe pointers.
type Ring struct {
	nodes    []unsafe.Pointer
	capacity int
	head     int
	depth    int
	sync.RWMutex
}

// New creates a new ring with defined capacity.
// The capacity must be at least 2.
func New(capacity int) *Ring {
	if capacity < 2 {
		panic(fmt.Errorf("ring capacity too low, expected at least 2"))
	}
	return &Ring{
		nodes:    make([]unsafe.Pointer, capacity),
		capacity: capacity,
		head:     0,
		depth:    0,
	}
}

// Add adds a new value to the ring structure.
func (r *Ring) Add(value unsafe.Pointer) {
	// set the write mode
	r.Lock()
	defer r.Unlock()

	// advance the head; loop back to zero if the top is reached
	r.head++
	if r.head == r.capacity {
		r.head = 0
	}

	// keep track of number of stored elements
	r.depth++
	if r.depth > r.capacity {
		r.depth = r.capacity
	}

	// store value
	r.nodes[r.head] = value
}

// Reset the ring depth.
func (r *Ring) Reset() {
	r.Lock()
	r.depth = 0
	r.Unlock()
}

// List returns a slice of at most depth items from the ring.
func (r *Ring) List(length int) []unsafe.Pointer {
	// handle the read mode
	r.RLock()
	defer r.RUnlock()

	// make sure to limit the pulled range to the ring depth
	tl := length
	if tl > r.depth {
		tl = r.depth
	}

	l := make([]unsafe.Pointer, tl)

	// load the data
	for i := 0; i < tl; i++ {
		// start from the current head and loop array indexes down
		// upper part is stitched to the bottom by adding capacity if needed
		ix := r.head - i
		if ix < 0 {
			ix += r.capacity
		}

		// if there is no content in the node, return all we have so far
		// since all the subsequent nodes will also be empty
		if nil == r.nodes[ix] {
			return l[:i]
		}
		l[i] = r.nodes[ix]
	}
	return l[:]
}
