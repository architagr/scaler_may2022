package kth_symbol

import (
	"log"
	"strconv"
)

func KthSymbol(A int, B int) int {
	row := getRow(A-1, B, "0")
	log.Println(row)
	x, _ := strconv.Atoi(string(row[B-1]))
	return x
}

func getRow(A, B int, str string) string {
	if A == 0 || len(str) > B {
		return str
	}
	x := []byte(str)
	n := len(str)
	for i := 0; i < n; i++ {
		if str[i] == '1' {
			x = append(x, '0')
		} else {
			x = append(x, '1')
		}
	}
	return getRow(A-1, B, string(x))
}
