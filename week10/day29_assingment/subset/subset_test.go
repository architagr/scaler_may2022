package subset

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	inputA   []int
	expected [][]int
}

func TestGetSubset(t *testing.T) {

	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputA:   []int{13, 12},
		expected: [][]int{{}, {12}, {12, 13}, {13}},
	})
	testCases = append(testCases, TestCase{
		inputA:   []int{1, 2, 3},
		expected: [][]int{{}, {1}, {1, 2}, {1, 2, 3}, {1, 3}, {2}, {2, 3}, {3}},
	})

	testCases = append(testCases, TestCase{
		inputA: []int{15, 20, 12, 19, 4},
		expected: [][]int{{}, {4}, {4, 12}, {4, 12, 15}, {4, 12, 15, 19}, {4, 12, 15, 19, 20}, {4, 12, 15, 20},
			{4, 12, 19}, {4, 12, 19, 20}, {4, 12, 20}, {4, 15}, {4, 15, 19}, {4, 15, 19, 20}, {4, 15, 20},
			{4, 19}, {4, 19, 20}, {4, 20}, {12}, {12, 15}, {12, 15, 19}, {12, 15, 19, 20}, {12, 15, 20},
			{12, 19}, {12, 19, 20}, {12, 20}, {15}, {15, 19}, {15, 19, 20}, {15, 20}, {19}, {19, 20}, {20}},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := GetSubset(testCase.inputA)
			if !reflect.DeepEqual(got, testCase.expected) {
				tb.Errorf("tested %d \nexpected\n%+v but \ngot\n%+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
