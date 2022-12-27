package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func useBufIO() {
	var s string
	reader := bufio.NewReader(os.Stdin)
	s, _ = reader.ReadString('\n')
	fmt.Printf("Your input content is: %s\n", s)
}

func insertInFile() {
	fileObj, err := os.OpenFile("/Users/cuiliquan/goproject//src/github.com/liqcui/go-study/24-gofile/insertFile.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("Open file failed, err: %v", err)
		return
	}
	defer fileObj.Close()

	//因为没有办法直接在文中插入内容，所以要借助一个临时文件
	tmpFile, err := os.OpenFile("/tmp/InsertFile.tmp", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Create tmp file failed, error:%v\n", err)
	}
	defer tmpFile.Close()

	//读取源文件，写入新文件
	var ret [1]byte
	n, err := fileObj.Read(ret[:])
	if err != nil {
		fmt.Printf("Read from file failed, err: %v", err)
		return
	}

	//写入临时文件
	tmpFile.Write(ret[:n])
	fileObj.Seek(1, 0)
	// var s []byte
	s := []byte{'c'}
	tmpFile.Write(s)
	//接着把原文件后续的内容写入临时文件
	var x [1024]byte
	for {
		n, err := fileObj.Read(x[:])
		if err == io.EOF {
			tmpFile.Write(x[:n])
			break
		}
		if err != nil {
			fmt.Printf("Read from file failed, err:%v\n", err)
		}
		tmpFile.Write(x[:n])
	}
	//原文件后续的也写入了临时文件中
	fileObj.Close()
	tmpFile.Close()
	os.Rename("/tmp/InsertFile.tmp", "/Users/cuiliquan/goproject//src/github.com/liqcui/go-study/24-gofile/insertFile.txt")

}
func main() {

	// fileObj, err := os.Open("/etc/hosts")
	// if err != nil {
	// 	fmt.Printf("open file failed, err:%v", err)
	// 	return
	// }

	// //关闭文件
	// defer fileObj.Close()
	// //读文件
	// //var tmp=make([]byte,128)
	// var tmp = make([]byte, 128)
	// n, err := fileObj.Read(tmp[:])
	// if err != nil {
	// 	fmt.Printf("Read from file failed, err:%v", err)
	// }
	// fmt.Printf("读了%d个字节\n", n)

	// content, err := ioutil.ReadFile("/etc/hosts")
	// if err != nil {
	// 	fmt.Println("read file failed, err:", err)
	// 	return
	// }
	// fmt.Println(string(content))

	// fileObj, err = os.OpenFile("./1.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	// if err != nil {
	// 	fmt.Println("open file failed")
	// }
	// defer fileObj.Close()
	// //创建一个写对象

	//useBufIO()
	insertInFile()
}
