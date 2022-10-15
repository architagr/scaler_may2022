package serialize_binary_tree

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	head     *treeNode
	expected []int
}

func TestSerializeBinaryTree(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		head:     GetTestCase1(),
		expected: []int{1, 2, 3, 4, 5, -1, -1, -1, -1, -1, -1},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := SerializeBinaryTree(testCase.head)

			if !reflect.DeepEqual(x, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}

func GetTestCase1() *treeNode {
	node11 := treeNode_new(1)
	node12 := treeNode_new(2)
	node13 := treeNode_new(3)
	node14 := treeNode_new(4)
	node15 := treeNode_new(5)

	node11.left = node12
	node11.right = node13

	node12.left = node14
	node12.right = node15

	return node11
}
