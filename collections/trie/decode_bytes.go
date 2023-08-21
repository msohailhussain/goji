package trie

import (
	"fmt"
)

func toByte[T any](item T) []byte {
	return []byte(fmt.Sprintf("%v", item))
}
