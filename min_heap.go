package heap

import "golang.org/x/exp/constraints"

// Min defines the Less function and it is used for Min Heap construction
func Min[T constraints.Integer](t1, t2 T) bool {
	return t1 < t2
}
