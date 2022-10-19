package balanced_binary_tree

import (
	"fmt"
	"testing"
)

type TestCase struct {
	head     *treeNode
	expected int
}

func TestReorderList(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		head:     GetTestCase1(),
		expected: 0,
	})

	testCases = append(testCases, TestCase{
		head:     GetTestCase2(),
		expected: 1,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := isBalanced(testCase.head)

			if x != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}
func GetTestCase2() *treeNode {
	node1 := treeNode_new(1)
	node21 := treeNode_new(2)
	node1.right = node21

	return node1
}

func GetTestCase1() *treeNode {
	node1 := treeNode_new(25)

	node21 := treeNode_new(7)
	node22 := treeNode_new(10)
	node1.left = node21
	node1.right = node22

	node31 := treeNode_new(5)
	node32 := treeNode_new(9)
	node21.left = node31
	node21.right = node32

	node33 := treeNode_new(12)
	node34 := treeNode_new(2)
	node22.left = node33
	node22.right = node34

	node41 := treeNode_new(11)
	node31.right = node41

	node42 := treeNode_new(1)
	node43 := treeNode_new(8)
	node33.left = node42
	node33.right = node43

	node44 := treeNode_new(4)
	node34.right = node44

	node51 := treeNode_new(6)
	node52 := treeNode_new(3)
	node41.left = node51
	node41.right = node52

	return node1
}
