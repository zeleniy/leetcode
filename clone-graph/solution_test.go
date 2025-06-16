package clone_graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Тестовые данные
type TestCaseDataSet struct {
	answer [][]int
	node   *Node
}

func BenchmarkCloneGraph(b *testing.B) {

	for i := 0; i < b.N; i++ {
		for _, data := range getTestDataSet() {
			cloneGraph(data.node)
		}
	}
}

func TestCloneGraph(t *testing.T) {
	for _, data := range getTestDataSet() {

		head := cloneGraph(data.node)

		if data.answer == nil {
			assert.Nil(t, head)
		} else {
			assert.Equal(t, len(data.answer[0]), len(head.Neighbors))
		}
	}
}

func getTestDataSet() []TestCaseDataSet {

	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}

	node1.Neighbors = []*Node{node2, node4}
	node2.Neighbors = []*Node{node1, node3}
	node3.Neighbors = []*Node{node2, node4}
	node4.Neighbors = []*Node{node1, node3}

	singleNode := &Node{Val: 1}

	return []TestCaseDataSet{
		{
			answer: [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}},
			node:   node1,
		},
		{
			answer: [][]int{{}},
			node:   singleNode,
		},
		{
			answer: nil,
			node:   nil,
		},
	}
}
