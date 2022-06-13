package single_number_3

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	input, expected []int
}

func TestFundSingleNumer(t *testing.T) {
	testcases := make([]TestCase, 0)
	testcases = append(testcases, TestCase{
		input:    []int{1, 2, 3, 1, 2, 4},
		expected: []int{3, 4},
	})

	testcases = append(testcases, TestCase{
		input:    []int{1, 2},
		expected: []int{1, 2},
	})
	testcases = append(testcases, TestCase{
		input:    []int{186, 256, 102, 377, 186, 377},
		expected: []int{102, 256},
	})

	for _, testCase := range testcases {
		t.Run(fmt.Sprintf("finding for %+v", testCase.input), func(tb *testing.T) {
			got := FundSingleNumer(testCase.input)

			if !reflect.DeepEqual(got, testCase.expected) {
				tb.Errorf("Tested for %+v, expected %+v but got %+v", testCase.input, testCase.expected, got)
			}
		})
	}
}
