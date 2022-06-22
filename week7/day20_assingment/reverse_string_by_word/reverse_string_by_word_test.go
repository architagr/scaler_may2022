package reverse_string_by_word

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	inputArr string
	expected string
}

func TestReverseStringByWord(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputArr: "the sky is blue",
		expected: "blue is sky the",
	})

	testCases = append(testCases, TestCase{
		inputArr: "this is ib",
		expected: "ib is this",
	})
	testCases = append(testCases, TestCase{
		inputArr: "crulgzfkif gg ombt vemmoxrgf qoddptokkz  op xdq hv ",
		expected: "hv xdq op qoddptokkz vemmoxrgf ombt gg crulgzfkif",
	})
	for _, testcase := range testCases {
		t.Run(fmt.Sprintf("testing %v", testcase.inputArr), func(tb *testing.T) {
			got := ReverseStringByWord(testcase.inputArr)
			if !reflect.DeepEqual(got, testcase.expected) {
				tb.Errorf("Tested %+v to  expected %+v but got %+v", testcase.inputArr, testcase.expected, got)
			}
		})
	}
}
