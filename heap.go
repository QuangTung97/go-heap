package heap

type MaxHeap struct {
	data []uint64
}

func NewMaxHeap(elems []uint64) *MaxHeap {
	data := make([]uint64, len(elems))
	copy(data, elems)
	h := &MaxHeap{data: data}
	h.heapify()
	return h
}

func (h *MaxHeap) heapifyAt(i int) {
	for {
		left := 2*i + 1
		right := 2*i + 2
		max_elem := i
		num := len(h.data)

		if left < num && h.data[left] > h.data[max_elem] {
			max_elem = left
		}

		if right < num && h.data[right] > h.data[max_elem] {
			max_elem = right
		}

		if max_elem == i {
			return
		}
		// swap i and max_elem
		tmp := h.data[i]
		h.data[i] = h.data[max_elem]
		h.data[max_elem] = tmp

		i = max_elem
	}
}

func (h *MaxHeap) heapify() {
	num := len(h.data)
	num = (num + 1) / 2
	if num == 0 {
		return
	}
	for i := num; i > 0; i-- {
		h.heapifyAt(i - 1)
	}
}

func (h *MaxHeap) FindMax() uint64 {
	return h.data[0]
}

func (h *MaxHeap) ExtractMax() uint64 {
	result := h.data[0]
	last := len(h.data) - 1

	// swap 0 and end
	tmp := h.data[0]
	h.data[0] = h.data[last]
	h.data[last] = tmp

	h.data = h.data[:last]
	h.heapifyAt(0)

	return result
}

func (h *MaxHeap) Insert(elem uint64) int {
	i := len(h.data)
	h.data = append(h.data, elem)

	for {
		if i == 0 {
			return 0
		}
		parent := (i+1)/2 - 1
		if h.data[parent] >= h.data[i] {
			return i
		}

		// swap i and parent
		tmp := h.data[i]
		h.data[i] = h.data[parent]
		h.data[parent] = tmp
		i = parent
	}
}

func (h *MaxHeap) FindTop(k uint64) []uint64 {
	result := make([]uint64, 0, k)
	for i := uint64(0); i < k; i++ {
		result = append(result, h.ExtractMax())
	}
	for _, e := range result {
		h.Insert(e)
	}
	return result
}

func (h *MaxHeap) DeleteAt(index int) {
	last := len(h.data) - 1

	// swap index and last
	tmp := h.data[last]
	h.data[last] = h.data[index]
	h.data[index] = tmp

	h.data = h.data[:last]

	parent := (index+1)/2 - 1
	if h.data[parent] < h.data[index] {
		for {
			if index == 0 {
				return
			}
			parent := (index+1)/2 - 1
			if h.data[parent] >= h.data[index] {
				return
			}

			// swap i and parent
			tmp := h.data[index]
			h.data[index] = h.data[parent]
			h.data[parent] = tmp
			index = parent
		}
	} else {
		h.heapifyAt(index)
	}
}

func IsHeap(data []uint64) bool {
	for i := range data {
		left := 2*i + 1
		right := 2*i + 2
		if left < len(data) && data[i] < data[left] {
			return false
		}
		if right < len(data) && data[i] < data[right] {
			return false
		}
	}
	return true
}
