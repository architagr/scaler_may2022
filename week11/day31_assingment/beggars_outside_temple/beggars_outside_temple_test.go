package beggars_outside_temple

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	input    int
	b        [][]int
	expected []int
}

func TestBeggarsOutsideTemple(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    5,
		b:        [][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}},
		expected: []int{10, 55, 45, 25, 25},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d ", (i+1)), func(tb *testing.T) {
			got := BeggarsOutsideTemple(testCase.input, testCase.b)

			if !reflect.DeepEqual(got, testCase.expected) {
				tb.Errorf("testing %d, expected %+v, got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
