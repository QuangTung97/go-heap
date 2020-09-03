package heap

// MaxHeap a max heap
type MaxHeap struct {
	elems []Elem
}

// NewMaxHeap create a max heap
func NewMaxHeap(elems []Elem) *MaxHeap {
	heapElems := make([]Elem, len(elems))
	copy(heapElems, elems)
	h := &MaxHeap{elems: heapElems}
	h.heapify()
	return h
}

func (h *MaxHeap) heapifyAt(i int) {
	for {
		left := 2*i + 1
		right := 2*i + 2
		maxElem := i
		num := len(h.elems)

		if left < num && h.elems[left].Key > h.elems[maxElem].Key {
			maxElem = left
		}

		if right < num && h.elems[right].Key > h.elems[maxElem].Key {
			maxElem = right
		}

		if maxElem == i {
			return
		}
		// swap i and maxElem
		tmp := h.elems[i]
		h.elems[i] = h.elems[maxElem]
		h.elems[maxElem] = tmp

		i = maxElem
	}
}

func (h *MaxHeap) heapify() {
	num := len(h.elems)
	num = (num + 1) / 2
	if num == 0 {
		return
	}
	for i := num; i > 0; i-- {
		h.heapifyAt(i - 1)
	}
}

// FindMax find the max element
func (h *MaxHeap) FindMax() Elem {
	return h.elems[0]
}

// ExtractMax find and pop the max element
func (h *MaxHeap) ExtractMax() Elem {
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

// Insert insert a element
func (h *MaxHeap) Insert(elem Elem) int {
	i := len(h.elems)
	h.elems = append(h.elems, elem)

	for {
		if i == 0 {
			return 0
		}
		parent := (i+1)/2 - 1
		if h.elems[parent].Key >= h.elems[i].Key {
			return i
		}

		// swap i and parent
		tmp := h.elems[i]
		h.elems[i] = h.elems[parent]
		h.elems[parent] = tmp
		i = parent
	}
}

// FindTop get the top k elements
func (h *MaxHeap) FindTop(k uint) []Elem {
	result := make([]Elem, 0, k)
	for i := uint(0); i < k; i++ {
		result = append(result, h.ExtractMax())
	}
	for _, e := range result {
		h.Insert(e)
	}
	return result
}

// DeleteAt delete a position
func (h *MaxHeap) DeleteAt(index int) {
	last := len(h.elems) - 1

	// swap index and last
	tmp := h.elems[last]
	h.elems[last] = h.elems[index]
	h.elems[index] = tmp

	h.elems = h.elems[:last]

	parent := (index+1)/2 - 1
	if index != 0 && h.elems[parent].Key < h.elems[index].Key {
		for {
			if index == 0 {
				return
			}
			parent := (index+1)/2 - 1
			if h.elems[parent].Key >= h.elems[index].Key {
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

// Array returns the internal array (COPY) of heap
func (h *MaxHeap) Array() []Elem {
	result := make([]Elem, 0, len(h.elems))
	for _, e := range h.elems {
		result = append(result, e)
	}
	return result
}
