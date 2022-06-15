package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"scaler-may-22/common"
	"scaler-may-22/week5/day13_homework/hollow_dimond_star"
	"strconv"
)

func main() {
	// YOUR CODE GOES HERE
	// Please take input and print output to standard input/output (stdin/stdout)
	// E.g. 'fmt.Scanf' for input & 'fmt.Printf' for output
	in := bufio.NewReader(os.Stdin)
	t := common.ReadInt(in)

	hollow_dimond_star.PrintHollowDimondStar(t)

	x := solve([]int{2, 1, 3})
	fmt.Println("area", x)

	x = solve([]int{10})
	fmt.Println("area", x)

	x = solve([]int{2, 1, 3, 2, 4})
	fmt.Println("area", x)
}

func solve(A []int) string {
	var ans float64 = 0
	n := len(A)
	if n > 0 {
		ans += float64(A[0]) * 0.5
		if n == 1 {
			ans += float64(A[0]) * 0.5
		} else {

			for i := 1; i < n; i++ {
				fmt.Printf("ans %v \n", ans)
				ans += (max(A[i], A[i-1]) - (math.Abs(float64((A[i] - A[i-1])) * 0.5)))
			}
			fmt.Printf("ans %v \n", ans)
			ans += float64(A[n-1]) * 0.5
		}
	}
	return strconv.Itoa(int(ans))
}

func max(a, b int) float64 {
	if a > b {
		return float64(a)
	}
	return float64(b)
}
