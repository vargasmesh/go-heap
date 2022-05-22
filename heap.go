package heap

// Less returns true if t1 is lesser than t2
type Less[T any] func(t1, t2 T) bool

// Heapify changes a given slice of T to build a heap
// using the Less function as decision factor to swap values in the slice
func Heapify[T any](t []T, less Less[T]) {
	for nodeIdx := len(t)/2 - 1; nodeIdx >= 0; nodeIdx-- {
		down(nodeIdx, len(t)-1, less, t)
	}
}

func down[T any, L Less[T]](currentIdx, endIdx int, less L, heap []T) {
	childOneIdx := currentIdx*2 + 1
	for childOneIdx <= endIdx {
		idxToSwap := childOneIdx
		childTwoIdx := childOneIdx + 1
		if childTwoIdx <= endIdx && less(heap[childTwoIdx], heap[childOneIdx]) {
			idxToSwap = childTwoIdx
		}

		if less(heap[idxToSwap], heap[currentIdx]) {
			heap[currentIdx], heap[idxToSwap] = heap[idxToSwap], heap[currentIdx]
			currentIdx = idxToSwap
			childOneIdx = currentIdx*2 + 1
		} else {
			return
		}
	}
}
