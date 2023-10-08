package tree

import "fmt"

// The main difference is pendingUpdates which holds for each segment its pending updates,
// When mergeUpdates is set all pending updates for each node are continuously merged,
// Otherwise they are all queued
type LazySegmentTree[Q comparable] struct {
	merge        func(q1 Q, q2 Q) Q
	mergeUpdates func(func(Q) Q, func(Q) Q) func(Q) Q

	numberElements int
	pendingUpdates [][]func(Q) Q
	segments       []Q
}

// Pass nil as mergeUpdates if you never call UpdateRange or if you can't merge updates
// Assumptions:
// - len(elements) > 0
// - query != nil
// - mergeUpdates != nil
func NewLazySegmentTree[E any, Q comparable](
	elements []E,
	query func(element E) Q,
	merge func(q1 Q, q2 Q) Q,
	mergeUpdates func(func(Q) Q, func(Q) Q) func(Q) Q,
) *LazySegmentTree[Q] {
	segments := make([]Q, 4*len(elements))

	var build func(i, l, r int)
	build = func(i, l, r int) {
		if l == r {
			segments[i] = query(elements[l])
			return
		}
		m := (l + r) / 2
		build(2*i+1, l, m)
		build(2*i+2, m+1, r)
		segments[i] = merge(segments[2*i+1], segments[2*i+2])
	}
	build(0, 0, len(elements)-1)

	pendingUpdates := make([][]func(Q) Q, 4*len(elements))
	for i := 0; i < 4*len(elements); i++ {
		pendingUpdates[i] = make([]func(Q) Q, 0)
	}

	return &LazySegmentTree[Q]{
		merge:        merge,
		mergeUpdates: mergeUpdates,

		numberElements: len(elements),
		pendingUpdates: pendingUpdates,
		segments:       segments,
	}
}

func (s *LazySegmentTree[Q]) insertPendingUpdate(i int, update func(Q) Q) {
	if s.mergeUpdates == nil {
		s.pendingUpdates[i] = append(s.pendingUpdates[i], update)
	} else {
		if len(s.pendingUpdates[i]) == 0 {
			s.pendingUpdates[i] = append(s.pendingUpdates[i], update)
		} else {
			s.pendingUpdates[i][0] = s.mergeUpdates(s.pendingUpdates[i][0], update)
		}
	}
}

func (s *LazySegmentTree[Q]) push(i, l, r int) {
	for _, f := range s.pendingUpdates[i] {
		s.segments[i] = f(s.segments[i])
	}
	if l != r {
		for _, f := range s.pendingUpdates[i] {
			s.insertPendingUpdate(2*i+1, f)
			s.insertPendingUpdate(2*i+2, f)
		}
	}
	s.pendingUpdates[i] = make([]func(Q) Q, 0)
}

// Performs a query to [start, end]
// Assumptions:
//   - start <= end
//   - start and end are valid
func (s *LazySegmentTree[Q]) Query(start int, end int) Q {
	var queryRecRight func(i, l, r int) Q
	queryRecRight = func(i, l, r int) Q {
		s.push(i, l, r)
		if r == end {
			return s.segments[i]
		}
		m := (l + r) / 2
		if end <= m {
			return queryRecRight(2*i+1, l, m)
		} else {
			return s.merge(
				s.segments[2*i+1],
				queryRecRight(2*i+2, m+1, r),
			)
		}
	}
	var queryRecLeft func(i, l, r int) Q
	queryRecLeft = func(i, l, r int) Q {
		s.push(i, l, r)
		m := (l + r) / 2
		if start >= m+1 {
			return queryRecLeft(2*i+2, m+1, r)
		} else {
			if l == start {
				return s.segments[i]
			}
			return s.merge(
				queryRecLeft(2*i+1, l, m),
				s.segments[2*i+2],
			)
		}
	}
	var queryRec func(i, l, r int) Q
	queryRec = func(i, l, r int) Q {
		s.push(i, l, r)
		if l == r {
			return s.segments[i]
		}
		m := (l + r) / 2
		if end <= m {
			return queryRec(2*i+1, l, m)
		} else if start >= m+1 {
			return queryRec(2*i+2, m+1, r)
		} else {
			return s.merge(
				queryRecLeft(2*i+1, l, m),
				queryRecRight(2*i+2, m+1, r),
			)
		}
	}
	return queryRec(0, 0, s.numberElements-1)
}

