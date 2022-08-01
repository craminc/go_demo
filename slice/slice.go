package slice

import "fmt"

func Slice() {
	words := [5]string{"abc", "bcd", "cde", "def", "efg"}
	fmt.Println(words)
	w := words[1:3]
	fmt.Printf("w: %s, len(w): %d, cap(w): %d\n", w, len(w), cap(w))
	r := w[1:4]
	fmt.Printf("r: %s, len(r): %d, cap(r): %d\n", r, len(r), cap(r))
	arr := make([]string, 2, 10)
	fmt.Println(arr)
	s := make([]int, 0, 0)
	s = append(s, 1, 2, 3, 4)
	fmt.Printf("s: %v, ptr: %p\n", s, s)
	fmt.Println(cap(s))
	s = append(s[:2], s[3:]...)
	fmt.Printf("s: %v, ptr: %p\n", s, s)
	fmt.Println(cap(s))

	a := make([]string, 0, 10)
	for i := 0; i < 10; i++ {
		fmt.Println(a)
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(a)
}
