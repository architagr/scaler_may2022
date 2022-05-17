package main

import (
	"fmt"
	multiplicationtable "scaler-may-22/week1/day1_homework/multiplication-table"
)

func main() {
	fmt.Println("Enter a number")

	var n int
	fmt.Scanf("%d", &n)
	multiplicationtable.PrintTable(n)
}
