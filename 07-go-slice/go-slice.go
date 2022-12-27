package main

import (
	"fmt"
	"strings"
)

func main() {
	//切片的定义
	var s1 []int    //定义一个存放int类型元素的切片,没有分配内存
	var s2 []string //定义一个存放string类型元素的切片,没有分配内存
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

	//切片的初始化，,没有分配内存
	s1 = []int{1, 2, 3}
	s2 = []string{"beijing", "shanghai", "tianjin"}
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

	//切片的长度和容量
	fmt.Printf("len(s1):%d  cap[s1]:%d\n", len(s1), cap(s1))
	fmt.Printf("len(s2):%d  cap[s2]:%d\n", len(s2), cap(s2))

	//2.由数组得到切片
	a1 := [...]int{1, 3, 4, 5, 6, 8, 9, 11, 12}
	ss1 := a1[0:4] //基于一个数组的切割，左包含右不包含（左开右闭）
	fmt.Println(ss1)
	ss2 := a1[3:8]
	fmt.Println(ss2)
	ss3 := a1[:4]
	ss4 := a1[3:]
	ss5 := a1[:]
	fmt.Println(ss3, ss4, ss5)

	//切片的容量是指底层数组的容量
	fmt.Printf("len(ss3):%d  cap[ss3]:%d\n", len(ss3), cap(ss3))
	fmt.Printf("len(ss4):%d  cap[ss4]:%d\n", len(ss4), cap(ss4))
	fmt.Printf("len(ss5):%d  cap[ss5]:%d\n", len(ss5), cap(ss5))

	//切片再切片
	s8 := ss4[3:]
	fmt.Printf("len(s8):%d  cap[s8]:%d\n", len(s8), cap(s8))
	//切片时引用类型，都指向了底层的一个数组
	fmt.Println("ss4: ", ss4)
	a1[6] = 1300 //修改了底层数组的值
	fmt.Println("ss4: ", ss4)
	fmt.Println("s8: ", s8)

	//make函数创造切片，分配内存
	//切片就是一个框，框住了一块连续的内存
	ss6 := make([]int, 5, 10)
	fmt.Printf("ss6=%v len(ss6)=%d cap(ss6)=%d\n", ss6, len(ss6), cap(ss6))
	ss7 := make([]int, 0, 10)
	fmt.Printf("ss7=%v len(ss7)=%d cap(ss7)=%d\n", ss7, len(ss7), cap(ss7))
	//要判断一个切片是否为空，用len(s)==0来判断，不能用s==nil来判断

	//切片的赋值
	a3 := []int{1, 3, 5}
	a4 := a3
	fmt.Println(a3, a4)
	//a3和a4指向同一个数组
	a3[0] = 100
	fmt.Println(a3, a4)

	//切片的遍历
	//1.索引遍历
	for i := 0; i < len(a3); i++ {
		fmt.Println(a3[i])
	}
	//2. for range遍历
	for i, v := range a4 {
		fmt.Println(i, ":", v)
	}

	//切片的追加
	s10 := []string{"北京", "上海", "广州"}
	// s10[3] = "深圳"  //错误用法，panic: runtime error: index out of range [3] with length 3
	fmt.Printf("s10=%v len(s10)=%d cap(s10)=%d\n", s10, len(s10), cap(s10))
	s10 = append(s10, "沈阳")
	//调用append函数必须用原来的切片变量接收返回值
	//append追加元素，原来的底层数组放不下的时候，Go底层就会把底层数组换一个
	//如果新申请的容量（cap）大于2倍的旧容量（old.cap),最终容量（newcap）就是新申请的容量（cap）
	//否则判断，如果切片的长度小于1024，则最终容量（newcap）就是就容量（old.cap）的两倍，即newcap=doublecap
	//否则判断，如果就切片的长度大于等于1024，则最终容量（newcap）从旧容量（oldcap）的开始循环增加原来容量的1/4，直到最终容量（newcap）大于等于新申请的容量（cap）
	//如果最终容量（cap）计算溢出，则最终容量就是新申请的容量
	fmt.Printf("s10=%v len(s10)=%d cap(s10)=%d\n", s10, len(s10), cap(s10))

	s11 := []string{"长春", "天津"}
	s10 = append(s10, s11...) //...表示拆开
	fmt.Printf("s10=%v len(s10)=%d cap(s10)=%d\n", s10, len(s10), cap(s10))

	//Copy用法
	c1 := []int{1, 3, 5}
	c2 := c1 //赋值
	var c3 = make([]int, 3)
	copy(c3, c1)
	fmt.Println(c1, c2, c3)
	c1[0] = 100
	fmt.Println(c1, c2, c3)

	//将c1中的索引为1的3的这个元素删掉
	c1 = append(c1[:1], c1[2:]...)
	fmt.Println(c1)

	x1 := []int{1, 3, 5}
	x2 := x1[:]
	fmt.Println(x2, len(x2), cap(x2))
	//1.切片不保存具体的值
	//2.切片对应一个底层数组
	//3.底层数组都是占用一块连续的内存
	fmt.Printf("%p\n", &x2[0]) //地址指针相同
	x2 = append(x2[:1], x2[2:]...)
	fmt.Printf("%p\n", &x2[0]) //地址指针相同
	fmt.Println(x2, len(x2), cap(x2))

	//指针
	//1. &取地址
	//2. *：根据地址取值
	n := 19
	p := &n
	fmt.Println(p)
	fmt.Printf("%T\n", p)
	//2. *:根据地址取值
	m := *p
	fmt.Println(m)
	fmt.Printf("%T\n", m)

	slice1()

}

func slice1() {
	// 数据类型为 切片，长度为5，容量为10
	a := make([]int, 5, 10)
	fmt.Println(a, cap(a), len(a)) // out put : [0 0 0 0 0] 10 5

	// 切片追加元素，当超过原来的的容量的时候，会翻倍扩容，但不是一定翻倍，如果容量太大不会再翻倍
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	// 再对值进行修改
	for i := 0; i < 10; i++ {
		a[i] = i
	}
	fmt.Println(a, cap(a), len(a)) // [0 1 2 3 4 5 6 7 8 9 5 6 7 8 9] 20 15

	countWords()
}

func countWords() {
	//Count "How do you do"单词出现的次数
	s1 := "How do you do"
	//吧字符串按照空格切割得到切片
	s2 := strings.Split(s1, " ")
	m1 := make(map[string]int, 10)
	for _, w := range s2 {
		//如果原来map中不存在w这个key那么出现次数=1
		if _, ok := m1[w]; !ok {
			m1[w] = 1
		} else {
			m1[w]++
		}
		//如果map中存在w这个key，那么出现次数+1
	}
	//累计出现的次数
	for key, value := range m1 {
		fmt.Println(key, value)
	}
}
