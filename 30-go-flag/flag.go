package main

import (
	"flag"
	"fmt"
)

func main() {
	// 定义几个变量，用于接收命令行的参数值
	var user string
	var pwd string
	var host string
	var port int

	// &user 就是接收用户命令行中输入的 -user 后面的参数值
	// "user" 就是 -user 指定的参数
	// "" 默认值
	// "用户名，默认为空" 说明
	flag.StringVar(&user, "user", "", "用户名，默认为空")
	flag.StringVar(&pwd, "password", "", "密码，默认为空")
	flag.StringVar(&host, "host", "localhost", "主机名，默认为 localhost")
	flag.IntVar(&port, "port", 3306, "duan端口号，默认3306")

	// 【必须调用】从 arguments 中解析注册的 flag
	flag.Parse()

	// 输出结果
	fmt.Printf("\n user=%v \n password=%v \n host=%v \n port=%v \n", user, pwd, host, port)
}
