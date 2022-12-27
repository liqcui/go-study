package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i1 = 101
	fmt.Printf("%d\n", i1)
	fmt.Printf("%b\n", i1) //把十进制转换成二进制
	fmt.Printf("%o\n", i1) //把十进制转换成八进制
	fmt.Printf("%x\n", i1) //把十进制转换成十六进制
	//八进制
	i2 := 077
	fmt.Printf("%d\n", i2)
	//十六进制
	i3 := 0x1234567
	fmt.Printf("%d\n", i3)
	//查看变量的类型
	fmt.Printf("%T\n", i3)
	f1 := 1.23456
	fmt.Printf("%T\n", f1)
	f2 := float32(1.23456)
	fmt.Printf("%T\n", f2)
	//数值比较
	v1 := 4
	v2 := 4
	fmt.Println(reflect.DeepEqual(v1, v2))

	v3 := 4
	v4 := 5
	fmt.Println(reflect.DeepEqual(v3, v4))
	a := 1
	a = a - 1
	fmt.Println(a)

	fmt.Println("int to hex")
	var i int
	i = 2
	for i <= 20 {
		i = i + 2
		h := fmt.Sprintf("%x", i)
		fmt.Println(h)
	}

	var b1 bool
	fmt.Println(b1)

}
