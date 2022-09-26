package linked_list

import "fmt"

type Node struct {
	Data int
	Next *Node
}

var head *Node

func Insert_node(position, value int) {
	// @params position, integer
	// @params value, integer
	newNode := new(Node)
	newNode.Data = value
	if head == nil {
		head = newNode
		return
	}
	if position == 0 {
		newNode.Next = head
		head = newNode
		return
	}
	c := 0
	x := head
	for x.Next != nil && c < position-1 {
		x = x.Next
		c++
	}
	newNode.Next = x.Next
	x.Next = newNode
}

func Delete_node(position int) {
	// @params position, integer
	if head == nil {
		return
	}
	if position == 0 {
		head = head.Next
		return
	}
	c := 0
	x := head
	for x.Next != nil && c < position-1 {
		x = x.Next
		c++
	}

	if c <= position && x.Next != nil {
		x.Next = x.Next.Next
	}

}

func Print_ll() {
	// Output each element followed by a space
	x := head
	s := ""
	if x != nil {
		s = fmt.Sprintf("%d", x.Data)
		x = x.Next
	}
	for x != nil {
		s = fmt.Sprintf("%s %d", s, x.Data)
		x = x.Next
	}
	fmt.Printf("\n%s", s)

}
