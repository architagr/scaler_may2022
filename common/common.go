package common

import (
	"bufio"
	"strconv"
	"strings"
)

func ReadString(in *bufio.Reader) string {
	nStr, _ := in.ReadString('\n')
	nStr = strings.ReplaceAll(nStr, "\r", "")
	nStr = strings.ReplaceAll(nStr, "\n", "")
	nStr = strings.TrimSpace(nStr)
	nStr = strings.Trim(nStr, "\t \n")
	return nStr
}
func ReadInt(in *bufio.Reader) int {
	nStr := ReadString(in)
	n, _ := strconv.Atoi(nStr)
	return n
}

func ReadStrArr(in *bufio.Reader) []string {
	line := ReadString(in)
	numbs := strings.Split(line, " ")
	return numbs
}

func ReadArrInt(in *bufio.Reader) []int {
	numbs := ReadStrArr(in)
	arr := make([]int, 0)
	for _, n := range numbs {
		val, _ := strconv.Atoi(n)
		arr = append(arr, val)
	}
	return arr
}

func MaxInt() int {
	const MaxUint = ^uint(0)
	const MaxInt = int(MaxUint >> 1)
	return MaxInt
}
func MinInt() int {
	const MaxUint = ^uint(0)
	const MaxInt = int(MaxUint >> 1)
	const MinInt = -MaxInt - 1
	return MinInt
}
