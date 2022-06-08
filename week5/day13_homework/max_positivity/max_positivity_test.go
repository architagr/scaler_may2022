package max_positivity

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	input    []int
	expected []int
}

func TestFindMaxPositivitySubArray(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    []int{-5173778, -8176176, 1694510, 7089884, -1394259, 1146372, -2502339, 1544618, 6602022, 4330572},
		expected: []int{1544618, 6602022, 4330572},
	})
	testCases = append(testCases, TestCase{
		input:    []int{5, 6, -1, 7, 8},
		expected: []int{5, 6},
	})
	testCases = append(testCases, TestCase{
		input:    []int{5, 6, -1, 7, 8, 9},
		expected: []int{7, 8, 9},
	})
	testCases = append(testCases, TestCase{
		input:    []int{1, 2, 3, 4, 5, 6},
		expected: []int{1, 2, 3, 4, 5, 6},
	})
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing case no %d", (i+1)), func(tb *testing.T) {
			got := FindMaxPositivitySubArray(testCase.input)
			if !reflect.DeepEqual(got, testCase.expected) {
				tb.Errorf("tested case number %d, expected %+v and got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
