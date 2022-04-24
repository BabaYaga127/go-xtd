package algo

import (
	"github.com/ndkimhao/go-xtd/iterator"
)

// Swap swaps the value of two iterator
func Swap[T any](a, b iterator.Iterator[T]) {
	va := a.Value()
	vb := b.Value()
	a.SetValue(vb)
	b.SetValue(va)
}

// Reverse reverse the elements in the range [first, last]
func Reverse[T any](first, last iterator.BidIterator[T]) {
	left := first.Clone().(iterator.BidIterator[T])
	right := last.Clone().(iterator.BidIterator[T]).Prev().(iterator.BidIterator[T])
	for !left.Equal(right) {
		Swap[T](left, right)
		left.Next()
		if left.Equal(right) {
			break
		}
		right.Prev()
	}
}
