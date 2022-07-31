package sub_matrix_sum_queries

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	input                [][]int
	b, c, d, e, expected []int
}

func TestSubMatrixSumQueries(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		b:        []int{1, 2},
		c:        []int{1, 2},
		d:        []int{2, 3},
		e:        []int{2, 3},
		expected: []int{12, 28},
	})

	testCases = append(testCases, TestCase{
		input:    [][]int{{5, 17, 100, 11}, {0, 0, 2, 8}},
		b:        []int{1, 1},
		c:        []int{1, 4},
		d:        []int{2, 2},
		e:        []int{2, 4},
		expected: []int{22, 19},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d ", (i+1)), func(tb *testing.T) {
			got := SubMatrixSumQueries(testCase.input, testCase.b, testCase.c, testCase.d, testCase.e)

			if !reflect.DeepEqual(got, testCase.expected) {
				tb.Errorf("testing %d, expected %+v, got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
