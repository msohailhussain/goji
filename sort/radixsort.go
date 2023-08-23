package sort

import (
	"unsafe"

	"github.com/lorenzotinfena/goji/collections"
	"github.com/lorenzotinfena/goji/utils/constraints"
)

func RadixSort[T any, S constraints.Unsigned](v []T, getStructure func(T) []S) {
	data := make([]*collections.Pair[T, []S], len(v))
	foo := make([]collections.Pair[T, []S], len(v))
	for i := range v {
		foo[i] = collections.MakePair[T, []S](v[i], getStructure(v[i]))
		data[i] = &foo[i]
	}

	res := make([]*collections.Pair[T, []S], len(v))

	// get size of structure
	J := (int(unsafe.Sizeof(v[0])) * 8) - 1

	for b := len(data[0].Second) - 1; b >= 0; b-- {
		for j := 0; j <= J; j++ {

			// empty lists
			zero := 0
			one := 0

			for _, item := range data {
				if item.Second[b]&(1<<j) == 0 {
					zero++
				} else {
					one++
				}
			}
			
			one += zero

			for i := len(data) - 1; i >= 0; i-- {
				if data[i].Second[b]&(1<<j) == 0 {
					zero--
					res[zero] = data[i]
				} else {
					one--
					res[one] = data[i]
				}
			}
			res, data = data, res
		}
	}
	for i := range data {
		v[i] = data[i].First
	}
}

// linked list based approach
func RadixSortNaive[T any, S constraints.Unsigned](v []T, getStructure func(T) []S) {
	if len(v) <= 1 {
		return
	}

	type singleLinkedListNode struct {
		value *collections.Pair[T, []S]
		next  *singleLinkedListNode
	}

	v1 := make([]collections.Pair[T, []S], len(v))
	for i := range v {
		v1[i] = collections.MakePair[T, []S](v[i], getStructure(v[i]))
	}

	dataFirst := &singleLinkedListNode{value: &v1[0]}
	dataLast := dataFirst
	for i := 1; i < len(v1); i++ {
		dataLast.next = &singleLinkedListNode{value: &v1[i]}
		dataLast = dataLast.next
	}

	// get size of structure
	J := (int(unsafe.Sizeof(v[0])) * 8) - 1

	var zeroFirst, zeroLast, oneFirst, oneLast *singleLinkedListNode

	for b := len(v1[0].Second) - 1; b >= 0; b-- {
		for j := 0; j <= J; j++ {

			// empty lists
			zeroFirst = &singleLinkedListNode{}
			zeroLast = zeroFirst
			oneFirst = &singleLinkedListNode{}
			oneLast = oneFirst

			tmp := dataFirst
			for tmp != nil {
				if tmp.value.Second[b]&(1<<j) == 0 {
					zeroLast.next = tmp
					zeroLast = zeroLast.next
				} else {
					oneLast.next = tmp
					oneLast = oneLast.next
				}
				tmp = tmp.next
			}

			// merge lists
			oneLast.next = nil
			zeroLast.next = oneFirst.next
			dataFirst = zeroFirst.next
		}
	}
	i := 0
	tmp := dataFirst
	for tmp != nil {
		v[i] = tmp.value.First
		tmp = tmp.next
		i++
	}
}
