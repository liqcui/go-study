package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readFromFilebyBufio() {
	fileObj, err := os.Open("/etc/hosts")
	if err != nil {
		fmt.Printf("Open file failed, error:%v", err)
		return
	}

	defer fileObj.Close()
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("Read file failed, error:%v", err)
			return
		}
		fmt.Print(line)
	}
}

func readFromFileByIOUtil() {
	ret, err := ioutil.ReadFile("/etc/hosts")
	if err != nil {
		fmt.Printf("Read file failed, err:%v\n", err)
		return
	}
	fmt.Println(string(ret))
}

func createFile() {
	fileObj, err := os.OpenFile("/tmp/gofile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Open file failed:%v\n", err)
		return
	}
	fileObj.Write([]byte("binary file was written\n"))
	fileObj.WriteString("string write to file\n")
	fileObj.Close()

}

func createYamlByString(s string) {
	fileObj, err := os.OpenFile("/var/tmp/sro-cli.yaml", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Create sro-cli.yaml failed:%v\n", err)
		return
	}
	fileObj.Write([]byte(s))
	//fileObj.WriteString(s)
	fileObj.Close()

}

// 	if err != nil {
// 		fmt.Printf("Openfile failed, err:%v", err)
// 	}
// 	defer fileObj.Close()
// 	fileObj.Seek(3, 0)
// 	var ret [1]byte
// 	n, err := fileObj.Read(ret[:])
// 	if err != nil {
// 		fmt.Printf("Read from file failed,err:%v", err)
// 		return
// 	}
// 	fmt.Println(string(ret[:n]))
// }

func main() {
	// readFromFilebyBufio()
	// fmt.Println("===================================")
	readFromFileByIOUtil()
	// fmt.Println("===================================")
	// createFile()
	// fmt.Println("===================================")
	// createYamlByString(SROOperatorGroup)
}
