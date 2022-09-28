package reverse_linked_list

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	head     *listNode
	expected []int
}

func TestReverseLinkedList(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		head:     GetTestCase1(),
		expected: []int{35, 76, 46},
	})
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := listToArr(ReverseLinkedList(testCase.head))

			if !reflect.DeepEqual(x, testCase.expected) {
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

func listToArr(head *listNode) []int {
	arr := make([]int, 0)
	if head == nil {
		return arr
	}
	for head != nil {
		arr = append(arr, head.value)
		head = head.next
	}
	return arr
}
