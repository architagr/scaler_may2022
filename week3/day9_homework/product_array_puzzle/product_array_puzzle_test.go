package product_array_puzzle

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	inputArr []int
	expected []int
}

func TestProductArrayPuzzle(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputArr: []int{1, 2, 3, 4, 5},
		expected: []int{120, 60, 40, 30, 24},
	})

	testCases = append(testCases, TestCase{
		inputArr: []int{5, 1, 10, 1},
		expected: []int{10, 50, 5, 50},
	})

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %+v", testCase.inputArr), func(tb *testing.T) {
			got := ProductArrayPuzzle(testCase.inputArr)
			if !reflect.DeepEqual(got, testCase.expected) {
				tb.Errorf("Tested %+v extected %+v and got %+v", testCase.inputArr, testCase.expected, got)
			}
		})
	}
}
