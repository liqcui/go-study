package main

import (
	"fmt"
	"os"
)

var smr studentMgr

//菜单函数
func showMenu() {
	fmt.Println("welcome SMS!")
	fmt.Println(`
	1. 查看所有学生
	2. 添加学生
	3. 修改学生
	4. 删除学生
	5.退出
	`)
}

func main() {
	smr = studentMgr{
		allStudent: make(map[int64]student, 100),
	}
	for {
		showMenu()
		fmt.Print("请输入序号：")
		var choice int
		fmt.Scanln(&choice)
		fmt.Println("你输入的是：", choice)
		switch choice {
		case 1:
			smr.showStudents()
		case 2:
			smr.addStudents()
		case 3:
			smr.modifyStudents()
		case 4:
			smr.deleteStudents()
		case 5:
			os.Exit(1)
		default:
			fmt.Println("Invalid input")
		}
	}
}
