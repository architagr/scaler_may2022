package palindrome_list

import (
	"fmt"
	"testing"
)

type TestCase struct {
	head     *listNode
	expected int
}

func TestFindPalindrome(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		head:     GetTestCase1(),
		expected: 0,
	})
	testCases = append(testCases, TestCase{
		head:     GetTestCase2(),
		expected: 0,
	})

	testCases = append(testCases, TestCase{
		head:     GetTestCase3(),
		expected: 1,
	})
	testCases = append(testCases, TestCase{
		head:     GetTestCase4(),
		expected: 1,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := FindPalindrome(testCase.head)

			if x != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}

func GetTestCase1() *listNode {
	node11 := listNode_new(1)
	node12 := listNode_new(2)
	node13 := listNode_new(3)
	node11.next = node12
	node12.next = node13

	return node11
}

func GetTestCase2() *listNode {
	node11 := listNode_new(1)
	node12 := listNode_new(2)
	node13 := listNode_new(3)
	node14 := listNode_new(4)

	node11.next = node12
	node12.next = node13
	node13.next = node14

	return node11

}
func GetTestCase3() *listNode {
	node11 := listNode_new(1)
	node12 := listNode_new(2)
	node13 := listNode_new(1)

	node11.next = node12
	node12.next = node13

	return node11
}

func GetTestCase4() *listNode {
	node11 := listNode_new(1)
	node12 := listNode_new(2)
	node13 := listNode_new(2)
	node14 := listNode_new(1)

	node11.next = node12
	node12.next = node13
	node13.next = node14

	return node11
}
