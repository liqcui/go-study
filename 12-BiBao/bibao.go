package main

import (
	"fmt"
	"strings"
)

func invokedfunc(x, y int) {
	fmt.Println("this is invoked function")
	fmt.Println(x + y)
}

func encapfunc(x func(int, int), m, n int) func() {
	tmp := func() {
		x(m, n)
	}
	//tmp() //x(m,n)
	return tmp
}

//闭包是什么
//闭包是个函数，这个函数包含了他外部作用域的一个变量
//底层原理
//1. 函数可以作为返回值
//2. 函数内部查找变了的顺序，先在自己内部找，找不到往外层找
//闭包等于函数+外部变量的引用
//例如 外部变量x+ func(y int) int {
//	        	x += y
//		        return x
//              }
func calAdd(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	ret := encapfunc(invokedfunc, 100, 320)
	ret()

	retAdd := calAdd(100)
	ret2 := retAdd(200)
	fmt.Println(ret2)

	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test"))
	fmt.Println(jpgFunc("aaa.jpg"))
	fmt.Println(txtFunc("back"))
	fmt.Println(txtFunc("bbb.txt"))

	f1, f2 := calc(10)
	fmt.Println(f1(1), f2(2))
	fmt.Println(f1(3), f2(4))
}
