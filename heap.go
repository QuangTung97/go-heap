package heap

// Elem represents a heap element
type Elem struct {
	Key   uint64
	Value uint64
}

// IsMaxHeap check if array is a max heap
func IsMaxHeap(elems []Elem) bool {
	for i := range elems {
		left := 2*i + 1
		right := 2*i + 2
		if left < len(elems) && elems[i].Key < elems[left].Key {
			return false
		}
		if right < len(elems) && elems[i].Key < elems[right].Key {
			return false
		}
	}
	return true
}

// IsMinHeap check if array is a min heap
func IsMinHeap(elems []Elem) bool {
	for i := range elems {
		left := 2*i + 1
		right := 2*i + 2
		if left < len(elems) && elems[i].Key > elems[left].Key {
			return false
		}
		if right < len(elems) && elems[i].Key > elems[right].Key {
			return false
		}
	}
	return true
}
