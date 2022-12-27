package main

import (
	"fmt"
)

//自定义数据类型
type myInt int

//结构体
type s1 struct {
	name      string
	namespace string
}

//结构体是值类型
type person struct {
	name, gender string
}

//go语言中函数参数永远是拷贝
func f(x person) {
	x.name = "clark"
	x.gender = "女" //修改的是副本的gender
	fmt.Println(x.name)
}

func f2(x *person) {
	//(*x).gender = "女" 根据内存地址找到那个原变量，修改的是原来的变量
	(*x).gender = "女" //自动根据指针找到对应的变量
}

//结构体匿名字段
type people struct {
	string
	int
}

//嵌套结构体
type personInfo struct {
	name     string
	age      int
	province string
	city     string
	compnay  compnay
}

type compnay struct {
	companyName string
}

//结构体继承
type animal struct {
	name string
}

func (a animal) move() {
	fmt.Printf("%s会动！", a.name)
}

type dog struct {
	feet uint8
	animal
}

func (d dog) wang() {
	fmt.Printf("%s is shutting：", d.name)
}

type location struct {
	x     int
	y     int
	valid bool
}

func main() {

	var b myInt
	b = 100
	fmt.Println(b)

	//结构体1
	a := s1{
		name: "test",
	}
	fmt.Printf("%T,%v", a.namespace, a.namespace)
	fmt.Println(len(a.namespace))

	//匿名结构体,临时场景
	var s struct {
		x string
		y int
	}
	s.x = "good"
	s.y = 100
	fmt.Printf("type:%T value:%v", s, s)

	var p person
	p.name = "Mike"
	p.gender = "Male"
	f(p)
	fmt.Println(p.gender) //Male

	f2(&p) //0x234ac5
	fmt.Println(p.gender)

	var p2 = new(person)
	p2.name = "good"
	fmt.Printf("%T\n%p\n", p2, p2)
	// wait.Poll(time.Second, time.Second*5, func() (done bool, err error) {
	// 	fmt.Println(time.Now())
	// 	return false, nil
	// })

	p3 := people{
		"zhouzongli",
		9000,
	}
	fmt.Println(p3)
	fmt.Println(p3.string)

	//嵌套结构体
	p4 := personInfo{
		name:     "Good",
		age:      28,
		province: "JL",
		city:     "CC",
		compnay: compnay{
			companyName: "RH",
		},
	}
	fmt.Println(p4)
	fmt.Println(p4.compnay.companyName)

	//结构体继承
	d1 := dog{
		animal: animal{name: "zhoulin"},
		feet:   4,
	}

	fmt.Println(d1)
	d1.move()
	d1.wang()

	//初始化结构体
	loc := new(location)
	loc.x = 10
	fmt.Println(loc.x)
	loc2 := location{
		x: 11,
	}
	fmt.Println(loc2.x)

	//slice of structs
	//create struct and append it to the slice.
	places := []*location{}
	loc = new(location)
	loc.x = 10
	loc.y = 20
	loc.valid = true

	places = append(places, loc)

	loc = new(location)
	loc.x = 15
	loc.y = 25
	loc.valid = true

	places = append(places, loc)

	for i := range places {
		place := places[i]
		fmt.Println("Location:", place)
	}
}

//应用其他package的结构体，要研究一下
