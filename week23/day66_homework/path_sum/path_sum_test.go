package path_sum

import (
	"fmt"
	"testing"
)

type TestCase struct {
	head     *treeNode
	b        int
	expected int
}

func TestPathSum(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		head:     GetTestCase1(),
		b:        22,
		expected: 1,
	})

	testCases = append(testCases, TestCase{
		head:     GetTestCase2(),
		b:        1000,
		expected: 0,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := HasPathSum(testCase.head, testCase.b)

			if x != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}
func GetTestCase2() *treeNode {
	node1 := treeNode_new(1000)

	node21 := treeNode_new(200)
	node1.right = node21

	return node1
}
func GetTestCase1() *treeNode {
	node1 := treeNode_new(5)

	node21 := treeNode_new(4)
	node22 := treeNode_new(8)
	node1.left = node21
	node1.right = node22

	node31 := treeNode_new(11)
	node21.left = node31

	node33 := treeNode_new(13)
	node34 := treeNode_new(4)
	node22.left = node33
	node22.right = node34

	node41 := treeNode_new(2)
	node42 := treeNode_new(2)
	node31.right = node41
	node31.left = node42

	node43 := treeNode_new(1)
	node34.right = node43

	return node1
}
