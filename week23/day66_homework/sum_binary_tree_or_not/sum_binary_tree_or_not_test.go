package sum_binary_tree_or_not

import (
	"fmt"
	"testing"
)

type TestCase struct {
	head     *treeNode
	expected int
}

func TestIsSumBinaryTree(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		head:     GetTestCase1(),
		expected: 1,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := IsSumBinaryTree(testCase.head)

			if x != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}

func GetTestCase1() *treeNode {
	node11 := treeNode_new(10)

	node211 := treeNode_new(4)
	node221 := treeNode_new(2)
	node11.left = node211
	node11.right = node221

	node311 := treeNode_new(2)
	node321 := treeNode_new(2)
	node211.right = node311
	node211.left = node321

	return node11
}
