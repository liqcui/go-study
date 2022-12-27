package main

import "fmt"

//递归，函数自己调用自己
//递归适合处理那种问题相同、问题规模越来越小的场景
//递归一定要有一个明确的退出条件
//永远不要高估自己

//3!=3*2*1      =3*2!
//4!=4*3*2*1    =4*3!
//5!=5*4*3*2*1  =5*4!

func f(n int64) int64 {
	if n <= 1 {
		return 1
	}
	return n * f(n-1)
}

//上台阶
//n个台阶，一次可以走一步，也可以走2步，有多少种走法
func steps(n uint64) uint64 {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}
	return steps(n-1) + steps(n-2)
}

func main() {
	ret := f(7)
	fmt.Println(ret)

	ret2 := steps(5)
	fmt.Println(ret2)
}
