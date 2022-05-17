package main

import (
	"fmt"
	armstrongnumbers "scaler-may-22/week1/day1_homework/armstrong-numbers"
)

func main() {
	fmt.Println("Enter a number")

	var n int
	fmt.Scanf("%d", &n)
	result := armstrongnumbers.GetArmstrongInRange(n)
	for _, val := range result {
		fmt.Println(val)
	}
}
