package heap

type HeapElem struct {
	Key   uint64
	Value uint64
}

func IsMaxHeap(elems []HeapElem) bool {
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

func IsMinHeap(elems []HeapElem) bool {
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
