package copy_list_with_random_pointer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCaseDataSet struct {
	answer [][]any
	list   [][]any
}

func TestCopyRandomList(t *testing.T) {

	for _, dataSet := range getTestDataSet() {

		list := copyRandomList(toList(dataSet.list))
		for _, pair := range dataSet.answer {

			assert.Equal(t, pair[0], list.Val)

			if list.Random == nil {
				assert.Nil(t, list.Random)
			} else {
				assert.Equal(t, pair[1], list.Random.Val)
			}

			list = list.Next
		}
	}
}

func getTestDataSet() []TestCaseDataSet {

	return []TestCaseDataSet{
		{[][]any{{7, nil}}, [][]any{{7, nil}}},
		{[][]any{{7, nil}, {13, 7}, {11, 1}, {10, 11}, {1, 7}}, [][]any{{7, nil}, {13, 0}, {11, 4}, {10, 2}, {1, 0}}},
		{[][]any{{2, 2}}, [][]any{{2, 0}}},
	}
}

func toList(data [][]any) *Node {

	if len(data) == 0 {
		return nil
	}

	nodes := make([]*Node, len(data))
	for i, pair := range data {
		val := pair[0].(int)
		nodes[i] = &Node{Val: val}
	}

	for i, pair := range data {
		if i < len(nodes)-1 {
			nodes[i].Next = nodes[i+1]
		}
		randomIndex := pair[1]
		if randomIndex != nil {
			nodes[i].Random = nodes[randomIndex.(int)]
		}
	}

	return nodes[0]
}
