package main

import (
	"fmt"
)

type student struct {
	id   int64
	name string
}

//造一个学生的管理者
type studentMgr struct {
	allStudent map[int64]student
}

//查看学生
func (s studentMgr) showStudents() {
	for _, stu := range s.allStudent {
		fmt.Printf("学号：%d 姓名：%s", stu.id, stu.name)
	}
}

//增加学生
func (s studentMgr) addStudents() {
	var (
		stuID   int64
		stuName string
	)
	fmt.Print("请输入学号：")
	fmt.Scanln(&stuID)
	fmt.Print("请输入姓名：")
	fmt.Scanln(&stuName)
	newStu := student{
		id:   stuID,
		name: stuName,
	}
	s.allStudent[newStu.id] = newStu
	fmt.Println("添加成功")
}

//修改学生
func (s studentMgr) modifyStudents() {
     //1. 获取用户输入的学号
     var stuID int64
	 fmt.Print("请输入学号")
	 fmt.Scanln(&stuID)
	 stuObj,ok :=s.allStudent[stuID]
	 if !ok {
		 fmt.Println("查无此人")
		 return 
	 }
	 fmt.Println("你要修改的学生信息如下：学号:%d 姓名：%s \n",stuObj.id,stuObj.name)
    var newName string
    fmt.Scanln(&newName)
	stuObj.name=newName
	s.allStudent[stuID]=stuObj
	
}

//删除学生
func (s studentMgr) deleteStudents() {
 var stuID int64
 fmt.Print("请输入要删除的学生的学号：")
 fmt.Scanln(&stuID)
 _,ok:=s.allStudent[stuID]
 if !ok{
	 fmt.Println("查无此人")
	 return
 }
 delete(s.allStudent,stuID)
 fmt.Println("删除成功")
}
