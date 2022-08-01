package function

import (
	"errors"
	"fmt"
)

func Calc(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

type cal func(int, int) int

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func Math() {
	var c cal
	c = add
	fmt.Printf("type of c: %T\n", c)
	fmt.Println("add result: ", c(2, 3))
	c = sub
	fmt.Printf("type of c: %T\n", c)
	fmt.Println("sub result: ", c(4, 3))

	fmt.Println("function param: ", funcParam(10, 5, sub))
}

func funcParam(x, y int, fu cal) int {
	return fu(x, y)
}

func FuncReturn(s string) (func(int, int) int, error) {
	switch s {
	case "add":
		return add, nil
	case "sub":
		return sub, nil
	default:
		err := errors.New("unknown function")
		return nil, err
	}
}

func AnoFunc() func(int, int) int {
	return func(x, y int) int {
		return x + y
	}
}

func Closure(a string) func(string) string {
	return func(b string) string {
		return b + a
	}
}

func DeferExec() {
	fmt.Println("start")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end1")
	fmt.Println("end2")
}
