package middle_element_of_linked_list

import (
	"fmt"
	"testing"
)

type TestCase struct {
	head     *listNode
	expected int
}

func TestGetMiddleEle(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		head:     GetTestCase1(),
		expected: 76,
	})
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := GetMiddleEle(testCase.head)
			if x != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}

func GetTestCase1() *listNode {
	node1 := listNode_new(46)
	node2 := listNode_new(76)
	node3 := listNode_new(35)
	node1.next = node2
	node2.next = node3
	return node1
}
