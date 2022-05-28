package heap_test

import (
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vargasmesh/go-heap"
)

func TestExtractRoot(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	n := rand.Intn(50) + 1

	min := heap.Min[int]

	for i := 0; i < 5; i++ {
		t.Run("", func(t *testing.T) {

			lowestNumber := math.MaxInt
			slice := make([]int, n, n)
			for i := 0; i < n; i++ {
				number := -1000000 + rand.Intn(2000000)
				if min(number, lowestNumber) {
					lowestNumber = number
				}
				slice[i] = number
			}

			heap.Heapify(slice, min)

			root, slice := heap.ExtractRoot(slice, min)

			if assert.Equal(t, lowestNumber, root) {
				verifyMinHeap(t, slice)
			}

		})
	}

}

func TestInsert(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	n := rand.Intn(50) + 1

	min := heap.Min[int]

	for i := 0; i < 5; i++ {
		t.Run("", func(t *testing.T) {

			lowestNumber := math.MaxInt
			slice := make([]int, n, n)
			for i := 0; i < n; i++ {
				number := -1000000 + rand.Intn(2000000)
				if min(number, lowestNumber) {
					lowestNumber = number
				}
				slice[i] = number
			}

			heap.Heapify(slice, min)
			slice = heap.Insert((-1000000 + rand.Intn(2000000)), slice, min)
			verifyMinHeap(t, slice)
		})
	}
}
