package maxminarray

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputArr                 []int
	expectedMax, expectedMin int
}

func TestMaxMinArray(t *testing.T) {
	var testcases = make([]TestCase, 0)
	testcases = append(testcases, TestCase{
		inputArr:    []int{1, 2, 3, 4, 5},
		expectedMax: 5,
		expectedMin: 1,
	})

	testcases = append(testcases, TestCase{
		inputArr:    []int{10, 50, 40, 80},
		expectedMax: 80,
		expectedMin: 10,
	})

	for _, testcase := range testcases {
		t.Run(fmt.Sprintf("testing %v", testcase.inputArr), func(tb *testing.T) {
			gotMax, gotMin := MaxMinArray(testcase.inputArr)
			if gotMax != testcase.expectedMax || gotMin != testcase.expectedMin {
				tb.Errorf("Tested %+v, expected max %d but got %d, expected min %d but got %d", testcase.inputArr, testcase.expectedMax, gotMax, testcase.expectedMin, gotMin)
			}
		})
	}

}
