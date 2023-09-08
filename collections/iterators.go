package collections

type dfsIterator[T comparable] struct {
	getNexts func(T) []T
	toVisit  Stack[T]
	visited  Set[T]
}

func (it *dfsIterator[T]) HasNext() bool {
	return it.toVisit.Len() != 0
}

func (it *dfsIterator[T]) Next() T {
	cur := it.toVisit.Pop()
	nexts := it.getNexts(cur)
	for _, v := range nexts {
		if !it.visited.Contains(v) {
			it.toVisit.Push(v)
		}
	}
	return cur
}

func NewIteratorDFS[T comparable](root T, getNexts func(T) []T) *dfsIterator[T] {
	toVisit := *NewStack[T](nil)
	toVisit.Push(root)

	return &dfsIterator[T]{
		getNexts: getNexts,
		toVisit:  toVisit,
		visited:  *NewSet[T](),
	}
}

type bfsIterator[T comparable] struct {
	getNexts func(T) []T
	toVisit  Queue[T]
	visited  Set[T]
}

func (it *bfsIterator[T]) HasNext() bool {
	return it.toVisit.Len() != 0
}

func (it *bfsIterator[T]) Next() T {
	cur := it.toVisit.Dequeue()
	nexts := it.getNexts(cur)
	for _, v := range nexts {
		if !it.visited.Contains(v) {
			it.toVisit.Enqueue(v)
		}
	}
	return cur
}

func NewIteratorBFS[T comparable](root T, getNexts func(T) []T) *bfsIterator[T] {
	toVisit := *NewQueue[T](nil)
	toVisit.Enqueue(root)

	return &bfsIterator[T]{
		getNexts: getNexts,
		toVisit:  toVisit,
		visited:  *NewSet[T](),
	}
}
