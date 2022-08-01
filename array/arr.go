package array

import "fmt"

func Array() {
	var arr [3]int
	arr = [...]int{1, 2, 3}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
	fmt.Println(arr)
}

func Equal(x [3]int, y [3]int) bool {
	return x == y
}
