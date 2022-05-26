package separate_even_odd

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	inputArr, expectedEven, expectedOdd []int
}

func TestSeparateEvenOdd(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputArr:     []int{1, 2, 3, 4, 5},
		expectedEven: []int{2, 4},
		expectedOdd:  []int{1, 3, 5},
	})

	for _, testcase := range testCases {
		t.Run(fmt.Sprintf("testing %v", testcase.inputArr), func(tb *testing.T) {
			gotEven, gotOdd := SeparateEvenOdd(testcase.inputArr)
			if !reflect.DeepEqual(gotEven, testcase.expectedEven) || !reflect.DeepEqual(gotOdd, testcase.expectedOdd) {
				tb.Errorf("Tested %+v, even expected %+v but got %+v, odd expected %+v but got %+v", testcase.inputArr, testcase.expectedEven, gotEven, testcase.expectedOdd, gotOdd)
			}
		})
	}
}
