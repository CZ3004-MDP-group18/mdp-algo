package common

import (
	"container/heap"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestNode struct {
	cost int
}

func (t *TestNode) Less(o Node) bool {
	return t.cost < o.(*TestNode).cost
}

func TestHeap(t *testing.T) {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{[]int{1}, []int{1}},
		{[]int{1, 5, 3}, []int{1, 3, 5}},
	}
	for _, testCase := range testCases {
		h := NewHeap()
		for input := range testCase.input {
			h.Push(&TestNode{cost: input})
		}
		for expected := range testCase.expected {
			assert.Equal(t, expected, h.Pop().(*TestNode).cost)
		}
	}
}

func TestHeapify(t *testing.T) {
	inputs := []int{1, 5, 3, 9, 10}
	var nodes []*TestNode
	for _, input := range inputs {
		nodes = append(nodes, &TestNode{cost: input})
	}

	h := NewHeap()
	for _, node := range nodes {
		h.Push(node)
	}

	// Change node with cost 5 -> 2 and cost 9 -> 11
	nodes[1].cost = 2
	nodes[3].cost = 11
	h.Heapify()

	expecteds := []int{1, 2, 3, 10, 11}
	for _, expected := range expecteds {
		assert.Equal(t, expected, h.Pop().(*TestNode).cost)
	}
}

func TestPriorityQueue(t *testing.T) {
	inputs := []int{1, 5, 3, 9, 10}
	var nodes []*TestNode
	for _, input := range inputs {
		nodes = append(nodes, &TestNode{cost: input})
	}

	pq := NewPriorityQueue()
	heap.Init(pq)
	for _, node := range nodes {
		heap.Push(pq, node)
	}

	// Change node with cost 5 -> 2 and cost 9 -> 11
	nodes[1].cost = 2
	pq.Update(nodes[1])
	nodes[3].cost = 11
	pq.Update(nodes[3])

	expecteds := []int{1, 2, 3, 10, 11}
	for _, expected := range expecteds {
		assert.Equal(t, expected, heap.Pop(pq).(*TestNode).cost)
	}
}
