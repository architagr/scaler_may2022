package main

import (
	"bufio"
	"fmt"
	"os"
	"scaler-may-22/common"
	"scaler-may-22/week2/day8_homework/separate_even_odd"
)

func main() {
	// YOUR CODE GOES HERE
	// Please take input and print output to standard input/output (stdin/stdout)
	// E.g. 'fmt.Scanf' for input & 'fmt.Printf' for output
	in := bufio.NewReader(os.Stdin)
	t := common.ReadInt(in)
	result := make([][]int, 0)

	for i := 0; i < t; i++ {
		common.ReadInt(in)
		arr := common.ReadArrInt(in)
		even, odd := separate_even_odd.SeparateEvenOdd(arr)
		result = append(result, odd)
		result = append(result, even)

	}
	for _, val := range result {
		for _, val1 := range val {
			fmt.Printf("%d ", val1)
		}
		fmt.Printf("\n")
	}
}
