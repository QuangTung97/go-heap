package heap

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMaxHeap(t *testing.T) {
	data := []uint64{5, 6, 2, 1, 8, 9, 3, 4, 7}
	h := NewMaxHeap(data)

	expected := []uint64{9, 8, 5, 7, 6, 2, 3, 4, 1}
	assert.Equal(t, expected, h.data)
}

func TestExtractMax(t *testing.T) {
	data := []uint64{5, 6, 2, 1, 8, 9, 3, 4, 7}
	h := NewMaxHeap(data)

	var expected uint64

	expected = 9
	assert.Equal(t, expected, h.ExtractMax())

	expected = 8
	assert.Equal(t, expected, h.ExtractMax())

	expected = 7
	assert.Equal(t, expected, h.ExtractMax())

	expected = 6
	assert.Equal(t, expected, h.ExtractMax())
}

func TestInsert(t *testing.T) {
	data := []uint64{5, 6, 2, 1, 8, 9, 3, 4, 7}
	h := NewMaxHeap(data)

	h.Insert(12)

	expected := []uint64{12, 9, 5, 7, 8, 2, 3, 4, 1, 6}
	assert.Equal(t, expected, h.data)
}

func TestFindTop(t *testing.T) {
	data := []uint64{5, 6, 2, 1, 8, 9, 3, 4, 7}
	h := NewMaxHeap(data)

	top := h.FindTop(4)

	expected := []uint64{9, 8, 7, 6}
	assert.Equal(t, expected, top)

	assert.Equal(t, 9, len(h.data))
}

func randomNumbers(num int) []uint64 {
	data := make([]uint64, 0, num)
	for i := 0; i < num; i++ {
		data = append(data, rand.Uint64())
	}
	return data
}

func BenchmarkNewMaxHeap(b *testing.B) {
	const numElems = 1000000
	data := randomNumbers(numElems)

	b.Run("random", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			NewMaxHeap(data)
		}
	})
}

var v uint64

func BenchmarkExtractMax(b *testing.B) {
	const numElems = 1000000
	data := randomNumbers(numElems)
	h := NewMaxHeap(data)

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
	data := randomNumbers(numElems)
	h := NewMaxHeap(data)

	var numTop uint64

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