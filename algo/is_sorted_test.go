package algo_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ndkimhao/go-xtd/algo"
	"github.com/ndkimhao/go-xtd/slice"
)

func TestIsSortedUntil(t *testing.T) {
	t.Run("Normal", func(t *testing.T) {
		s := slice.Of(1, 2, 3, 3, 2, 3, 4)
		it := algo.IsSortedUntil[int](s.Begin(), s.End())
		assert.Equal(t, reflect.TypeOf(it).Name(), "Iterator[int]")
		assert.True(t, it.Equal(s.Begin().Add(4)))
	})
	t.Run("One", func(t *testing.T) {
		s := slice.Of(1, 0, 1)
		it := algo.IsSortedUntil[int](s.Begin(), s.End())
		assert.True(t, it.Equal(s.Begin().Add(1)))
	})
	t.Run("Empty", func(t *testing.T) {
		s := slice.Of[int]()
		it := algo.IsSortedUntil[int](s.Begin(), s.End())
		assert.True(t, it.Equal(s.End()))
	})
}

func TestIsSorted(t *testing.T) {
	s := slice.Of[int]()
	assert.True(t, algo.IsSorted[int](s.Begin(), s.End()))
	s = slice.Of(1)
	assert.True(t, algo.IsSorted[int](s.Begin(), s.End()))
	s = slice.Of(1, 2)
	assert.True(t, algo.IsSorted[int](s.Begin(), s.End()))
	s = slice.Of(1, 2, 3, 3, 3, 5)
	assert.True(t, algo.IsSorted[int](s.Begin(), s.End()))

	s = slice.Of(1, 2, 3, 2, 3)
	assert.False(t, algo.IsSorted[int](s.Begin(), s.End()))
	s = slice.Of(2, 2, 2, 1)
	assert.False(t, algo.IsSorted[int](s.Begin(), s.End()))
}
