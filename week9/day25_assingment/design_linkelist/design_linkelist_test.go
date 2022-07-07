package design_linkelist

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	input    [][]int
	expected []int
}

func TestInsertNode(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input: [][]int{{1, 13, -1},
			{3, 0, -1},
			{3, 1, -1},
			{2, 15, 0},
			{3, 0, -1},
			{1, 12, -1},
			{3, 0, -1},
			{1, 19, -1},
			{1, 13, -1},
			{3, 0, -1},
			{0, 12, -1},
			{1, 13, -1},
			{3, 2, -1},
		},
		expected: []int{12, 13},
	})

	testCases = append(testCases, TestCase{
		input: [][]int{
			{3, 1, -1},
			{3, 1, -1},
			{1, 18, -1},
			{2, 12, 1},
			{1, 17, -1},
			{2, 11, 3},
			{1, 19, -1},
			{3, 0, -1},
			{0, 12, -1},
		},
		expected: []int{12, 12, 17, 11, 19},
	})

	testCases = append(testCases, TestCase{
		input: [][]int{{0, 1, -1},
			{1, 2, -1},
			{2, 3, 1},
			{0, 4, -1},
			{3, 1, -1},
			{3, 2, -1},
		},
		expected: []int{4, 3},
	})

	testCases = append(testCases, TestCase{
		input: [][]int{{2, 18, 0},
			{2, 5, 1},
			{2, 8, 0},
			{1, 7, -1},
			{1, 5, -1},
		},
		expected: []int{8, 18, 5, 7, 5},
	})

	testCases = append(testCases, TestCase{
		input: [][]int{{0, 1, -1},
			{1, 2, -1},
			{2, 3, 1}},
		expected: []int{1, 3, 2},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := solve(testCase.input)
			x := GetLitListToArray(got)

			if !reflect.DeepEqual(x, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}
func GetLitListToArray(head *listNode) []int {
	arr := make([]int, 0)

	for head != nil {
		arr = append(arr, head.value)
		head = head.next
	}
	return arr
}
