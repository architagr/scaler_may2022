package main

import (
	"bufio"
	"fmt"
	"os"
	"scaler-may-22/common"
	rotationgame "scaler-may-22/week2/day8_assingment/rotation_game"
)

func main() {
	// YOUR CODE GOES HERE
	// Please take input and print output to standard input/output (stdin/stdout)
	// E.g. 'fmt.Scanf' for input & 'fmt.Printf' for output
	in := bufio.NewReader(os.Stdin)
	arr := common.ReadArrInt(in)
	k := common.ReadInt(in)

	result := rotationgame.RotateArray(arr, k)
	for _, val := range result {
		fmt.Printf("%d ", val)
	}
}
