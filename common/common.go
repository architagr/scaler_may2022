package common

import (
	"bufio"
	"strconv"
	"strings"
)

func ReadInt(in *bufio.Reader) int {
	nStr, _ := in.ReadString('\n')
	nStr = strings.ReplaceAll(nStr, "\r", "")
	nStr = strings.ReplaceAll(nStr, "\n", "")
	nStr = strings.TrimSpace(nStr)
	nStr = strings.Trim(nStr, "\t \n")
	n, _ := strconv.Atoi(nStr)
	return n
}

func ReadLineNumbs(in *bufio.Reader) []string {
	line, _ := in.ReadString('\n')
	line = strings.ReplaceAll(line, "\r", "")
	line = strings.ReplaceAll(line, "\n", "")
	line = strings.TrimSpace(line)
	line = strings.Trim(line, "\t \n")
	numbs := strings.Split(line, " ")
	return numbs
}

func ReadArrInt(in *bufio.Reader) []int {
	numbs := ReadLineNumbs(in)
	arr := make([]int, 0)
	for _, n := range numbs[1:] {
		val, _ := strconv.Atoi(n)
		arr = append(arr, val)
	}
	return arr
}
