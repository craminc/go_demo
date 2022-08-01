package tinterface

import "fmt"

type Singer interface {
	sing()
}

type Bird struct {
}

func (bird Bird) sing() {
	fmt.Println("叽叽喳喳")
}

type Dog struct {
}

func (dog Dog) sing() {
	fmt.Println("汪汪汪")
}

type Cat struct {
}

func (cat *Cat) sing() {
	fmt.Println("喵喵喵")
}
func SingASong(singer Singer) {
	singer.sing()
}

// WashingMachine 洗衣机
type WashingMachine interface {
	wash()
	dry()
}

// 甩干器
type dryer struct{}

// 实现WashingMachine接口的dry()方法
func (d dryer) dry() {
	fmt.Println("甩一甩")
}

// 海尔洗衣机
type washer struct {
	dryer //嵌入甩干器
}

// 实现WashingMachine接口的wash()方法
func (w washer) wash() {
	fmt.Println("洗一洗")
}

func WashClothes() {
	w := washer{}
	w.wash()
	w.dryer.dry()
}
