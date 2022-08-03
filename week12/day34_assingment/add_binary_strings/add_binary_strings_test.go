package add_binary_strings

import (
	"fmt"
	"testing"
)

type TestCase struct {
	a, b, expected string
}

func TestAddBinaryStrings(t *testing.T) {
	testCases := make([]TestCase, 0)
	// testCases = append(testCases, TestCase{
	// 	a:        "100",
	// 	b:        "11",
	// 	expected: "111",
	// })
	// testCases = append(testCases, TestCase{
	// 	a:        "00",
	// 	b:        "11",
	// 	expected: "11",
	// })

	testCases = append(testCases, TestCase{
		a:        "1010110111001101101000",
		b:        "1000011011000000111100110",
		expected: "1001110001111010101001110",
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d ", (i+1)), func(tb *testing.T) {
			got := AddBinaryStrings(testCase.a, testCase.b)

			if got != testCase.expected {
				tb.Errorf("testing %d, expected %+v, got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
