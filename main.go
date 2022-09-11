package main

import (
	"bufio"
	"fmt"
	"os"
	"scaler-may-22/common"
)

func main() {
	// YOUR CODE GOES HERE
	// Please take input and print output to standard input/output (stdin/stdout)
	// E.g. 'fmt.Scanf' for input & 'fmt.Printf' for output

	in := bufio.NewReader(os.Stdin)
	str := common.ReadString(in)
	str = reverseString(str)
	fmt.Println(str)
}

func reverseString(str string) string {
	if len(str) == 0 {
		return ""
	}
	return reverseString(str[1:]) + string(str[0])
}