// Assumptions:
//   - start and end are valid
//   - update != nil
func (s *LazySegmentTree[Q]) UpdateRange(start, end int, update func(Q) Q) {
	lazyUpdate := func(i int, l, r int, update func(Q) Q) {
		s.push(i, l, r) // TODO I don't know why but if I remove this it stops working
		s.insertPendingUpdate(i, update)
		s.push(i, l, r)
	}
	var updateRecRight func(i, l, r int)
	updateRecRight = func(i, l, r int) {
		if r == end {
			lazyUpdate(i, l, r, update)
			return
		}

		s.push(i, l, r)
		m := (l + r) / 2
		if end <= m {

			updateRecRight(2*i+1, l, m)
			s.push(2*i+2, m+1, r)
		} else {

			lazyUpdate(2*i+1, l, m, update)
			updateRecRight(2*i+2, m+1, r)
		}
		s.segments[i] = s.merge(s.segments[2*i+1], s.segments[2*i+2])
	}
	var updateRecLeft func(i, l, r int)
	updateRecLeft = func(i, l, r int) {
		m := (l + r) / 2
		if start >= m+1 {
			s.push(i, l, r)
			updateRecLeft(2*i+2, m+1, r)
			s.push(2*i+1, l, m)
		} else {
			if l == start {
				lazyUpdate(i, l, r, update)
				return
			}
			s.push(i, l, r)
			lazyUpdate(2*i+2, m+1, r, update)
			updateRecLeft(2*i+1, l, m)
		}
		s.segments[i] = s.merge(s.segments[2*i+1], s.segments[2*i+2])
	}
	var updateRec func(i, l, r int)
	updateRec = func(i, l, r int) {
		if start == l && end == r {
			lazyUpdate(i, l, r, update)
			return
		}
		s.push(i, l, r)
		m := (l + r) / 2
		if end <= m {
			updateRec(2*i+1, l, m)
			s.push(2*i+2, m+1, r)
		} else if start >= m+1 {
			updateRec(2*i+2, m+1, r)
			s.push(2*i+1, l, m)
		} else {
			updateRecLeft(2*i+1, l, m)
			updateRecRight(2*i+2, m+1, r)
		}
		s.segments[i] = s.merge(s.segments[2*i+1], s.segments[2*i+2])
	}
	updateRec(0, 0, s.numberElements-1)
}

func (s *LazySegmentTree[Q]) Len() int {
	return s.numberElements
}

func (s *LazySegmentTree[Q]) String() string {
	var rec func(l, r int) *TreeNode[string]
	rec = func(l, r int) *TreeNode[string] {
		node := &TreeNode[string]{
			Value: fmt.Sprint(l) + "━━━" + fmt.Sprint(r) + "\n" + fmt.Sprint(s.Query(l, r)),
		}
		if l != r {
			m := (l + r) / 2
			node.Children = []*TreeNode[string]{rec(l, m), rec(m+1, r)}
		}
		return node
	}
	return rec(0, s.Len()-1).String()
}

func (s *LazySegmentTree[Q]) String1() string {
	var rec func(i, l, r int) *TreeNode[string]
	rec = func(i, l, r int) *TreeNode[string] {
		node := &TreeNode[string]{
			Value: fmt.Sprint(l+1) + "━━━" + fmt.Sprint(r+1) + "\n" + fmt.Sprint(s.segments[i]) + "\n" + fmt.Sprint(len(s.pendingUpdates[i])),
		}
		if l != r {
			m := (l + r) / 2
			node.Children = []*TreeNode[string]{rec(2*i+1, l, m), rec(2*i+2, m+1, r)}
		}
		return node
	}
	return rec(0, 0, s.Len()-1).String()
}
