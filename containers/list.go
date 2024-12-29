package containers

type List[T any] struct {
	Items []T
}

func NewList[T any]() *List[T] {
	return &List[T]{
		Items: make([]T, 0),
	}
}

func (l *List[T]) Add(item T) {
	l.Items = append(l.Items, item)
}
func (l *List[T]) Remove(i int) {
	l.Items = append(l.Items[:i], l.Items[i+1:]...)
}
func (l *List[T]) PermutionIterator(yield func([]T) bool) {
	items := make([]T, len(l.Items))
	copy(items, l.Items)

	var generate func(k int)
	keepGoing := true

	// Heap's algorithm
	generate = func(k int) {
		if k == 1 {
			// Create a copy of the slice to avoid modifying the original array
			keepGoing = yield(items)
			return
		}

		for i := 0; i < k; i++ {
			if !keepGoing {
				return
			}
			generate(k - 1)
			if k%2 == 0 {
				items[i], items[k-1] = items[k-1], items[i]
			} else {
				items[0], items[k-1] = items[k-1], items[0]
			}
		}
	}

	generate(len(items))
}
