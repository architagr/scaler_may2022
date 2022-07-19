package path_to_given_node

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	inputA   *treeNode
	inputB   int
	expected []int
}

func TestFindPath(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputA:   GetTestCase1(),
		inputB:   5,
		expected: []int{1, 2, 5},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := FindPath(testCase.inputA, testCase.inputB)

			if !reflect.DeepEqual(got, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}

func GetTestCase1() *treeNode {

	var node31 *treeNode = new(treeNode)
	node31.value = 4

	var node32 *treeNode = new(treeNode)
	node32.value = 5

	var node33 *treeNode = new(treeNode)
	node33.value = 6

	var node34 *treeNode = new(treeNode)
	node34.value = 7

	var node21 *treeNode = new(treeNode)
	node21.value = 2
	node21.left = node31
	node21.right = node32

	var node22 *treeNode = new(treeNode)
	node22.value = 3
	node22.left = node33
	node22.right = node34

	var node1 *treeNode = new(treeNode)
	node1.value = 1
	node1.left = node21
	node1.right = node22

	return node1
}
