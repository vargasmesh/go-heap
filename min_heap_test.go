package heap_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/vargasmesh/go-heap"
)

func TestMin(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	n := rand.Intn(10000000)

	for i := 0; i < 5; i++ {
		t.Run("", func(t *testing.T) {

			slice := make([]int, n, n)
			for i := 0; i < n; i++ {
				slice[i] = -1000000 + rand.Intn(2000000)
			}

			heap.Heapify(slice, heap.Min[int])

			verifyMinHeap(t, slice)
		})
	}

}

func verifyMinHeap(t *testing.T, h []int) {
	t.Helper()

	less := heap.Min[int]

	length := len(h)

	for index, currentNode := range h {
		childLeft := index*2 + 1
		childRight := index*2 + 2

		if childLeft < length {
			if less(h[childLeft], currentNode) {
				t.FailNow()
			}
		}

		if childRight < length {
			if less(h[childRight], currentNode) {
				t.FailNow()
			}
		}

	}
}
