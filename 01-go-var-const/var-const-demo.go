package main

import "fmt"

//注意事项：
//1. 函数外的每个语句都必须以关键字开始（var,const,func等）
//2. ：=不能在函数外使用
//3._多用于占位，表示忽略值，匿名变量
//4.同一个作用域不能声明同名变量
//声明变量方式1
var name string
var age int
var isOK bool

//变量声明方式2，推荐使用驼峰命名
var (
	stuName string
	stuAge  int
	isFail  bool
)

//定义常量
const pi = 3.1415926
const (
	statusOK = 200
	notFound = 404
	n1       = 100 //批量声明常量时，如果某一行声明后没有赋值，默认就和上一行一致。
	n2
	n3
)
const (
	a1 = iota
	a2
	a3
	a4
	a5
)

//跳过
const (
	b1 = iota
	b2 = iota
	_  = iota
	b4 = iota
	b5 = iota
)

//插队
const (
	c1 = iota
	c2 = 100
	c3 = iota
	c4
	c5
)

//iota每增加一行，值加1，同一行值相等
const (
	d1, d2 = iota + 1, iota + 2
	d3, d4 = iota + 1, iota + 2
)

//定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	name = "目标"
	age = 18
	isOK = true
	stuName = "努力"
	stuAge = 20
	isFail = false
	//Go语言声明变量必须使用，否则编译不通过
	//声明变量同时赋值
	var s1 string = "who"
	//类型推导(根据值判断该变量是什么类型)
	var s2 = "20"
	//简短变量声明
	s3 := "hahaha"
	fmt.Print(isOK)                                                               //在终端输出打印的内容
	fmt.Println(name, age)                                                        //打印完指定内容之后会在后面加上一个换行符。
	fmt.Printf("Student Name is %s,Age is %d,failed？%v", stuName, stuAge, isFail) //%s占位符，使用stuName的变量值去替换占位符。
	fmt.Println(s1, s2, s3)
	//常量
	fmt.Println(pi, statusOK, notFound)
	fmt.Println(n1, n2, n3)
	fmt.Println(a1, a2, a3, a4, a5)
	fmt.Println(b1, b2, b4, b5)
	fmt.Println(c1, c2, c3, c4, c5)
	fmt.Println(d1, d2, d3, d4)
}
