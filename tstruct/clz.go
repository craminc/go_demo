package tstruct

import "fmt"

func init() {
	fmt.Println("clz started")
}

type student struct {
	name   string
	age    int
	habits []string
}

func (stu *student) SetHabits(habits []string) {
	stu.habits = habits
}

func Test() {
	s1 := student{name: "cramin", age: 22}
	habits := []string{"basketball", "swimming", "lol"}
	s1.SetHabits(habits)
	fmt.Println(s1)

	habits[1] = "football"
	fmt.Println(s1)
}

func Clz() {
	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	for _, stu := range stus {
		m[stu.name] = &stu
		fmt.Println(stu.name)
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}

func NewStu(name string, age int) *student {
	return &student{
		name: name,
		age:  age,
	}
}

func (stu student) GetName() (name string) {
	return stu.name
}

func (stu *student) SetName(name string) {
	stu.name = name
}

type Speaker string

func (s Speaker) Say() {
	fmt.Printf("I am %s, I want to say %s", s, "hello")
}
