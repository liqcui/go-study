package main

import (
	"fmt"
	"strings"
	"testing"
)

func main() {

	//数组
	//存放元素的容量
	//必须指定存放元素的类型和容量（长度）
	//数组的长度是数组类型的一部分
	var a1 [3]bool //[true,false, true]
	var a2 [5]int
	fmt.Printf("a1:%T\na2:%T", a1, a2)

	//数组的初始化
	//如果不初始化：默认的元素都是零值（布尔值：false，整型和浮点类型都是0，字符串：""）
	fmt.Printf("a1:%T a2:%T", a1, a2)

	//数组的初始化
	//1.初始化方式1
	a1 = [3]bool{true, true, false}
	fmt.Println(a1)
	//2.初始化方式2: 根据初始值自动判断数组的长度是多少
	a10 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	//3.初始化方式3:根据索引来初始化
	a3 := [5]int{0: 1, 4: 2}
	fmt.Println(a10, a3)

	//数组的遍历
	citys := [...]string{"beijing", "changchun", "shenyang"}
	//索引：0～2
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}

	for index, value := range citys {
		fmt.Println(index, value)
	}

	//多维数组
	//[[1 2]][3,4][5,6]
	var a11 [3][2]int
	a11 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(a11)

	//二维数组,多维数组只有最外层可以使用...
	var a12 = [...][2]int{
		[2]int{1, 3},
		[2]int{2, 5},
		[2]int{3, 7},
	}
	fmt.Println(a12)

	for _, v1 := range a11 {
		fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

	//数组是值类型
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[0] = 100
	fmt.Println(b1, b2)

	//1.求数组[1,4,5,7,13]所有元素的和
	a5 := [...]int{1, 4, 5, 7, 13}
	sum := 0
	for _, v := range a5 {
		sum = sum + v
	}
	fmt.Println(sum)

	//找出数组中和为指定值的元素下标，比如数组[1,3,5,7,8]
	//中和为8的两个元素的下标分别为（0，3）（1，2）
	b3 := [...]int{1, 3, 5, 7, 8}
	for i := 0; i < len(b3); i++ {
		for j := i + 1; j < len(b3); j++ {
			if b3[i]+b3[j] == 8 {
				fmt.Printf("(%d,%d)\n", i, j)
			}
		}
	}

	//数组是值类型
	x := [3]int{1, 2, 3}
	y := x
	y[1] = 200
	fmt.Println(x)
	f1(x)
	fmt.Println(x)
	abc := lenarray()
	fmt.Println(abc)

	a6 := []int{0, 1, 2, 3, 4}
	//删除第i个元素
	i6 := 2
	a6 = append(a6[:i6], a6[i6+1:]...)
	fmt.Println(a6)
	split_arry()

	paoDeployOCPVersionList := []string{"4.6", "4.7", "4.8", "4.9", "4.10"}
	for _, v := range paoDeployOCPVersionList {
		fmt.Println(v)
	}

}

func f1(a [3]int) {
	//Go语言中的函数传递的都是值（ctrl+c，ctrl+v）
	a[1] = 100
}

func lenarray() int {
	var strarry []string
	//strarry := []string{"good", "bad", "fine", "well"}
	strarry = []string{}
	fmt.Println(strarry)
	fmt.Println(len(strarry))
	return len(strarry)
}

func split_arry() {
	//string分割空字符串，生成的字符串数组长度自动加一
	empty_string := ""
	result := strings.Split(empty_string, " ")
	fmt.Printf("The length of empty string after split is:%v", len(result)) // here length is 1

}

func square(arr *[3]int) {
	for i, num := range *arr {
		(*arr)[i] = num * num
	}
}

func TestArrayPointer(t *testing.T) {
	a := [...]int{1, 2, 3}
	square(&a)
	fmt.Println(a) // [1 4 9]
	if a[1] != 4 && a[2] != 9 {
		t.Fatal("failed")
	}
}
