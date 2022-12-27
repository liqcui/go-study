package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	var m1 map[string]int
	fmt.Println(m1 == nil)        //还没有进行初始化，还没在内存中分配空间
	m1 = make(map[string]int, 10) //预先估计好容量，避免在运行中动态分配扩容
	m1["Tom"] = 90
	m1["Jerry"] = 100
	fmt.Println(m1)
	fmt.Println(m1["Tom"])

	m1["John"] = 98
	value, OK := m1["John"]
	if !OK {
		fmt.Println("查无此人")
	} else {
		fmt.Println(value)
	}
	//map的遍历
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	//删除
	delete(m1, "Jerry")
	fmt.Println(m1)
	delete(m1, "not exist")
	fmt.Println(m1)

	//示例
	var scoremap = make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成一个stu0开始的字符串
		value := rand.Intn(100)          //生成0-99的随机数
		scoremap[key] = value
	}
	fmt.Println(scoremap)

	//取出map中的所有key存放到keys里面。
	var keys = make([]string, 0, 200)
	for key := range scoremap {
		keys = append(keys, key)
	}

	//对切片进行排序
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println(key, scoremap[key])
	}

	//元素类型为map的切片
	var s1 = make([]map[int]string, 10, 10)
	s1[0] = make(map[int]string, 1)
	s1[0][0] = "Shahe"
	fmt.Println(s1)

	//值为切片类型的map
	var s2 = make(map[string][]int, 10)
	s2["BJ"] = []int{1, 2, 3}
	fmt.Println(s2)

}
