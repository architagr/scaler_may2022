package linked_list

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	position, value int
	head            *Node
	expected        []int
}

func TestInsertNode(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		position: 0,
		value:    1,
		head:     nil,
		expected: []int{1},
	})
	testCases = append(testCases, TestCase{
		position: 2,
		value:    3,
		head:     GetTestCase1(),
		expected: []int{1, 2, 3, 4},
	})
	testCases = append(testCases, TestCase{
		position: 0,
		value:    3,
		head:     GetTestCase1(),
		expected: []int{3, 1, 2, 4},
	})
	testCases = append(testCases, TestCase{
		position: 10,
		value:    5,
		head:     GetTestCase1(),
		expected: []int{1, 2, 4, 5},
	})
	testCases = append(testCases, TestCase{
		position: 3,
		value:    5,
		head:     GetTestCase1(),
		expected: []int{1, 2, 4, 5},
	})
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			head = testCase.head
			Insert_node(testCase.position, testCase.value)
			temp := head
			x := GetLitListToArray(temp)

			if !reflect.DeepEqual(x, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}

func TestDeletetNode(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		position: 0,
		head:     nil,
		expected: []int{},
	})
	testCases = append(testCases, TestCase{
		position: 2,
		head:     GetTestCase1(),
		expected: []int{1, 2},
	})
	testCases = append(testCases, TestCase{
		position: 0,
		head:     GetTestCase1(),
		expected: []int{2, 4},
	})
	testCases = append(testCases, TestCase{
		position: 10,
		head:     GetTestCase1(),
		expected: []int{1, 2, 4},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			head = testCase.head
			Delete_node(testCase.position)
			temp := head
			x := GetLitListToArray(temp)

			if !reflect.DeepEqual(x, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}

func TestPrintll(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		head: nil,
	})
	testCases = append(testCases, TestCase{
		head: GetTestCase1(),
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			head = testCase.head
			Print_ll()

		})
	}
}

func GetTestCase1() *Node {
	var node3 *Node = new(Node)
	node3.Data = 4

	var node2 *Node = new(Node)
	node2.Data = 2
	node2.Next = node3

	var node1 *Node = new(Node)
	node1.Data = 1
	node1.Next = node2

	return node1
}

func GetLitListToArray(head *Node) []int {
	arr := make([]int, 0)

	for head != nil {
		arr = append(arr, head.Data)
		head = head.Next
	}
	return arr
}
