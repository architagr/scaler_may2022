package max_non_negative_subArray

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	input    []int
	expected []int
}

func TestMaxSet(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    []int{1, 2, 3},
		expected: []int{1, 2, 3},
	})
	testCases = append(testCases, TestCase{
		input:    []int{1, 2, 5, -7, 2, 3},
		expected: []int{1, 2, 5},
	})
	testCases = append(testCases, TestCase{
		input:    []int{10, -1, 2, 3, -4, 100},
		expected: []int{100},
	})

	testCases = append(testCases, TestCase{
		input:    []int{24831, 53057, 19790, -10679, 77006, -36098, 75319, -45223, -56760, -86126, -29506, 76770, 29094, 87844, 40534, -15394, 18738, 59590, -45159, -64947, 7217, -55539, 42385, -94885, 82420, -31968, -99915, 63534, -96011, 24152, 77295},
		expected: []int{76770, 29094, 87844, 40534},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d ", (i+1)), func(tb *testing.T) {
			got := MaxSet(testCase.input)

			if !reflect.DeepEqual(got, testCase.expected) {
				tb.Errorf("testing %d, expected %+v, got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
