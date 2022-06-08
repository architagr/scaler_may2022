package hollow_dimond_star

import "fmt"

func PrintHollowDimondStar(N int) {
	printUpperPart(N)
	printLowerPart(N)
}
func printLowerPart(n int) {
	for i := n - 1; i >= 0; i-- {
		firstHalfStarCount := n - i
		spaces := 2 * i

		for j := 0; j < firstHalfStarCount; j++ {
			fmt.Printf("*")
		}
		for j := 0; j < spaces; j++ {
			fmt.Printf(" ")
		}
		for j := 0; j < firstHalfStarCount; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

func printUpperPart(n int) {
	for i := 0; i < n; i++ {
		firstHalfStarCount := n - i
		spaces := 2 * i

		for j := 0; j < firstHalfStarCount; j++ {
			fmt.Printf("*")
		}
		for j := 0; j < spaces; j++ {
			fmt.Printf(" ")
		}
		for j := 0; j < firstHalfStarCount; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}
