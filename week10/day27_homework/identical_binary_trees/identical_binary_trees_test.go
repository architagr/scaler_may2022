package identical_binary_trees

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	inputA   *treeNode
	inputB   *treeNode
	expected int
}

func TestIsSameTree(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputA:   GetTestCase1(),
		inputB:   GetTestCase1(),
		expected: 1,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := isSameTree(testCase.inputA, testCase.inputB)

			if !reflect.DeepEqual(got, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}

func GetTestCase1() *treeNode {

	var node21 *treeNode = new(treeNode)
	node21.value = 3

	var node1 *treeNode = new(treeNode)
	node1.value = 1
	node1.right = node21

	return node1
}
