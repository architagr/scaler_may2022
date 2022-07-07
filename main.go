package main

import (
	"bufio"
	"os"
	"scaler-may-22/common"
	"scaler-may-22/week9/day25_assingment/linked_list"
	"strconv"
)

func main() {
	// YOUR CODE GOES HERE
	// Please take input and print output to standard input/output (stdin/stdout)
	// E.g. 'fmt.Scanf' for input & 'fmt.Printf' for output

	in := bufio.NewReader(os.Stdin)
	t := common.ReadInt(in)
	for i := 0; i < t; i++ {
		line := common.ReadLineNumbs(in)
		switch line[0] {
		case "i":
			pos, _ := strconv.Atoi(line[1])
			value, _ := strconv.Atoi(line[2])
			linked_list.Insert_node(pos, value)
		case "d":
			pos, _ := strconv.Atoi(line[1])
			linked_list.Delete_node(pos - 1)
		case "p":
			linked_list.Print_ll()
			common.ReadInt(in)
		}

	}
}
