package loop

import "fmt"

func ForLoop(x string) {
	for i, el := range x {
		fmt.Printf("index: %d, element: %c \n", i, el)
	}
}
