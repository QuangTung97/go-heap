package heap

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMaxHeap(t *testing.T) {
	elems := []MaxHeapElem{
		{Key: 5, Value: 1}, {Key: 6, Value: 2},
		{Key: 2, Value: 3}, {Key: 1, Value: 4},
		{Key: 8, Value: 5}, {Key: 9, Value: 6},
		{Key: 3, Value: 7}, {Key: 4, Value: 8},
		{Key: 7, Value: 9},
	}
	assert.False(t, IsHeap(elems))

	h := NewMaxHeap(elems)
	assert.True(t, IsHeap(h.elems))

	expected := []MaxHeapElem{
		{Key: 9, Value: 6}, {Key: 8, Value: 5},
		{Key: 5, Value: 1}, {Key: 7, Value: 9},
		{Key: 6, Value: 2}, {Key: 2, Value: 3},
		{Key: 3, Value: 7}, {Key: 4, Value: 8},
		{Key: 1, Value: 4},
	}
	assert.Equal(t, expected, h.elems)
}

func TestExtractMax(t *testing.T) {
	elems := []MaxHeapElem{
		{Key: 5, Value: 1}, {Key: 6, Value: 2},
		{Key: 2, Value: 3}, {Key: 1, Value: 4},
		{Key: 8, Value: 5}, {Key: 9, Value: 6},
		{Key: 3, Value: 7}, {Key: 4, Value: 8},
		{Key: 7, Value: 9},
	}
	h := NewMaxHeap(elems)

	var expected MaxHeapElem

	expected = MaxHeapElem{Key: 9, Value: 6}
	assert.Equal(t, expected, h.ExtractMax())

	expected = MaxHeapElem{Key: 8, Value: 5}
	assert.Equal(t, expected, h.ExtractMax())

	expected = MaxHeapElem{Key: 7, Value: 9}
	assert.Equal(t, expected, h.ExtractMax())

	expected = MaxHeapElem{Key: 6, Value: 2}
	assert.Equal(t, expected, h.ExtractMax())

	assert.True(t, IsHeap(h.elems))
}

func TestInsert(t *testing.T) {
	elems := []MaxHeapElem{
		{Key: 5, Value: 1}, {Key: 6, Value: 2},
		{Key: 2, Value: 3}, {Key: 1, Value: 4},
		{Key: 8, Value: 5}, {Key: 9, Value: 6},
		{Key: 3, Value: 7}, {Key: 4, Value: 8},
		{Key: 7, Value: 9},
	}
	h := NewMaxHeap(elems)

	var index int
	var expected []MaxHeapElem

	index = h.Insert(MaxHeapElem{Key: 12, Value: 10})

	expected = []MaxHeapElem{
		{Key: 12, Value: 10}, {Key: 9, Value: 6},
		{Key: 5, Value: 1}, {Key: 7, Value: 9},
		{Key: 8, Value: 5}, {Key: 2, Value: 3},
		{Key: 3, Value: 7}, {Key: 4, Value: 8},
		{Key: 1, Value: 4}, {Key: 6, Value: 2},
	}
	assert.Equal(t, expected, h.elems)
	assert.Equal(t, 0, index)

	index = h.Insert(MaxHeapElem{Key: 7, Value: 11})

	expected = []MaxHeapElem{
		{Key: 12, Value: 10}, {Key: 9, Value: 6},
		{Key: 5, Value: 1}, {Key: 7, Value: 9},
		{Key: 8, Value: 5}, {Key: 2, Value: 3},
		{Key: 3, Value: 7}, {Key: 4, Value: 8},
		{Key: 1, Value: 4}, {Key: 6, Value: 2},
		{Key: 7, Value: 11},
	}
	assert.Equal(t, expected, h.elems)
	assert.Equal(t, 10, index)

	assert.True(t, IsHeap(h.elems))
}

func TestFindTop(t *testing.T) {
	elems := []MaxHeapElem{
		{Key: 5, Value: 1}, {Key: 6, Value: 2}, {Key: 2, Value: 3}, {Key: 1, Value: 4},
		{Key: 8, Value: 5}, {Key: 9, Value: 6}, {Key: 3, Value: 7}, {Key: 4, Value: 8},
		{Key: 7, Value: 9},
	}
	h := NewMaxHeap(elems)

	top := h.FindTop(4)

	expected := []MaxHeapElem{
		{Key: 9, Value: 6}, {Key: 8, Value: 5},
		{Key: 7, Value: 9}, {Key: 6, Value: 2},
	}
	assert.Equal(t, expected, top)

	assert.Equal(t, 9, len(h.elems))

	assert.True(t, IsHeap(h.elems))
}

func TestDeleteAt(t *testing.T) {
	t.Run("bubble down", func(t *testing.T) {
		elems := []MaxHeapElem{
			{Key: 5, Value: 1}, {Key: 6, Value: 2}, {Key: 2, Value: 3}, {Key: 1, Value: 4},
			{Key: 8, Value: 5}, {Key: 9, Value: 6}, {Key: 3, Value: 7}, {Key: 4, Value: 8},
			{Key: 7, Value: 9},
		}
		h := NewMaxHeap(elems)

		var expected []MaxHeapElem

		expected = []MaxHeapElem{
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
		expected = []MaxHeapElem{
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

		assert.True(t, IsHeap(h.elems))
	})

	t.Run("bubble up", func(t *testing.T) {
		elems := []MaxHeapElem{
			{Key: 5, Value: 1}, {Key: 6, Value: 2}, {Key: 2, Value: 3}, {Key: 1, Value: 4},
			{Key: 8, Value: 5}, {Key: 9, Value: 6}, {Key: 3, Value: 7}, {Key: 4, Value: 8},
			{Key: 7, Value: 9},
		}
		h := NewMaxHeap(elems)

		var expected []MaxHeapElem

		expected = []MaxHeapElem{
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

		expected = []MaxHeapElem{
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

		assert.True(t, IsHeap(h.elems))
	})
}

func randomNumbers(num int) []MaxHeapElem {
	elems := make([]MaxHeapElem, 0, num)
	for i := 0; i < num; i++ {
		elems = append(elems, MaxHeapElem{Key: rand.Uint64(), Value: uint64(i)})
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

var v MaxHeapElem

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
