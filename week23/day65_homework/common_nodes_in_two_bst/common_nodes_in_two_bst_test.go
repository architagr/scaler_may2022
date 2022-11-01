package common_nodes_in_two_bst

import (
	"fmt"
	"testing"
)

type TestCase struct {
	headA, headB *treeNode
	expected     int
}

func TestCommonNodes(t *testing.T) {
	testCases := make([]TestCase, 0)

	headA, headB := GetTestCase1()
	testCases = append(testCases, TestCase{
		headA:    headA,
		headB:    headB,
		expected: 17,
	})

	head2A, head2B := GetTestCase2()
	testCases = append(testCases, TestCase{
		headA:    head2A,
		headB:    head2B,
		expected: 46,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := CommonNodes(testCase.headA, testCase.headB)

			if x != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}

func GetTestCase2() (*treeNode, *treeNode) {
	node11 := treeNode_new(7)

	node211 := treeNode_new(1)
	node221 := treeNode_new(10)
	node11.left = node211
	node11.right = node221

	node311 := treeNode_new(2)
	node211.right = node311

	node321 := treeNode_new(15)
	node221.left = node321

	node411 := treeNode_new(11)
	node321.left = node411

	// 2nd tree
	node12 := treeNode_new(7)

	node212 := treeNode_new(1)
	node222 := treeNode_new(10)
	node12.left = node212
	node12.right = node222

	node312 := treeNode_new(2)
	node212.right = node312

	node322 := treeNode_new(15)
	node222.left = node322

	node412 := treeNode_new(11)
	node322.left = node412

	return node11, node12
}

func GetTestCase1() (*treeNode, *treeNode) {
	node11 := treeNode_new(5)

	node211 := treeNode_new(2)
	node221 := treeNode_new(8)
	node11.left = node211
	node11.right = node221

	node311 := treeNode_new(3)
	node211.right = node311

	node321 := treeNode_new(15)
	node221.left = node321

	node411 := treeNode_new(9)
	node321.left = node411

	// 2nd tree
	node12 := treeNode_new(7)

	node212 := treeNode_new(1)
	node222 := treeNode_new(10)
	node12.left = node212
	node12.right = node222

	node312 := treeNode_new(2)
	node212.right = node312

	node322 := treeNode_new(15)
	node222.left = node322

	node412 := treeNode_new(11)
	node322.left = node412

	return node11, node12
}
