package main

import "fmt"

//接口是一种类型，是一种特殊的类型
//我不关心一个变量是什么类型，我只关心能调用它的什么方法
//定义一个能叫的类型
type speaker interface {
	speak() //只要实现了speak方法的变量就是speaker类型。
}

type cat struct{}

type dog struct{}

type person struct{}

func (c cat) speak() {
	fmt.Println("miao miao")
}

func (d dog) speak() {
	fmt.Println("wang wang")
}

func (p person) speak() {
	fmt.Println("haha")
}

func beat(x speaker) {
	//接收一个参数，传递来什么，我就打什么
	x.speak() //挨打了就叫
}

// 不管什么牌子的车，都能跑
//定义一个car接口类型
//不管是什么结构体，只要run方法都能上car类型
type car interface {
	run()
}

type falali struct {
	brand string
}

type baoshijie struct {
	brand string
}

func (f falali) run() {
	fmt.Printf("%s速度是70迈～\n", f.brand)
}

func (b baoshijie) run() {
	fmt.Printf("%s速度是700迈～\n", b.brand)
}

//drive函数接收一个car类型的变量
func driver(c car) {
	c.run()
}

type animal interface {
	move()
	eat(food string)
}

type goat struct {
	name string
	feet int8
}

func (c goat) move() {
	fmt.Println("走羊步...")
}

func (c goat) eat(food string) {
	fmt.Printf("羊吃%v...\n", food)
}

type chicken struct {
	feet int8
}

func (c chicken) move() {
	fmt.Println("鸡冻")
}

func (c chicken) eat() {
	fmt.Println("吃鸡饲料")
}

func main() {
	var c1 cat
	var d1 dog
	var p1 person

	beat(c1)
	beat(d1)
	beat(p1)

	var f1 = falali{
		brand: "法拉利",
	}

	var b1 = baoshijie{
		brand: "保时捷",
	}

	driver(f1)
	driver(b1)

	var ss speaker
	ss = c1
	ss = d1
	ss = p1
	fmt.Println(ss)

	var a1 animal
	bc := goat{
		name: "mimi",
		feet: 4,
	}

	a1 = bc
	a1.eat("青草")
	fmt.Println(a1)

	//空接口
	//interface 关键字
	//interface{}:空接口类型
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 16)
	m1["name"] = "zhou"
	m1["age"] = 18
	m1["married"] = true
	m1["hobby"] = []string{"sing", "jump", "rap"}
	fmt.Println(m1)
}
