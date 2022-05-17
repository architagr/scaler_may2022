package multiplicationtable

import "fmt"

func PrintTable(n int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d * %d = %d\n", n, i, (n * i))
	}
}
