package binary_tree_from_inorder_and_postorder

import (
	"fmt"
	"reflect"
	"scaler-may-22/week22/day62_assingment/inorder_traversal"
	"testing"
)

type TestCase struct {
	inOrder, postOrder []int
	expected           []int
}

func TestGetNumbers(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		inOrder:   []int{2, 1, 3},
		postOrder: []int{2, 3, 1},
		expected:  []int{2, 1, 3},
	})

	testCases = append(testCases, TestCase{
		inOrder:   []int{7, 5, 6, 2, 3, 1, 4},
		postOrder: []int{5, 6, 3, 1, 4, 2, 7},
		expected:  []int{7, 5, 6, 2, 3, 1, 4},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := inorder_traversal.InorderTraversal(BuildTree(testCase.inOrder, testCase.postOrder))

			if !reflect.DeepEqual(x, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}
