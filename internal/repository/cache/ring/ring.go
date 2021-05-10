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
	nodes  []unsafe.Pointer
	length int
	head   int
	sync.RWMutex
}

// New creates a new ring with defined capacity.
// The capacity must be at least 2.
func New(capacity int) *Ring {
	if capacity < 2 {
		panic(fmt.Errorf("ring capacity too low, expected at least 2"))
	}
	return &Ring{
		nodes:  make([]unsafe.Pointer, capacity),
		length: capacity,
		head:   0,
	}
}

// Add adds a new value to the ring structure.
func (r *Ring) Add(value unsafe.Pointer) {
	// set the write mode
	r.Lock()
	defer r.Unlock()

	// advance the head
	r.head++
	if r.head == r.length {
		r.head = 0
	}

	// store value
	r.nodes[r.head] = value
}

// List returns a slice of at most length-1 items from the ring.
func (r *Ring) List(length int) []unsafe.Pointer {
	// how many we can pull? make sure to stop at the ring capacity
	tl := length
	if length > r.length {
		length = r.length
	}

	// make the container
	l := make([]unsafe.Pointer, tl)

	// set the read mode
	r.RLock()
	defer r.RUnlock()

	// load the data
	for i := 0; i < tl; i++ {
		// get the index of the next node to get
		ix := r.head - i
		if ix < 0 {
			ix = r.length + ix
		}

		// no data in the next node? return all we have so far
		if nil == r.nodes[ix] {
			return l[:i]
		}

		// copy the value
		l[i] = r.nodes[ix]
	}

	// return it all
	return l[:]
}
