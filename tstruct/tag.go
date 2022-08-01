package tstruct

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	ID     int    `json:"id"`
	Gender string `json:"gender"`
	name   string
}

func Tag() {
	p1 := Person{
		ID:     1,
		Gender: "male",
		name:   "cramin",
	}
	data, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("json marshal failed!")
	} else {
		fmt.Printf("json str: %s\n", data)
	}
}
