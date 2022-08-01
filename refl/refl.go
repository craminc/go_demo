package refl

import (
	"fmt"
	"math"
	"reflect"
)

func reflectType(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("type: %v\n", t)
	fmt.Printf("type.name: %v\n", t.Name())
	fmt.Printf("type.kind: %v\n", t.Kind())
	fmt.Println()
}

func ReflectType() {
	var a *float32
	reflectType(a)

	var b rune = 100
	reflectType(b)

	reflectType("cramin")

	type person struct {
		name string
		age  int
	}
	ming := person{
		"ming", 22,
	}
	reflectType(ming)

	reflectType([3]int{1, 2, 3})
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Printf("value: %v\n", v)
	k := v.Kind()
	switch k {
	case reflect.Float64:
		fmt.Printf("parse: %+v\n", v.Float())
	case reflect.Ptr:
		fmt.Printf("parse: %+v\n", v.Pointer())
	case reflect.Int:
		fmt.Printf("parse: %+v\n", v.Int())
	}
	fmt.Printf("value.kind: %v\n", k)

}

func reflectSetValue(x interface{}) {
	fmt.Println(x)
	v := reflect.ValueOf(x)
	v.Elem().SetFloat(math.Pi)
	fmt.Println(x)
}

func ReflectValue() {
	a := 3.14
	reflectValue(a)
	b := &a
	reflectValue(b)
	fmt.Println(a)
	reflectSetValue(b)
	fmt.Println(a)
	c := &b
	fmt.Println(**c)
}

type student struct {
	Name  string `json:"name" short:"n"`
	Score int    `json:"score" short:"s"`
}

func (s student) Study(x string) (msg string) {
	msg = x + "好好学习，天天向上"
	fmt.Println(msg)
	return
}

func (s student) Sleep(x string) (msg string) {
	msg = x + "吃饱饱，睡觉觉"
	fmt.Println(msg)
	return
}

func ReflectStructFiled() {
	stu := student{
		Name:  "cramin",
		Score: 100,
	}

	t := reflect.TypeOf(stu)
	fmt.Printf("name: %v, kind: %v\n", t.Name(), t.Kind())

	for i := 0; i < t.NumField(); i++ {
		filed := t.Field(i)
		fmt.Printf("name: %v, index: %d, type: %v, tag: %v\n", filed.Name, filed.Index, filed.Type, filed.Tag.Get("short"))
	}

	if f, ok := t.FieldByName("Name"); ok {
		fmt.Printf("name: %v, index: %d, type: %v, tag: %v\n", f.Name, f.Index, f.Type, f.Tag.Get("json"))
	}
}

func ReflectStructMethod() {
	stu := student{
		Name:  "cramin",
		Score: 100,
	}

	v := reflect.ValueOf(stu)
	fmt.Printf("value: %v, kind: %v\n", v, v.Kind())

	for i := 0; i < v.NumMethod(); i++ {
		method := v.Method(i)
		method.Call([]reflect.Value{reflect.ValueOf("cramin ")})
	}

	v.MethodByName("Sleep").Call([]reflect.Value{reflect.ValueOf("chengruimin ")})
}
