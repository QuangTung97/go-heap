package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMinHeap(t *testing.T) {
	elems := []HeapElem{
		{Key: 5, Value: 1}, {Key: 6, Value: 2},
		{Key: 2, Value: 3}, {Key: 1, Value: 4},
		{Key: 8, Value: 5}, {Key: 9, Value: 6},
		{Key: 3, Value: 7}, {Key: 4, Value: 8},
		{Key: 7, Value: 9},
	}
	assert.False(t, IsMinHeap(elems))

	h := NewMinHeap(elems)
	assert.True(t, IsMinHeap(h.elems))

	expected := []HeapElem{
		{Key: 1, Value: 4},
		{Key: 4, Value: 8}, {Key: 2, Value: 3},
		{Key: 5, Value: 1}, {Key: 8, Value: 5}, {Key: 9, Value: 6}, {Key: 3, Value: 7},
		{Key: 6, Value: 2}, {Key: 7, Value: 9},
	}
	assert.Equal(t, expected, h.elems)

	expected = []HeapElem{
		{Key: 1, Value: 4},
		{Key: 2, Value: 3},
		{Key: 3, Value: 7},
	}
	assert.Equal(t, expected, h.FindBottom(3))

}

func TestMinHeapDeleteAt(t *testing.T) {
	elems := []HeapElem{
		{Key: 5, Value: 1}, {Key: 6, Value: 2},
		{Key: 2, Value: 3}, {Key: 1, Value: 4},
		{Key: 8, Value: 5}, {Key: 9, Value: 6},
		{Key: 3, Value: 7}, {Key: 4, Value: 8},
		{Key: 7, Value: 9},
	}
	h := NewMinHeap(elems)

	h.DeleteAt(0)
	expected := []HeapElem{
		{Key: 2, Value: 3},
		{Key: 4, Value: 8}, {Key: 3, Value: 7},
		{Key: 5, Value: 1}, {Key: 8, Value: 5}, {Key: 9, Value: 6}, {Key: 7, Value: 9},
		{Key: 6, Value: 2},
	}
	assert.Equal(t, expected, h.elems)
	assert.True(t, IsMinHeap(h.elems))
}
