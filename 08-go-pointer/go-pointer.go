package main

import "fmt"

func main() {
	//1. &取内存地址
	n := 18
	p := &n //int类型的指针
	fmt.Printf("type is :%T, value is: %v\n", p, p)

	//2. *根据地址取值
	m := *p
	fmt.Printf("type is :%T, value is %v\n", m, m)

	//new申请一个内存地址,很少用，一般是用来给基本数据类型申请地址，int，string等
	var a1 = new(int)
	// *a1 = 20
	fmt.Println(a1)
	fmt.Println(*a1)
	*a1 = 100
	fmt.Println(*a1)
	fmt.Println(a1)

	var a2 *int
	fmt.Println(a2)

	//make申请内存地址，只用于map和slice和channel，make函数返回的是三种数据类型本身，因为这三种数据类型是引用类型。

}
