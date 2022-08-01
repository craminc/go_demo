package ifelse

func Judge(x, y int) {
	if x -= 2; x > y {
		println("x is bigger")
	} else if x < y {
		println("x is smaller")
	} else {
		println("x equals y")
	}
}
