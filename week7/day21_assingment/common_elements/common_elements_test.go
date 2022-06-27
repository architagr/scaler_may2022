package common_elements

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	inputA, inputB, expected []int
}

func TestFindCommonELements(t *testing.T) {

	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputA:   []int{1, 2, 2, 1},
		inputB:   []int{2, 3, 1, 2},
		expected: []int{1, 2, 2},
	})
	testCases = append(testCases, TestCase{
		inputA:   []int{2, 1, 4, 10},
		inputB:   []int{3, 6, 2, 10, 10},
		expected: []int{2, 10},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := FindCommonELements(testCase.inputA, testCase.inputB)
			if !reflect.DeepEqual(got, testCase.expected) {
				tb.Errorf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}

}
