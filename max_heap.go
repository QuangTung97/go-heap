package heap

type MaxHeap struct {
	elems []HeapElem
}

func NewMaxHeap(elems []HeapElem) *MaxHeap {
	heapElems := make([]HeapElem, len(elems))
	copy(heapElems, elems)
	h := &MaxHeap{elems: heapElems}
	h.heapify()
	return h
}

func (h *MaxHeap) heapifyAt(i int) {
	for {
		left := 2*i + 1
		right := 2*i + 2
		max_elem := i
		num := len(h.elems)

		if left < num && h.elems[left].Key > h.elems[max_elem].Key {
			max_elem = left
		}

		if right < num && h.elems[right].Key > h.elems[max_elem].Key {
			max_elem = right
		}

		if max_elem == i {
			return
		}
		// swap i and max_elem
		tmp := h.elems[i]
		h.elems[i] = h.elems[max_elem]
		h.elems[max_elem] = tmp

		i = max_elem
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

func (h *MaxHeap) FindMax() HeapElem {
	return h.elems[0]
}

func (h *MaxHeap) ExtractMax() HeapElem {
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

func (h *MaxHeap) Insert(elem HeapElem) int {
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

func (h *MaxHeap) FindTop(k uint) []HeapElem {
	result := make([]HeapElem, 0, k)
	for i := uint(0); i < k; i++ {
		result = append(result, h.ExtractMax())
	}
	for _, e := range result {
		h.Insert(e)
	}
	return result
}

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

func (h *MaxHeap) Array() []HeapElem {
	result := make([]HeapElem, 0, len(h.elems))
	for _, e := range h.elems {
		result = append(result, e)
	}
	return result
}
