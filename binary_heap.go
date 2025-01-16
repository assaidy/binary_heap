package binary_heap

// CompareFunc defines a generic function type for comparing two elements.
type CompareFunc[T any] func(a, b T) bool

// Heap represents a generic binary heap structure.
type Heap[T any] struct {
	elems   []T
	compare CompareFunc[T]
}

// NewHeap creates a new heap with the given elements and comparison function.
// It creates a new internal array, and doesn't affect the one passed as a parameter.
func NewHeap[T any](elems []T, compare CompareFunc[T]) *Heap[T] {
	heap := Heap[T]{compare: compare}
	heap.elems = make([]T, len(elems))
	copy(heap.elems, elems)
	Heapify(heap.elems, compare)
	return &heap
}

// Push adds a new element to the heap and maintains heap property.
func (me *Heap[T]) Push(elem T) {
	me.elems = append(me.elems, elem)
	heapifyUp(me.elems, len(me.elems)-1, me.compare)
}

// Pop removes and returns the root element of the heap.
func (me *Heap[T]) Pop() T {
	if len(me.elems) == 0 {
		panic("heap is empty")
	}
	result := me.elems[0]
	n := len(me.elems)
	me.elems[0] = me.elems[n-1]
	me.elems = me.elems[:n-1]
	heapifyDown(me.elems, 0, me.compare)
	return result
}

// Length returns the number of elements in the heap.
func (me *Heap[T]) Length() int {
	return len(me.elems)
}

// IsEmpty checks if the heap is empty.
func (me *Heap[T]) IsEmpty() bool {
	return len(me.elems) == 0
}

// Heapify transforms a slice into a valid heap based on the comparison function.
func Heapify[T any](elems []T, compare CompareFunc[T]) {
	// Start from the last parent node and heapify down.
	for i := parentIdx(len(elems) - 1); i >= 0; i-- {
		heapifyDown(elems, i, compare)
	}
}

// HeapSort sorts a slice in-place using the heap sort algorithm.
func HeapSort[T any](elems []T, compare CompareFunc[T]) {
	if len(elems) == 0 {
		return
	}
	reverseCompare := func(a, b T) bool { return compare(b, a) } // Reverse the comparison.
	Heapify(elems, reverseCompare)
	heap := Heap[T]{elems: elems, compare: reverseCompare}
	for i := len(elems) - 1; i >= 0; i-- {
		elems[i] = heap.Pop()
	}
}

// heapifyDown restores the heap property by sinking an element down the tree.
func heapifyDown[T any](elems []T, i int, compare CompareFunc[T]) {
	var (
		n     = len(elems)
		curr  = i
		left  = leftIdx(i)
		right = rightIdx(i)
	)
	if left < n && compare(elems[left], elems[curr]) {
		curr = left
	}
	if right < n && compare(elems[right], elems[curr]) {
		curr = right
	}
	if curr != i {
		elems[i], elems[curr] = elems[curr], elems[i]
		heapifyDown(elems, curr, compare)
	}
}

// heapifyUp restores the heap property by bubbling an element up the tree.
func heapifyUp[T any](elems []T, i int, compare CompareFunc[T]) {
	if i > 0 {
		p := parentIdx(i)
		if compare(elems[i], elems[p]) {
			elems[i], elems[p] = elems[p], elems[i]
			heapifyUp(elems, p, compare)
		}
	}
}

// parentIdx calculates the index of the parent node.
func parentIdx(i int) int {
	return (i - 1) / 2
}

// leftIdx calculates the index of the left child node.
func leftIdx(i int) int {
	return 2*i + 1
}

// rightIdx calculates the index of the right child node.
func rightIdx(i int) int {
	return 2*i + 2
}
