package add_one_to_number

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	input    []int
	expected []int
}

func TestPlusOne(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    []int{1, 2, 3},
		expected: []int{1, 2, 4},
	})
	testCases = append(testCases, TestCase{
		input:    []int{9, 9},
		expected: []int{1, 0, 0},
	})
	testCases = append(testCases, TestCase{
		input:    []int{7, 9},
		expected: []int{8, 0},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d ", (i+1)), func(tb *testing.T) {
			got := PlusOne(testCase.input)

			if !reflect.DeepEqual(got, testCase.expected) {
				tb.Errorf("testing %d, expected %+v, got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
