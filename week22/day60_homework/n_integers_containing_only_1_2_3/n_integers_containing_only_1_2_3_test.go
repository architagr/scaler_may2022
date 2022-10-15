package n_integers_containing_only_1_2_3

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	input    int
	expected []int
}

func TestGetNumbers(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		input:    3,
		expected: []int{1, 2, 3},
	})

	testCases = append(testCases, TestCase{
		input:    7,
		expected: []int{1, 2, 3, 11, 12, 13, 21},
	})

	testCases = append(testCases, TestCase{
		input:    18,
		expected: []int{1, 2, 3, 11, 12, 13, 21, 22, 23, 31, 32, 33, 111, 112, 113, 121, 122, 123},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := GetNumbers(testCase.input)

			if !reflect.DeepEqual(x, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}
