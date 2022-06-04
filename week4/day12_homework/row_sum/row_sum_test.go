package row_sum

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	input    [][]int
	expected []int
}

func TestFindRowSum(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 2, 3, 4}},
		expected: []int{10, 26, 18},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing case no %d", (i+1)), func(tb *testing.T) {
			got := FindRowSum(testCase.input)
			if !reflect.DeepEqual(got, testCase.expected) {
				tb.Errorf("tested case number %d, expected %+v and got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
