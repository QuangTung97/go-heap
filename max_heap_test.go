package heap

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxHeapifyEmpty(t *testing.T) {
	h := NewMaxHeap([]Elem{})
	assert.Empty(t, h.Array())

	elems := []Elem{
		{Key: 5, Value: 1}, {Key: 6, Value: 2},
		{Key: 2, Value: 3}, {Key: 1, Value: 4},
		{Key: 8, Value: 5}, {Key: 9, Value: 6},
		{Key: 3, Value: 7}, {Key: 4, Value: 8},
		{Key: 7, Value: 9},
	}
	h = NewMaxHeap(elems)
	expected := []Elem{
		{Key: 9, Value: 6}, {Key: 8, Value: 5},
		{Key: 5, Value: 1}, {Key: 7, Value: 9},
		{Key: 6, Value: 2}, {Key: 2, Value: 3},
		{Key: 3, Value: 7}, {Key: 4, Value: 8},
		{Key: 1, Value: 4},
	}
	assert.Equal(t, expected, h.Array())
}

func TestNewMaxHeap(t *testing.T) {
	elems := []Elem{
		{Key: 5, Value: 1}, {Key: 6, Value: 2},
		{Key: 2, Value: 3}, {Key: 1, Value: 4},
		{Key: 8, Value: 5}, {Key: 9, Value: 6},
		{Key: 3, Value: 7}, {Key: 4, Value: 8},
		{Key: 7, Value: 9},
	}
	assert.False(t, IsMaxHeap(elems))

	h := NewMaxHeap(elems)
	assert.True(t, IsMaxHeap(h.elems))

	expected := []Elem{
		{Key: 9, Value: 6}, {Key: 8, Value: 5},
		{Key: 5, Value: 1}, {Key: 7, Value: 9},
		{Key: 6, Value: 2}, {Key: 2, Value: 3},
		{Key: 3, Value: 7}, {Key: 4, Value: 8},
		{Key: 1, Value: 4},
	}
	assert.Equal(t, expected, h.elems)
}

func TestExtractMax(t *testing.T) {
	elems := []Elem{
		{Key: 5, Value: 1}, {Key: 6, Value: 2},
		{Key: 2, Value: 3}, {Key: 1, Value: 4},
		{Key: 8, Value: 5}, {Key: 9, Value: 6},
		{Key: 3, Value: 7}, {Key: 4, Value: 8},
		{Key: 7, Value: 9},
	}
	h := NewMaxHeap(elems)

	var expected Elem

	expected = Elem{Key: 9, Value: 6}
	assert.Equal(t, expected, h.ExtractMax())

	expected = Elem{Key: 8, Value: 5}
	assert.Equal(t, expected, h.ExtractMax())

	expected = Elem{Key: 7, Value: 9}
	assert.Equal(t, expected, h.ExtractMax())

	expected = Elem{Key: 6, Value: 2}
	assert.Equal(t, expected, h.ExtractMax())

	assert.True(t, IsMaxHeap(h.elems))
}

func TestInsert(t *testing.T) {
	elems := []Elem{
		{Key: 5, Value: 1}, {Key: 6, Value: 2},
		{Key: 2, Value: 3}, {Key: 1, Value: 4},
		{Key: 8, Value: 5}, {Key: 9, Value: 6},
		{Key: 3, Value: 7}, {Key: 4, Value: 8},
		{Key: 7, Value: 9},
	}
	h := NewMaxHeap(elems)

	var index int
	var expected []Elem

	index = h.Insert(Elem{Key: 12, Value: 10})

	expected = []Elem{
		{Key: 12, Value: 10}, {Key: 9, Value: 6},
		{Key: 5, Value: 1}, {Key: 7, Value: 9},
		{Key: 8, Value: 5}, {Key: 2, Value: 3},
		{Key: 3, Value: 7}, {Key: 4, Value: 8},
		{Key: 1, Value: 4}, {Key: 6, Value: 2},
	}
	assert.Equal(t, expected, h.elems)
	assert.Equal(t, 0, index)

	index = h.Insert(Elem{Key: 7, Value: 11})

	expected = []Elem{
		{Key: 12, Value: 10}, {Key: 9, Value: 6},
		{Key: 5, Value: 1}, {Key: 7, Value: 9},
		{Key: 8, Value: 5}, {Key: 2, Value: 3},
		{Key: 3, Value: 7}, {Key: 4, Value: 8},
		{Key: 1, Value: 4}, {Key: 6, Value: 2},
		{Key: 7, Value: 11},
	}
	assert.Equal(t, expected, h.elems)
	assert.Equal(t, 10, index)

	assert.True(t, IsMaxHeap(h.elems))
}

func TestFindTop(t *testing.T) {
	elems := []Elem{
		{Key: 5, Value: 1}, {Key: 6, Value: 2}, {Key: 2, Value: 3}, {Key: 1, Value: 4},
		{Key: 8, Value: 5}, {Key: 9, Value: 6}, {Key: 3, Value: 7}, {Key: 4, Value: 8},
		{Key: 7, Value: 9},
	}
	h := NewMaxHeap(elems)

	top := h.FindTop(4)

	expected := []Elem{
		{Key: 9, Value: 6}, {Key: 8, Value: 5},
		{Key: 7, Value: 9}, {Key: 6, Value: 2},
	}
	assert.Equal(t, expected, top)

	assert.Equal(t, 9, len(h.elems))

	assert.True(t, IsMaxHeap(h.elems))
}

