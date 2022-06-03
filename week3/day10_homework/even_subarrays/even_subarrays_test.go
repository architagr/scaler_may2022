package even_subarrays

import (
	"fmt"
	"testing"
)

func TestFindEvenSubarray_OddLength(t *testing.T) {
	expected := "NO"

	got := FindEvenSubarray([]int{1, 2, 3})
	if got != expected {
		t.Errorf("for array of length odd the expected is %s but got %s", expected, got)
	}
}

type TestCase struct {
	input    []int
	expected string
}

func TestFindEvenSubarray_EvenLength(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    []int{2, 4, 8, 6},
		expected: "YES",
	})
	testCases = append(testCases, TestCase{
		input:    []int{2, 4, 8, 7, 6},
		expected: "NO",
	})
	testCases = append(testCases, TestCase{
		input:    []int{796, 483, 671, 998, 149, 374, 801, 191, 457, 802, 903, 329, 191, 286, 153, 600, 917, 189, 724, 741, 306, 253, 520, 372},
		expected: "YES",
	})
	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %+v", testCase.input), func(tb *testing.T) {
			got := FindEvenSubarray(testCase.input)
			if got != testCase.expected {
				tb.Errorf("for array %+v the expected is %s but got %s", testCase.input, testCase.expected, got)
			}
		})
	}

}
