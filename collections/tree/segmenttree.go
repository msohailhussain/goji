package tree

type SegmentTree[E any, Q any] struct {
	query  func(element E) Q
	merge  func(q1 Q, q2 Q) Q
	update func(oldQ Q, oldE E, newE E) (newQ Q)

	elements []E
	segments []Q
}

func NewSegmentTree[E any, Q any](
	elements []E,
	query func(element E) Q,
	merge func(q1 Q, q2 Q) Q,
	update func(oldQ Q, oldE E, newE E) (newQ Q),
) *SegmentTree[E, Q] {
	i := 1
	for i < len(elements) {
		i *= 2
	}
	segments := make([]Q, 2*i-1)

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
		query:  query,
		merge:  merge,
		update: update,

		elements: elements,
		segments: segments,
	}
}

func (s *SegmentTree[E, Q]) String() string {
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
	root := toString(0, 0, len(s.elements)-1)
	return root.String()
}
