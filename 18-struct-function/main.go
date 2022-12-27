package main

import (
	"fmt"
)

//构造函数

type person struct {
	name string
	age  int
}

//构造函数
//返回的结构体还是结构体指针
//当结构体比较大的时候尽量使用结构体指针，减少程序的内存开销
func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

//结构体2
type dog struct {
	name string
}

//构造函数
func newDog(name string) dog {
	return dog{
		name: name,
	}
}

//方法是作用于特定类型的函数
//接受者表示的是调用钙方法的具体类型变量，多用类型名首字母表示。
func (d dog) wang() {
	fmt.Printf("%s:汪汪汪", d.name)
}

type measure struct {
	size int
	unit string
}

func main() {
	p1 := newPerson("JC", 19)
	fmt.Println(p1)

	d1 := newDog("zhzhang")
	d1.wang()

	//A map with struct keys.
	storage := map[measure]bool{}

	m := new(measure)
	m.size = 10
	m.unit = "centimeters"
	storage[*m] = true

	m = new(measure)
	m.size = 15
	m.unit = "feet"
	storage[*m] = true

	fmt.Println("Map len", len(storage))

	key := new(measure)
	key.size = 21
	key.unit = "decibes"
	v := storage[*key]
	fmt.Println("Result", key, v)

}
