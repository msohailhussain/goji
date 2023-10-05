package heap

import "github.com/lorenzotinfena/goji/collections"

type BinaryHeapWithRemove[T comparable] struct {
	s     []T
	m     map[T]*collections.Set[int]
	prior func(T, T) bool
}

// Note for function Prior: It's a strict order relation
func NewBinaryHeapWithRemoveFromSlice[T comparable](
	s []T,
	prior func(T, T) bool,
) (h *BinaryHeapWithRemove[T]) {
	h = &BinaryHeapWithRemove[T]{
		s:     s,
		m:     make(map[T]*collections.Set[int]),
		prior: prior,
	}
	if h.Len() > 1 {
		for i := (h.Len() - 2) / 2; i >= 0; i-- {
			h.heapifyDown(i)
		}
	}
	return
}

func NewBinaryHeapWithRemove[T comparable](Prior func(T, T) bool) *BinaryHeapWithRemove[T] {
	return &BinaryHeapWithRemove[T]{s: make([]T, 0), m: make(map[T]*collections.Set[int]), prior: Prior}
}

func (h *BinaryHeapWithRemove[T]) Len() int {
	return len(h.s)
}

func (h *BinaryHeapWithRemove[T]) Push(value T) {
	set, exist := h.m[value]
	if !exist {
		h.m[value] = collections.NewSet[int]()
		set = h.m[value]
	}
	set.Add(len(h.s))
	h.s = append(h.s, value)
	h.heapifyUp(h.Len() - 1)
}

func (h *BinaryHeapWithRemove[T]) Pop() (res T) {
	res = h.s[0]
	set := h.m[h.s[0]]
	if set.Len() == 1 {
		delete(h.m, h.s[0])
	} else {
		set.Remove(0)
	}
	h.s[0] = h.s[h.Len()-1]
	set = h.m[h.s[0]]
	set.Remove(h.Len() - 1)
	set.Add(0)
	h.s = h.s[:h.Len()-1]
	h.heapifyDown(0)
	return
}

func (h *BinaryHeapWithRemove[T]) heapifyDown(index int) bool {
	origin := index
	for {
		j := index*2 + 2
		if j < h.Len() {
			if h.prior(h.s[j-1], h.s[j]) {
				j--
			}
		} else {
			j--
			if j >= h.Len() {
				break
			}
		}
		if h.prior(h.s[j], h.s[index]) {

			h.s[j], h.s[index] = h.s[index], h.s[j]
			if h.s[index] != h.s[j] {
				set := h.m[h.s[j]]
				set.Remove(index)
				set.Add(j)
				set = h.m[h.s[index]]
				set.Remove(j)
				set.Add(index)
			}
			index = j
		} else {
			break
		}
	}
	return origin != index
}

func (h *BinaryHeapWithRemove[T]) heapifyUp(index int) {
	for {
		if index == 0 {
			break
		}
		parent := (index - 1) / 2
		if h.prior(h.s[parent], h.s[index]) {
			break
		}
		h.s[index], h.s[parent] = h.s[parent], h.s[index]
		if h.s[index] != h.s[parent] {
			set := h.m[h.s[index]]
			set.Remove(parent)
			set.Add(index)
			set = h.m[h.s[parent]]
			set.Remove(index)
			set.Add(parent)
		}
		index = parent
	}
}

func (h *BinaryHeapWithRemove[T]) Preview() T {
	return h.s[0]
}

func (h BinaryHeapWithRemove[T]) String() string {
	return "" // #TODO
}

func (h *BinaryHeapWithRemove[T]) Remove(element T) {
	set := h.m[element]
	i := set.ToSlice()[0]
	if set.Len() == 1 {
		delete(h.m, element)
	} else {
		set.Remove(i) // to fix
	}
	if i == h.Len()-1 {
		h.s = h.s[:h.Len()-1]
	} else {
		set = h.m[h.s[h.Len()-1]]
		set.Remove(h.Len() - 1)
		set.Add(i)
		h.s[i] = h.s[h.Len()-1]
		h.s = h.s[:h.Len()-1]
		if !h.heapifyDown(i) {
			h.heapifyUp(i)
		}
	}
}