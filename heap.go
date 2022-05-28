package heap

// Less returns true if t1 is lesser than t2
type Less[T any] func(t1, t2 T) bool

// Heapify changes a given slice of T to build a heap using the Less function as
// decision factor to swap values in the slice
func Heapify[T any](t []T, less Less[T]) {
	for nodeIdx := len(t)/2 - 1; nodeIdx >= 0; nodeIdx-- {
		down(nodeIdx, len(t)-1, less, t)
	}
}

// ExtractRoot remove the root value and rebalance the heap using the Less function
func ExtractRoot[T any](t []T, less Less[T]) (T, []T) {
	var root T

	length := len(t)
	if length == 0 {
		return root, t
	}

	lastIndex := length - 1
	t[0], t[lastIndex] = t[lastIndex], t[0]
	root = t[lastIndex]

	t = t[:lastIndex]

	down(0, len(t)-1, less, t)

	return root, t
}

func down[T any](currentIdx, endIdx int, less Less[T], heap []T) {
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
