package trees

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/lorenzotinfena/goji/collections"
)

type TreeNode[T any] struct {
	Value    T
	Children []*TreeNode[T]
}

// Using code from: github.com/shivamMg/ppds/blob/master/tree/tree.go
func (root TreeNode[T]) String() string {
	const (
		BoxVer       = "│"
		BoxHor       = "─"
		BoxVerRight  = "├"
		BoxDownLeft  = "┐"
		BoxDownRight = "┌"
		BoxDownHor   = "┬"
		BoxUpRight   = "└"
		Gutter       = 2
	)
	parents := map[*TreeNode[T]]*TreeNode[T]{}
	var setParents func(parents map[*TreeNode[T]]*TreeNode[T], root *TreeNode[T])
	setParents = func(parents map[*TreeNode[T]]*TreeNode[T], root *TreeNode[T]) {
		for _, c := range root.Children {
			parents[c] = root
			setParents(parents, c)
		}
	}
	setParents(parents, &root)

	var lines func(root *TreeNode[T]) (s []string)

	lines = func(root *TreeNode[T]) (s []string) {
		data := fmt.Sprintf("%s %v ", BoxHor, root.Value)
		l := len(root.Children)
		if l == 0 {
			s = append(s, data)
			return
		}

		w := utf8.RuneCountInString(data) // remove import unidcode/utf8
		for i, c := range root.Children {
			for j, line := range lines(c) {
				if i == 0 && j == 0 {
					if l == 1 {
						s = append(s, data+BoxHor+line)
					} else {
						s = append(s, data+BoxDownHor+line)
					}
					continue
				}

				var box string
				if i == l-1 && j == 0 {
					box = BoxUpRight
				} else if i == l-1 {
					box = " "
				} else if j == 0 {
					box = BoxVerRight
				} else {
					box = BoxVer
				}
				s = append(s, strings.Repeat(" ", w)+box+line)
			}
		}
		return
	}

	var lines2 func(root *TreeNode[T]) (s []string)
	lines2 = func(root *TreeNode[T]) (s []string) {
		s = append(s, fmt.Sprintf("%v", root.Value))
		l := len(root.Children)
		if l == 0 {
			return
		}

		for i, c := range root.Children {
			for j, line := range lines2(c) {
				// first line of the last child
				if i == l-1 && j == 0 {
					s = append(s, BoxUpRight+BoxHor+" "+line)
				} else if j == 0 {
					s = append(s, BoxVerRight+BoxHor+" "+line)
				} else if i == l-1 {
					s = append(s, "   "+line)
				} else {
					s = append(s, BoxVer+"  "+line)
				}
			}
		}
		return
	}
	safeData := func(n *TreeNode[T]) string {
		data := fmt.Sprintf("%v", n.Value)
		if data == "" {
			return " "
		}
		return data
	}
	var width func(widths map[*TreeNode[T]]int, n *TreeNode[T]) int
	width = func(widths map[*TreeNode[T]]int, n *TreeNode[T]) int {
		if w, ok := widths[n]; ok {
			return w
		}

		w := utf8.RuneCountInString(safeData(n)) + Gutter
		widths[n] = w
		if len(n.Children) == 0 {
			return w
		}

		sum := 0
		for _, c := range n.Children {
			sum += width(widths, c)
		}
		if sum > w {
			widths[n] = sum
			return sum
		}
		return w
	}

	var setPaddings func(paddings map[*TreeNode[T]]int, widths map[*TreeNode[T]]int, pad int, root *TreeNode[T])
	setPaddings = func(paddings map[*TreeNode[T]]int, widths map[*TreeNode[T]]int, pad int, root *TreeNode[T]) {
		for _, c := range root.Children {
			paddings[c] = pad
			setPaddings(paddings, widths, pad, c)
			pad += width(widths, c)
		}
	}

	isLeftMostChild := func(n *TreeNode[T]) bool {
		p, ok := parents[n]
		if !ok {
			// root
			return true
		}
		return p.Children[0] == n
	}

	paddings := map[*TreeNode[T]]int{}

	setPaddings(paddings, map[*TreeNode[T]]int{}, 0, &root)

	q := collections.NewQueue[*TreeNode[T]]()
	q.Enqueue(&root)
	linesss := []string{}
	for q.Len() > 0 {
		branches, nodes := "", ""
		covered := 0
		qLen := q.Len()
		for i := 0; i < qLen; i++ {
			n := q.Dequeue()
			for _, c := range n.Children {
				q.Enqueue(c)
			}

			spaces := paddings[n] - covered
			data := safeData(n)
			nodes += strings.Repeat(" ", spaces) + data

			w := utf8.RuneCountInString(data)
			covered += spaces + w
			var preview *TreeNode[T] = nil
			if q.Len() > 0 {
				preview = q.Preview()
			}
			current, next := isLeftMostChild(n), isLeftMostChild(preview)
			if current {
				branches += strings.Repeat(" ", spaces)
			} else {
				branches += strings.Repeat(BoxHor, spaces)
			}

			if current && next {
				branches += BoxVer
			} else if current {
				branches += BoxVerRight
			} else if next {
				branches += BoxDownLeft
			} else {
				branches += BoxDownHor
			}

			if next {
				branches += strings.Repeat(" ", w-1)
			} else {
				branches += strings.Repeat(BoxHor, w-1)
			}
		}
		linesss = append(linesss, branches, nodes)
	}

	s := ""
	for _, line := range linesss[1:] {
		s += strings.TrimRight(line, " ") + "\n"
	}
	return s
}
