package range_sum_query

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCasePrefixSum struct {
	inputArr []int
	expected []int64
}

func TestFindPrefixSum(t *testing.T) {
	testCases := make([]TestCasePrefixSum, 0)
	testCases = append(testCases, TestCasePrefixSum{
		inputArr: []int{1, 2, 3, 4, 5, 6},
		expected: []int64{1, 3, 6, 10, 15, 21},
	})

	testCases = append(testCases, TestCasePrefixSum{
		inputArr: []int{1, 2, 3, 4, -5, 6},
		expected: []int64{1, 3, 6, 10, 5, 11},
	})

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %+v", testCase.inputArr), func(tb *testing.T) {
			got := FindPrefixSum(testCase.inputArr)
			if !reflect.DeepEqual(got, testCase.expected) {
				tb.Errorf("Tested %+v extected %+v and got %+v", testCase.inputArr, testCase.expected, got)
			}
		})
	}
}

type TestCaseRangeQuery struct {
	inputArr []int
	queries  [][]int
	expected []int64
}

func TestFindSumRange(t *testing.T) {
	testCases := make([]TestCaseRangeQuery, 0)
	testCases = append(testCases, TestCaseRangeQuery{
		inputArr: []int{1, 2, 3, 4, 5, 6},
		queries:  [][]int{{1, 4}, {2, 3}},
		expected: []int64{10, 5},
	})

	testCases = append(testCases, TestCaseRangeQuery{
		inputArr: []int{2, 2, 2},
		queries:  [][]int{{1, 1}, {2, 3}},
		expected: []int64{2, 4},
	})

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %+v", testCase.inputArr), func(tb *testing.T) {
			got := FindSumRange(testCase.inputArr, testCase.queries)
			if !reflect.DeepEqual(got, testCase.expected) {
				tb.Errorf("Tested %+v extected %+v and got %+v", testCase.inputArr, testCase.expected, got)
			}
		})
	}
}
