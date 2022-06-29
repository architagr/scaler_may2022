package main

import (
	"bufio"
	"fmt"
	"os"
	"scaler-may-22/common"
	"scaler-may-22/week8/day23_assingment/print_reverse_string"
)

func main() {
	// YOUR CODE GOES HERE
	// Please take input and print output to standard input/output (stdin/stdout)
	// E.g. 'fmt.Scanf' for input & 'fmt.Printf' for output
	in := bufio.NewReader(os.Stdin)
	str := common.ReadString(in)
	str1 := print_reverse_string.ReverseString(str, len(str)-1)
	fmt.Println(str1)
}
