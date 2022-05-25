package main

import (
	"bufio"
	"fmt"
	"os"
	searchelement "scaler-may-22/week2/day8_assingment/search_element"
	"strconv"
	"strings"
)

func main() {
	// YOUR CODE GOES HERE
	// Please take input and print output to standard input/output (stdin/stdout)
	// E.g. 'fmt.Scanf' for input & 'fmt.Printf' for output
	in := bufio.NewReader(os.Stdin)

	t := readInt(in)
	result := make([]int, t)
	for i := 0; i < t; i++ {
		arr := readArrInt(in)
		b := readInt(in)
		fmt.Println("b", b, "arr", len(arr))
		if searchelement.SearchElement(arr, b) {
			result[i] = 1
		} else {
			result[i] = 0
		}
	}
	for i := 0; i < t; i++ {
		fmt.Println(result[i])
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
	fmt.Println("line", len(line))
	line = strings.ReplaceAll(line, "\r", "")
	line = strings.ReplaceAll(line, "\n", "")
	numbs := strings.Split(line, " ")
	fmt.Println("numbs", len(numbs))

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
