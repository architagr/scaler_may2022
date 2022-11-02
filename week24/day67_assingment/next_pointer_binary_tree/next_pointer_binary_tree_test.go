package next_pointer_binary_tree

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	head     *treeNode
	expected []int
}

func TestConnect(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		head:     GetTestCase1(),
		expected: []int{4, 5, 6, 7},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := Connect(testCase.head)

			if !reflect.DeepEqual(InorderTraversal(x), testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}
func InorderTraversal(A *treeNode) []int {
	return Parse(A)
}

func Parse(head *treeNode) []int {
	if head == nil {
		return []int{}
	}
	result := make([]int, 0)

	result = append(result, Parse(head.left)...)
	result = append(result, head.value)
	result = append(result, Parse(head.right)...)
	return result
}

func GetTestCase1() *treeNode {
	node11 := treeNode_new(4)

	node211 := treeNode_new(5)
	node221 := treeNode_new(6)
	node11.left = node211
	node11.right = node221

	node311 := treeNode_new(7)
	node211.right = node311

	return node11
}
