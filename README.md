# Heap with Generics

A library to work with **Heaps** using **Generics**

The current heap solution in `container/heap` relies on the following interfaces:

```go
package sort

type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}
```

```go
package heap

type Interface interface {
	sort.Interface
	Push(x any)
	Pop() any
}
```

But with generics with **do not need** to **implement Interfaces**. Any slice will work as long as we define a `Less` function.

```go
type Less[T any] func(t1, t2 T) bool

func Heapify[T any](t []T, less Less[T])
```

This `Less` function allow us to create **any kind of Heap**, even with **custom algorithms** for complex data structures
