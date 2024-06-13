package main

import (
	"fmt"
)

func main() {

	//var age int
	age := 20
	if age >= 18 {
		fmt.Println("已经成年")
	} else if age >= 80 {
		fmt.Println("已经老年")
	} else {
		fmt.Println("未成年")
	}

	//For循环
	s := "Hello长春"
	for i, v := range s {
		fmt.Printf("%d,%c\n", i, v)
	}
	//死循环
	// for{
	// 	fmt.Println("for ever loop")
	// }
	for a := 1; a < 10; a++ {
		for j := 1; j <= a; j++ {
			fmt.Printf("%d*%d=%d\t", a, j, a*j)
		}
		fmt.Println()
	}
	for a := 1; a < 10; a++ {
		for j := a; j < 10; j++ {
			fmt.Printf("%d*%d=%d\t", a, j, a*j)
		}
		fmt.Println()
	}

	//流程控制，跳出for循环
	for i := 0; i <= 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("over")
	//流程控制，跳出本地循环，继续下一次循环
	for i := 0; i <= 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("over")

	//case语句
	n := 5
	switch n {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效数据")
	}

	ocpkind := "deployment"
	switch ocpkind {
	case "deployment":
		fmt.Println("deployment ")
	case "daemonset":
		fmt.Println("daemonset ")
	default:
		fmt.Println("daemonset ")
	}

	switch n := 3; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8, 10:
		fmt.Println("偶数")
	default:
		fmt.Println("超出范围")
	}

	s1 := "1"
	s2 := "2"
	if s1 == s2 {
		fmt.Println(s1 == s2)
	} else {
		fmt.Println(s1 == s2)
	}

	for i := 4; i > 0; i-- {
		fmt.Println(i)
	}

	workerNodes := []string{"a", "e", "f"}
	masterNodes := []string{"a"}
	matchCount := 0
	for i := 0; i < len(workerNodes); i++ {
		for j := 0; j < len(masterNodes); j++ {
			if workerNodes[i] == masterNodes[j] {
				matchCount++
			}
		}
	}
	fmt.Println(matchCount)
	i := 0
	for { //相当于while true
		if i < 10 {
			fmt.Println(i)
		} else {
			break
		}
		i++
	}
}

func is3Master() bool {
	if true {
		return true
	}
	return false

}
