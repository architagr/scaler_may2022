package shortest_unique_prefix

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	head, expected []string
}

func TestPrefix(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		head:     []string{"zebra", "dog", "duck", "dove"},
		expected: []string{"z", "dog", "du", "dov"},
	})

	testCases = append(testCases, TestCase{
		head:     []string{"apple", "ball", "cat"},
		expected: []string{"a", "b", "c"},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := Prefix(testCase.head)

			if !reflect.DeepEqual(x, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}
