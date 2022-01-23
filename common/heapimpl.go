package common

import "fmt"

type heapImpl struct {
	array []Node
}

// NewHeap :
func NewHeap() (h Heap) {
	return &heapImpl{}
}

func (h *heapImpl) heapConsistent() {
	if len(h.array) > 1 {
		for i := 0; i < len(h.array); i++ {
			left, right := childrenOf(i)
			if (h.inHeap(left) && h.array[left].Less(h.array[i])) || (h.inHeap(right) && h.array[right].Less(h.array[i])) {
				fmt.Println(i, left, right)
				panic("heap inconsistent")
			}
		}
	}
}

func childrenOf(i int) (int, int) {
	idx := i + 1
	idx1 := 2 * idx
	idx2 := 2*idx + 1
	return idx1 - 1, idx2 - 1
}

func parentOf(i int) int {
	idx := i + 1
	idx0 := idx / 2
	return idx0 - 1
}

func (h *heapImpl) isLeaf(i int) bool {
	child1, _ := childrenOf(i)
	return child1 >= len(h.array)
}

func (h *heapImpl) inHeap(i int) bool {
	return i >= 0 && i < len(h.array)
}

func (h *heapImpl) fixHeapTopDown(parent int) {
	if h.isLeaf(parent) == false {
		left, right := childrenOf(parent)
		var less int = left
		if h.inHeap(right) && h.array[right].Less(h.array[left]) {
			less = right
		}
		if h.array[less].Less(h.array[parent]) {
			h.array[parent], h.array[less] = h.array[less], h.array[parent]
			h.fixHeapTopDown(less)
		}
	}
}

func (h *heapImpl) fixHeapBottomUp(child int) {
	if child > 0 {
		parent := parentOf(child)
		left, right := childrenOf(parent)
		var less int = left
		if h.inHeap(right) && h.array[right].Less(h.array[left]) {
			less = right
		}
		if h.array[less].Less(h.array[parent]) {
			h.array[parent], h.array[less] = h.array[less], h.array[parent]
			h.fixHeapBottomUp(parent)
		}
	}
}

func (h *heapImpl) Len() int {
	return len(h.array)
}

func (h *heapImpl) Push(node Node) {
	h.array = append(h.array, node)
	h.fixHeapBottomUp(len(h.array) - 1)
	h.heapConsistent()
}

func (h *heapImpl) Pop() (node Node) {
	if len(h.array) > 0 {
		node = h.array[0]
		h.array[0] = h.array[len(h.array)-1]
		h.array = h.array[:len(h.array)-1]
		h.fixHeapTopDown(0)
	} else {
		node = nil
	}
	h.heapConsistent()
	return node
}

func (h *heapImpl) heapify(i int) {
	if !h.isLeaf(i) {
		left, right := childrenOf(i)
		h.heapify(left)
		h.heapify(right)
		h.fixHeapTopDown(i)
	}
}

func (h *heapImpl) Heapify() {
	h.heapify(0)
	h.heapConsistent()
}
