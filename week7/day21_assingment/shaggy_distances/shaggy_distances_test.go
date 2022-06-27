package shaggy_distances

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	input    []int
	expected int
}

func TestShaggyDistances(t *testing.T) {

	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    []int{7, 1, 3, 4, 1, 7},
		expected: 3,
	})
	testCases = append(testCases, TestCase{
		input:    []int{1, 1},
		expected: 1,
	})
	testCases = append(testCases, TestCase{
		input:    []int{7, 1, 3, 4, 2, 8},
		expected: -1,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := ShaggyDistances(testCase.input)
			if !reflect.DeepEqual(got, testCase.expected) {
				tb.Errorf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}

}
