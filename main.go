package main

import (
	"fmt"
	summationgame "scaler-may-22/week1/day1_assingment/summation-game"
)

func main() {
	fmt.Println("Enter a number")

	var n int
	fmt.Scanf("%d", &n)
	fmt.Println(summationgame.SummationGame(n))
}
