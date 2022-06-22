package toupper

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	input    []byte
	expected []byte
}

func TestToUpper(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    []byte{'S', 'c', 'A', 'L', 'E', 'r', 'A', 'c', 'a', 'D', 'e', 'm', 'y'},
		expected: []byte{'S', 'C', 'A', 'L', 'E', 'R', 'A', 'C', 'A', 'D', 'E', 'M', 'Y'},
	})
	testCases = append(testCases, TestCase{
		input:    []byte{'S', 'c', 'a', 'L', 'e', 'R', '#', '2', '0', '2', '0'},
		expected: []byte{'S', 'C', 'A', 'L', 'E', 'R', '#', '2', '0', '2', '0'},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := ToUpper(testCase.input)
			if !reflect.DeepEqual(got, testCase.expected) {
				tb.Errorf("tested %d expected %d but got %d", (i + 1), testCase.expected, got)
			}
		})
	}
}
