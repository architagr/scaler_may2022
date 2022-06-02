package sum_of_even_odd_index

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	inputArr []int
	expected int
}

func TestSumOfEvenOddIndex(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputArr: []int{2, 1, 6, 4},
		expected: 1,
	})

	testCases = append(testCases, TestCase{
		inputArr: []int{1, 1, 1},
		expected: 3,
	})

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %+v", testCase.inputArr), func(tb *testing.T) {
			got := SumOfEvenOddIndex(testCase.inputArr)
			if got != testCase.expected {
				tb.Errorf("Tested %+v extected %d and got %d", testCase.inputArr, testCase.expected, got)
			}
		})
	}
}

type TestCasePrefixSum struct {
	inputArr []int
	expected []int64
}

func TestFindEvenPrefixSum(t *testing.T) {
	testCases := make([]TestCasePrefixSum, 0)
	testCases = append(testCases, TestCasePrefixSum{
		inputArr: []int{1, 2, 3, 4, 5, 6},
		expected: []int64{1, 1, 4, 4, 9, 9},
	})

	testCases = append(testCases, TestCasePrefixSum{
		inputArr: []int{1, 2, 3, 4, -5, 6},
		expected: []int64{1, 1, 4, 4, -1, -1},
	})

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %+v", testCase.inputArr), func(tb *testing.T) {
			got := FindEvenPrefixSum(testCase.inputArr)
			if !reflect.DeepEqual(got, testCase.expected) {
				tb.Errorf("Tested %+v extected %+v and got %+v", testCase.inputArr, testCase.expected, got)
			}
		})
	}
}

func TestFindOddPrefixSum(t *testing.T) {
	testCases := make([]TestCasePrefixSum, 0)
	testCases = append(testCases, TestCasePrefixSum{
		inputArr: []int{1, 2, 3, 4, 5, 6},
		expected: []int64{0, 2, 2, 6, 6, 12},
	})

	testCases = append(testCases, TestCasePrefixSum{
		inputArr: []int{1, 2, 3, 4, -5, 6},
		expected: []int64{0, 2, 2, 6, 6, 12},
	})

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %+v", testCase.inputArr), func(tb *testing.T) {
			got := FindOddPrefixSum(testCase.inputArr)
			if !reflect.DeepEqual(got, testCase.expected) {
				tb.Errorf("Tested %+v extected %+v and got %+v", testCase.inputArr, testCase.expected, got)
			}
		})
	}
}
