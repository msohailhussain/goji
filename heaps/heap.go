type BinaryHeap[T any] struct {
	s     []T
	prior func(T, T) bool
}

// Note for function Prior: It's a strict order relation
func NewBinaryHeapFromSlice[T any](
	s []T,
	Prior func(T, T) bool,
) (h *BinaryHeap[T]) {
	h = &BinaryHeap[T]{
		s:     s,
		prior: Prior,
	}
	if h.Len() > 1 {
		for i := (int64(h.Len()) - 2) / 2; i >= 0; i-- {
			h.heapifyDown(int64(i))
		}
	}
	return
}

func NewBinaryHeap[T any](Prior func(T, T) bool) *BinaryHeap[T] {
	return &BinaryHeap[T]{s: make([]T, 0), prior: Prior}
}

func (h *BinaryHeap[T]) Len() uint64 {
	return uint64(len(h.s))
}

func (h *BinaryHeap[T]) Push(value T) {
	h.s = append(h.s, value)
	h.heapifyUp(int64(h.Len() - 1))
}

func (h *BinaryHeap[T]) Pop() (res T) {
	res = h.s[0]
	h.s[0] = h.s[h.Len()-1]
	h.s = h.s[:h.Len()-1]
	h.heapifyDown(0)
	return
}

func (h *BinaryHeap[T]) heapifyDown(index int64) bool {
	origin := index
	for {
		j := index*2 + 2
		if j < int64(h.Len()) {
			if h.prior(h.s[j-1], h.s[j]) {
				j--
			}
		} else {
			j--
			if j >= int64(h.Len()) {
				break
			}
		}
		if h.prior(h.s[j], h.s[index]) {
			h.s[j], h.s[index] = h.s[index], h.s[j]
			index = j
		} else {
			break
		}
	}
	return origin != index
}

func (h *BinaryHeap[T]) heapifyUp(index int64) {
	for {
		if index == 0 {
			break
		}
		parent := (index - 1) / 2
		if h.prior(h.s[parent], h.s[index]) {
			break
		}
		h.s[index], h.s[parent] = h.s[parent], h.s[index]
		index = parent
	}
}

func (h *BinaryHeap[T]) Preview() T {
	return h.s[0]
}

func (h BinaryHeap[T]) Log() string {
	return "" // #TODO
}
