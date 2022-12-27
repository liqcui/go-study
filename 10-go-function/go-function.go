package main

import (
	"fmt"
)

//返回值可以命名也可以不命名
//命名的返回值相当于函数声明对一个变量
//return 后面可指定返回值的名字，例如retrun ret，如果不加，也是表示retrun “ret”

func f1(x int, y int) (ret int) {
	ret = x + y
	return
}

//没有参数，但是有返回值
func f2() int {
	ret := 3
	return ret
}

//参数类型的简写
//当多个参数类型时，可以将同一类的非最后一个参数的类型省略
func f3(a, b, c string, d, e, f int, h, i bool) (string, int, bool) {
	return a + b + c, d + e + f, i && h
}

//可变长参数
func f4(s1 string, i ...int) {
	fmt.Println(s1)
	fmt.Printf("%T and %v\n", i, i)
}

func sum(x int, y int) (ret int) {
	return x + y
}

func funret(x int, y int) int {
	return x + y
}

//函数作为参数和返回值
func f5(x func() int) func(x int, y int) int {

	fmt.Printf("function is a parameter")
	return funret
}
func main() {

	a := f1(3, 5)
	fmt.Println(a)
	b := f2()
	fmt.Println(b)
	c, d, e := f3("a", "b", "c", 4, 5, 6, false, true)
	fmt.Println(c, d, e)
	f4("good", 1, 2, 3, 4, 5, 6)

	r := sum(10, 35)
	fmt.Println(r)

	f5(f2)

	//函数内部没有办法声明带名字的函数
	//匿名函数
	f1 := func(x, y int) {
		fmt.Println(x + y)
	}
	f1(3, 5)

	//如果只是调用一次的函数，还可以简写成立即执行的函数
	func(x, y int) {
		fmt.Println(x + y)
	}(100, 200)
}
