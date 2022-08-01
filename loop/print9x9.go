package loop

import "fmt"

func MultiTable(x int) {
	for i := 1; i <= x; i++ {
		for j := i; j <= x; j++ {
			fmt.Printf("%d x %d = %2d | ", i, j, i*j)
		}
		fmt.Println()
	}
}
