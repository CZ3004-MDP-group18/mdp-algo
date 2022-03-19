package common

import "container/heap"

type PriorityQueue struct {
	array    []Node
	indexMap map[Node]int
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{
		array:    nil,
		indexMap: make(map[Node]int),
	}
}

func (pq PriorityQueue) Len() int {
	return len(pq.array)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq.array[i].Less(pq.array[j])
}

func (pq PriorityQueue) Swap(i, j int) {
	pq.array[i], pq.array[j] = pq.array[j], pq.array[i]
	pq.indexMap[pq.array[i]] = i
	pq.indexMap[pq.array[j]] = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(pq.array)
	item := x.(Node)
	pq.indexMap[item] = n
	pq.array = append(pq.array, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := pq.array
	n := len(old)
	item := old[n-1]
	old[n-1] = nil            // avoid memory leak
	delete(pq.indexMap, item) // for safety
	pq.array = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Update(node Node) {
	heap.Fix(pq, pq.indexMap[node])
}
