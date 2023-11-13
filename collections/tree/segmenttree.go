package tree

import "fmt"

type SegmentTree[E any, Q comparable] struct {
	query func(element E) Q
	merge func(q1 Q, q2 Q) Q

	Elements []E
	Segments []Q
}

// Pass nil as update if you never call Update method
// Assumptions:
//   - len(elements) > 0
func NewSegmentTree[E any, Q comparable](
	elements []E,
	query func(element E) Q,
	merge func(q1 Q, q2 Q) Q,
) *SegmentTree[E, Q] {
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

	return &SegmentTree[E, Q]{
		query: query,
		merge: merge,

		Elements: elements,
		Segments: segments,
	}
}

// Performs a query to [start, end]
// Assumptions:
//   - start <= end
//   - start and end are valid
func (s *SegmentTree[E, Q]) Query(start int, end int) Q {
	var queryRecRight func(i, l, r int) Q
	queryRecRight = func(i, l, r int) Q {
		if r == end {
			return s.Segments[i]
		}
		m := (l + r) / 2
		if end <= m {
			return queryRecRight(2*i+1, l, m)
		} else {
			return s.merge(
				s.Segments[2*i+1],
				queryRecRight(2*i+2, m+1, r),
			)
		}
	}
	var queryRecLeft func(i, l, r int) Q
	queryRecLeft = func(i, l, r int) Q {
		m := (l + r) / 2
		if start >= m+1 {
			return queryRecLeft(2*i+2, m+1, r)
		} else {
			if l == start {
				return s.Segments[i]
			}
			return s.merge(
				queryRecLeft(2*i+1, l, m),
				s.Segments[2*i+2],
			)
		}
	}
	var queryRec func(i, l, r int) Q
	queryRec = func(i, l, r int) Q {
		if l == r {
			return s.Segments[i]
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
	return queryRec(0, 0, len(s.Elements)-1)
}

// Assumptions:
//   - index is valid
func (s *SegmentTree[E, Q]) Update(index int, update func(oldQ Q, oldE E) (newQ Q)) {
	var updateRec func(i, l, r int)
	updateRec = func(i, l, r int) {
		s.Segments[i] = update(s.Segments[i], s.Elements[index])
		if l == r {
			return
		}
		m := (l + r) / 2
		if index <= m {
			updateRec(2*i+1, l, m)
		} else {
			updateRec(2*i+2, m+1, r)
		}
	}
	updateRec(0, 0, len(s.Elements)-1)
}

func (s *SegmentTree[E, Q]) Len() int {
	return len(s.Elements)
}

func (s *SegmentTree[E, Q]) String() string {
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
