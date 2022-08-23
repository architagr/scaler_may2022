package tower_of_hanoi

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	a        int
	expected [][]int
}

func TestTowerOfHanoi(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		a:        2,
		expected: [][]int{{1, 1, 2}, {2, 1, 3}, {1, 2, 3}},
	})

	testCases = append(testCases, TestCase{
		a:        3,
		expected: [][]int{{1, 1, 3}, {2, 1, 2}, {1, 3, 2}, {3, 1, 3}, {1, 2, 1}, {2, 2, 3}, {1, 1, 3}},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d ", (i+1)), func(tb *testing.T) {
			got := TowerOfHanoi(testCase.a)

			if !reflect.DeepEqual(got, testCase.expected) {
				tb.Errorf("testing %d, expected %+v, got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
