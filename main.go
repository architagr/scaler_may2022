package main

import (
	"fmt"
	isitperfect "scaler-may-22/week1/day1_assingment/is-it-perfect"
	_ "scaler-may-22/week1/day1_assingment/isitprime"
)

func main() {
	fmt.Println("Enter a number")

	var n int8
	fmt.Scanf("%d", &n)
	io := make([]bool, 0, n)
	var i int8 = 0
	for ; i < n; i++ {
		var a int64
		fmt.Scanf("%d", &a)
		check := isitperfect.IsItPerfect(a)
		io = append(io, check)
	}
	for _, check := range io {
		if check {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}

}
