package main

import "fmt"

type person struct {
	name string
	age  int16
}

func (p *person) tell() {
	fmt.Printf("my name is %v, i'm %v years old\n", p.name, p.age)
}

func multiparm(args ...string) {
	for i := 0; i < len(args); i++ {
		fmt.Printf("%v\n", args[i])
	}
}

func main() {
	var p *person = &person{name: "clark", age: 45}
	p.tell()
	multiparm("a", "b", "c", "d")
}
