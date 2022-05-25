package main

import (
	"bufio"
	"fmt"
	"os"
	rotationgame "scaler-may-22/week2/day8_assingment/rotation_game"
	"strconv"
	"strings"
)

func main() {
	// YOUR CODE GOES HERE
	// Please take input and print output to standard input/output (stdin/stdout)
	// E.g. 'fmt.Scanf' for input & 'fmt.Printf' for output
	in := bufio.NewReader(os.Stdin)
	arr := readArrInt(in)
	k := readInt(in)

	result := rotationgame.RotateArray(arr, k)
	for _, val := range result {
		fmt.Printf("%d ", val)
	}
}

func readInt(in *bufio.Reader) int {
	nStr, _ := in.ReadString('\n')
	nStr = strings.ReplaceAll(nStr, "\r", "")
	nStr = strings.ReplaceAll(nStr, "\n", "")
	n, _ := strconv.Atoi(nStr)
	return n
}

func readLineNumbs(in *bufio.Reader) []string {
	line, _ := in.ReadString('\n')
	line = strings.ReplaceAll(line, "\r", "")
	line = strings.ReplaceAll(line, "\n", "")
	line = strings.TrimSpace(line)
	line = strings.Trim(line, "\t \n")
	numbs := strings.Split(line, " ")
	return numbs
}

func readArrInt(in *bufio.Reader) []int {
	numbs := readLineNumbs(in)
	arr := make([]int, 0)
	for _, n := range numbs[1:] {
		val, _ := strconv.Atoi(n)
		arr = append(arr, val)
	}
	return arr
}
