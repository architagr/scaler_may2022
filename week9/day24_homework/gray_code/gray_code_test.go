package gray_code

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	input    int
	expected []int
}

func TestGetGrayCode(t *testing.T) {

	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    2,
		expected: []int{0, 1, 3, 2},
	})
	testCases = append(testCases, TestCase{
		input:    3,
		expected: []int{0, 1, 3, 2, 6, 7, 5, 4},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := GetGrayCode(testCase.input)
			if !reflect.DeepEqual(got, testCase.expected) {
				tb.Errorf("tested %d expected %d but got %d", (i + 1), testCase.expected, got)
			}
		})
	}
}