func TestDeleteAt(t *testing.T) {
	t.Run("bubble down", func(t *testing.T) {
		elems := []Elem{
			{Key: 5, Value: 1}, {Key: 6, Value: 2}, {Key: 2, Value: 3}, {Key: 1, Value: 4},
			{Key: 8, Value: 5}, {Key: 9, Value: 6}, {Key: 3, Value: 7}, {Key: 4, Value: 8},
			{Key: 7, Value: 9},
		}
		h := NewMaxHeap(elems)

		var expected []Elem

		expected = []Elem{
			{Key: 9, Value: 6},
			{Key: 8, Value: 5},
			{Key: 5, Value: 1},
			{Key: 7, Value: 9},
			{Key: 6, Value: 2},
			{Key: 2, Value: 3},
			{Key: 3, Value: 7},
			{Key: 4, Value: 8},
			{Key: 1, Value: 4},
		}
		assert.Equal(t, expected, h.elems)

		h.DeleteAt(3)
		expected = []Elem{
			{Key: 9, Value: 6},
			{Key: 8, Value: 5},
			{Key: 5, Value: 1},
			{Key: 4, Value: 8},
			{Key: 6, Value: 2},
			{Key: 2, Value: 3},
			{Key: 3, Value: 7},
			{Key: 1, Value: 4},
		}
		assert.Equal(t, expected, h.elems)
		assert.True(t, IsMaxHeap(h.elems))
	})

	t.Run("bubble up", func(t *testing.T) {
		elems := []Elem{
			{Key: 5, Value: 1}, {Key: 6, Value: 2}, {Key: 2, Value: 3}, {Key: 1, Value: 4},
			{Key: 8, Value: 5}, {Key: 9, Value: 6}, {Key: 3, Value: 7}, {Key: 4, Value: 8},
			{Key: 7, Value: 9},
		}
		h := NewMaxHeap(elems)

		var expected []Elem

		expected = []Elem{
			{Key: 9, Value: 6},
			{Key: 8, Value: 5},
			{Key: 5, Value: 1},
			{Key: 7, Value: 9},
			{Key: 6, Value: 2},
			{Key: 2, Value: 3},
			{Key: 3, Value: 7},
			{Key: 4, Value: 8},
			{Key: 1, Value: 4},
		}
		assert.Equal(t, expected, h.elems)

		h.DeleteAt(5)

		expected = []Elem{
			{Key: 9, Value: 6},
			{Key: 8, Value: 5},
			{Key: 5, Value: 1},
			{Key: 7, Value: 9},
			{Key: 6, Value: 2},
			{Key: 1, Value: 4},
			{Key: 3, Value: 7},
			{Key: 4, Value: 8},
		}
		assert.Equal(t, expected, h.elems)
		assert.True(t, IsMaxHeap(h.elems))
	})

	t.Run("at 0", func(t *testing.T) {
		elems := []Elem{
			{Key: 5, Value: 1}, {Key: 6, Value: 2}, {Key: 2, Value: 3}, {Key: 1, Value: 4},
			{Key: 8, Value: 5}, {Key: 9, Value: 6}, {Key: 3, Value: 7}, {Key: 4, Value: 8},
			{Key: 7, Value: 9},
		}
		h := NewMaxHeap(elems)

		var expected []Elem

		expected = []Elem{
			{Key: 9, Value: 6},
			{Key: 8, Value: 5}, {Key: 5, Value: 1},
			{Key: 7, Value: 9}, {Key: 6, Value: 2}, {Key: 2, Value: 3}, {Key: 3, Value: 7},
			{Key: 4, Value: 8}, {Key: 1, Value: 4},
		}
		assert.Equal(t, expected, h.elems)

		h.DeleteAt(0)

		expected = []Elem{
			{Key: 8, Value: 5},
			{Key: 7, Value: 9}, {Key: 5, Value: 1},
			{Key: 4, Value: 8}, {Key: 6, Value: 2}, {Key: 2, Value: 3}, {Key: 3, Value: 7},
			{Key: 1, Value: 4},
		}
		assert.Equal(t, expected, h.elems)
		assert.True(t, IsMaxHeap(h.elems))
	})
}

func randomNumbers(num int) []Elem {
	elems := make([]Elem, 0, num)
	for i := 0; i < num; i++ {
		elems = append(elems, Elem{Key: rand.Uint64(), Value: uint64(i)})
	}
	return elems
}

func BenchmarkNewMaxHeap(b *testing.B) {
	const numElems = 1000000
	elems := randomNumbers(numElems)

	b.Run("random", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			NewMaxHeap(elems)
		}
	})
}

var v Elem

func BenchmarkExtractMax(b *testing.B) {
	const numElems = 1000000
	elems := randomNumbers(numElems)
	h := NewMaxHeap(elems)

	b.Run("simple-extract-insert-1000000", func(b *testing.B) {
		b.StartTimer()
		for n := 0; n < b.N; n++ {
			v = h.ExtractMax()
			h.Insert(v)
		}
		b.StopTimer()
	})
}

func BenchmarkFindTop(b *testing.B) {
	const numElems = 1000000
	elems := randomNumbers(numElems)
	h := NewMaxHeap(elems)

	var numTop uint

	numTop = 10
	b.Run("top-10", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			h.FindTop(numTop)
		}
	})

	numTop = 100
	b.Run("top-100", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			h.FindTop(numTop)
		}
	})

	numTop = 1000
	b.Run("top-1000", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			h.FindTop(numTop)
		}
	})
}
