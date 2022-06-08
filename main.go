package main

import (
	"bufio"
	"os"
	"scaler-may-22/common"
	"scaler-may-22/week5/day13_homework/hollow_dimond_star"
)

func main() {
	// YOUR CODE GOES HERE
	// Please take input and print output to standard input/output (stdin/stdout)
	// E.g. 'fmt.Scanf' for input & 'fmt.Printf' for output
	in := bufio.NewReader(os.Stdin)
	t := common.ReadInt(in)

	hollow_dimond_star.PrintHollowDimondStar(t)
}
