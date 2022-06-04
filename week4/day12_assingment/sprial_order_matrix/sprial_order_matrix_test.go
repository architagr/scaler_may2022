package sprial_order_matrix

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	input    int
	expected [][]int
}

func TestGetSprialOrderMatrix(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    4,
		expected: [][]int{{1, 2, 3, 4}, {12, 13, 14, 5}, {11, 16, 15, 6}, {10, 9, 8, 7}},
	})
	testCases = append(testCases, TestCase{
		input:    5,
		expected: [][]int{{1, 2, 3, 4, 5}, {16, 17, 18, 19, 6}, {15, 24, 25, 20, 7}, {14, 23, 22, 21, 8}, {13, 12, 11, 10, 9}},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing case %d", (i+1)), func(tb *testing.T) {
			got := GetSprialOrderMatrix(testCase.input)
			if !reflect.DeepEqual(got, testCase.expected) {
				tb.Errorf("tested case no %d expected %+v, but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
