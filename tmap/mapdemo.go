package tmap

import "fmt"

func TMap() {
	//m := make(map[string]int, 8)
	m := map[string]int{
		"c": 3,
		"d": 4,
	}
	m["a"] = 1
	m["b"] = 2
	fmt.Println(m, m["a"], m["b"], m["c"])
	fmt.Printf("type of m: %T \nlen of m: %d\n", m, len(m))
	val, exist := m["d"]
	fmt.Printf("exist: %t, value: %d\n", exist, val)
	delete(m, "a")
	for k, v := range m {
		fmt.Printf("key: %s, value: %d\n", k, v)
	}
}

func Test() {
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s)
	m["q1mi"] = s
	s = append(s[:1], s[2:]...)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", m["q1mi"])
}
