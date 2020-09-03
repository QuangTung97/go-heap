package heap

// MinHeap a min heap
type MinHeap struct {
	elems []Elem
}

// NewMinHeap create a min heap
func NewMinHeap(elems []Elem) *MinHeap {
	heapElems := make([]Elem, len(elems))
	copy(heapElems, elems)
	h := &MinHeap{elems: heapElems}
	h.heapify()
	return h
}

func (h *MinHeap) heapifyAt(i int) {
	for {
		left := 2*i + 1
		right := 2*i + 2
		minElem := i
		num := len(h.elems)

		if left < num && h.elems[left].Key < h.elems[minElem].Key {
			minElem = left
		}

		if right < num && h.elems[right].Key < h.elems[minElem].Key {
			minElem = right
		}

		if minElem == i {
			return
		}
		// swap i and minElem
		tmp := h.elems[i]
		h.elems[i] = h.elems[minElem]
		h.elems[minElem] = tmp

		i = minElem
	}
}

func (h *MinHeap) heapify() {
	num := len(h.elems)
	num = (num + 1) / 2
	if num == 0 {
		return
	}
	for i := num; i > 0; i-- {
		h.heapifyAt(i - 1)
	}
}

// FindMin get the min element
func (h *MinHeap) FindMin() Elem {
	return h.elems[0]
}

// ExtractMin get and pop the min element
func (h *MinHeap) ExtractMin() Elem {
	result := h.elems[0]
	last := len(h.elems) - 1

	// swap 0 and end
	tmp := h.elems[0]
	h.elems[0] = h.elems[last]
	h.elems[last] = tmp

	h.elems = h.elems[:last]
	h.heapifyAt(0)

	return result
}

// Insert inserts a element
func (h *MinHeap) Insert(elem Elem) int {
	i := len(h.elems)
	h.elems = append(h.elems, elem)

	for {
		if i == 0 {
			return 0
		}
		parent := (i+1)/2 - 1
		if h.elems[parent].Key <= h.elems[i].Key {
			return i
		}

		// swap i and parent
		tmp := h.elems[i]
		h.elems[i] = h.elems[parent]
		h.elems[parent] = tmp
		i = parent
	}
}

// FindBottom find the k bottom elements
func (h *MinHeap) FindBottom(k uint) []Elem {
	result := make([]Elem, 0, k)
	for i := uint(0); i < k; i++ {
		result = append(result, h.ExtractMin())
	}
	for _, e := range result {
		h.Insert(e)
	}
	return result
}

// DeleteAt delete a element
func (h *MinHeap) DeleteAt(index int) {
	last := len(h.elems) - 1

	// swap index and last
	tmp := h.elems[last]
	h.elems[last] = h.elems[index]
	h.elems[index] = tmp

	h.elems = h.elems[:last]

	parent := (index+1)/2 - 1
	if index != 0 && h.elems[parent].Key > h.elems[index].Key {
		for {
			if index == 0 {
				return
			}
			parent := (index+1)/2 - 1
			if h.elems[parent].Key <= h.elems[index].Key {
				return
			}

			// swap i and parent
			tmp := h.elems[index]
			h.elems[index] = h.elems[parent]
			h.elems[parent] = tmp
			index = parent
		}
	} else {
		h.heapifyAt(index)
	}
}

// Array returns the internal array of min heap
func (h *MinHeap) Array() []Elem {
	result := make([]Elem, 0, len(h.elems))
	for _, e := range h.elems {
		result = append(result, e)
	}
	return result
}
