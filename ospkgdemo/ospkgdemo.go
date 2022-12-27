package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

//检查是否文件存在
func checkIfFileExist() {
	Parms := os.Args
	if len(Parms) == 1 {
		fmt.Println("Please input a parameter")
		return
	}
	path_file := Parms[1]
	fileInfo, err := os.Stat(path_file)
	if err != nil {
		fmt.Printf("The file not found:%v", err)
	}
	//检查是否是Regular文件
	fileMode := fileInfo.Mode()
	if fileMode.IsRegular() {
		fmt.Println("The file is regular file")
	}
	//读取文件内容，打印到终端
	file, err := os.Open(path_file)
	if err != nil {
		fmt.Printf("Open file %v error:%v", file, err)
	}
	defer file.Close()

	buf := make([]byte, 18)
	if _, err := io.ReadFull(file, buf); err != nil {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
			fmt.Println(err)
		}
	}
	io.WriteString(os.Stdout, string(buf))
	fmt.Println()

}

func readFilebyLine(file string) error {
	fd, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer fd.Close()
	reader := bufio.NewReader(fd)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("Faild to read file %v with error:%v", fd, err)
			break
		}
		fmt.Print(line)
	}
	return nil
}

func main() {
	// openfile()
	//checkIfFileExist()
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Printf("usage: lByL <file1> [<file2> ...]\n")
		return
	}

	for _, file := range flag.Args() {
		err := readFilebyLine(file)
		if err != nil {
			fmt.Println(err)
		}
	}
}
