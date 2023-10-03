package tree

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

func (s *LazySegmentTree[Q]) String() string {
	var toString func(i, l, r int) *TreeNode[Q]
	toString = func(i, l, r int) *TreeNode[Q] {
		children := make([]*TreeNode[Q], 0)
		if l != r {
			m := (l + r) / 2
			children = append(children, toString(2*i+1, l, m))
			children = append(children, toString(2*i+2, m+1, r))
		}
		return &TreeNode[Q]{
			Value:    s.segments[i],
			Children: children,
		}
	}
	root := toString(0, 0, s.numberElements-1)
	return root.String()
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

// Performs a query to [start, end]
// Assumptions:
//   - start <= end
//   - start and end are valid
func (s *LazySegmentTree[Q]) Query(start int, end int) Q {
	push := func(i, l, r int) {
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
	var queryRecRight func(i, l, r int) Q
	queryRecRight = func(i, l, r int) Q {
		push(i, l, r)
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
		push(i, l, r)
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
		push(i, l, r)
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
	var updateRecRight func(i, l, r int)
	updateRecRight = func(i, l, r int) {
		if r == end {
			s.insertPendingUpdate(i, update)
			return
		}
		m := (l + r) / 2
		if end <= m {
			updateRecRight(2*i+1, l, m)
		} else {
			s.insertPendingUpdate(2*i+1, update)
			updateRecRight(2*i+2, m+1, r)
		}
	}
	var updateRecLeft func(i, l, r int)
	updateRecLeft = func(i, l, r int) {
		m := (l + r) / 2
		if start >= m+1 {
			updateRecLeft(2*i+2, m+1, r)
		} else {
			if l == start {
				s.insertPendingUpdate(i, update)
				return
			}
			updateRecLeft(2*i+1, l, m)
			s.insertPendingUpdate(2*i+2, update)
		}
	}
	var updateRec func(i, l, r int)
	updateRec = func(i, l, r int) {
		if l == r {
			s.insertPendingUpdate(i, update)
			return
		}
		m := (l + r) / 2
		if end <= m {
			updateRec(2*i+1, l, m)
		} else if start >= m+1 {
			updateRec(2*i+2, m+1, r)
		} else {
			updateRecLeft(2*i+1, l, m)
			updateRecRight(2*i+2, m+1, r)
		}
	}
	updateRec(0, 0, s.numberElements-1)
}
