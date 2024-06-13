package main

import (
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {
	//算数运算符
	var (
		a = 5
		b = 3
	)
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a++ //单独语句，不能放在=右边赋值===> a=a+1
	b-- //单独语句，不能放在=右边赋值===> b=b-1
	fmt.Println(a, b)

	//关系运算符
	var (
		age = 21
	)
	if age < 18 && age > 60 {
		fmt.Println("Have to work")
	} else {
		fmt.Println("No need to work")
	}
	if age < 18 || age > 60 {
		fmt.Println("Have to work")
	} else {
		fmt.Println("No need to work")
	}
	//取反
	isOK := false
	fmt.Println(isOK)
	fmt.Println(!isOK)
	//位运算符，针对二进制
	//5的二进制是101
	//2的二进制是10
	//1的二进制是00000 0001
	fmt.Println(5 & 2)   //按位与（两位均为1才为1）
	fmt.Println(5 | 2)   //按位或（两位有一位为1才为1）
	fmt.Println(5 ^ 2)   //按位与（两位不一样才为1）
	fmt.Println(5 << 1)  //将二进制左移1位，从1010=>10
	fmt.Println(1 << 10) //将二进制左移10位，从1000 0000=>1024
	fmt.Println(5 >> 3)  //将二进制位右移3位，1 =>1

	x := 13
	x += 1
	fmt.Println(x)
	x -= 1
	fmt.Println(x)
	x *= 2
	fmt.Println(x)
	x /= 2
	fmt.Println(x)
	x %= 2
	fmt.Println(x)

	var desirednum string
	//var WindowsNodeNum int
	desirednum = "3"
	kindnames := "abc"
	//If there are windows worker nodes, the desired daemonset should be linux node's num
	WindowsNodeNum := 2
	if WindowsNodeNum > 0 && true {

		//Exclude windows nodes
		fmt.Printf("%v desirednum is: %v\n", kindnames, desirednum)
		desiredlinuxworkerNum, _ := strconv.Atoi(desirednum)
		fmt.Printf("desiredlinuxworkerNum is:%v\n", desiredlinuxworkerNum)
		fmt.Println(desiredlinuxworkerNum - WindowsNodeNum)
		desirednum = strconv.Itoa((desiredlinuxworkerNum - WindowsNodeNum))
	}
	fmt.Printf("desirednum:%T,%v", desirednum, desirednum)

	// a := 0
	// for a <= 20 {
	// 	a = a + 2

	// }
	fmt.Println("mask tranfer")
	fmt.Println(65535 >> 6) //2 cpus
	fmt.Println(65535 >> 4) //4 cpus
	fmt.Println(65535 >> 2) //6 cpus

	//Currently support 32core cpu
	smpbitMask := 0xffffffff
	smpbitMaskIntStr := fmt.Sprintf("%d", smpbitMask)
	fmt.Println(smpbitMaskIntStr)
	cpuNum := 4
	rightMoveBit := 32 - cpuNum
	smtMaskInt, err := strconv.Atoi(smpbitMaskIntStr)
	//smtMaskUnit, err := strconv.ParseUint(smtMaskIntStr, 10, 32)
	fmt.Println(smtMaskInt)
	fmt.Println(err)
	fmt.Println(smtMaskInt >> 14) //2 cpus
	fmt.Println(smtMaskInt >> 12) //4 cpus
	fmt.Println(smtMaskInt >> 8)  //8 cpus
	//smtMaskIntStr2 := strconv.Itoa(smtMaskInt >> 8)
	xx1 := fmt.Sprintf("0x"+"%x", smtMaskInt>>rightMoveBit)
	fmt.Println(xx1)

	smpAffinityMasdk2 := 0x2
	smtMaskInt2, _ := strconv.Atoi(xx1)
	// 1111 1111
	// 0000 0010
	// ---------
	// 1111 1101

	xx2 := fmt.Sprintf("%x", smtMaskInt2&smpAffinityMasdk2)
	fmt.Println(xx2)
	fmt.Println(xtob("f"))
	fmt.Println(xtob("2"))

	// smtMaskInt, _ = strconv.Atoi(xtob("f"))
	// cpuMaskInt, _ := strconv.Atoi(xtob("2"))
	a5 := 0xf
	b5 := 0x2
	xx3 := fmt.Sprintf("%x", a5^b5)
	fmt.Println(xx3)

	xx4 := fmt.Sprintf("%x", 15|2)
	fmt.Println(xx4)
	xx5 := fmt.Sprintf("%d", 0xf)
	fmt.Println(xx5)

	xx6 := "f"
	hx := hex.EncodeToString([]byte(xx6))
	fmt.Println("String to Hex Golang example")
	fmt.Println(xx6 + " ==> " + hx)

	xx7 := hexToInt("ff")
	xx8 := hexToInt("00000002")

	IntA, _ := strconv.Atoi(xx7)
	IntB, _ := strconv.Atoi(xx8)
	fmt.Println(strconv.Atoi(xx7))
	fmt.Println(strconv.Atoi(xx8))
	xx9 := fmt.Sprintf("%x", IntA^IntB)
	fmt.Println(xx9)

	isMatch := assertDefaultIRQSMPAffinityAffectedBitMask("fd", "2")
	fmt.Println(isMatch)

	b1 := false
	b2 := false
	fmt.Println(b1 == b2)

	fmt.Println(31 % 7)
	fmt.Println(30 % 7)
	fmt.Println(29 % 7)
	fmt.Println(28 % 7)

	fmt.Println(27 % 6)
	fmt.Println(26 % 6)
	fmt.Println(25 % 6)
	fmt.Println(24 % 6)

	fmt.Println(27 / 4)
	fmt.Println(26 / 4)
	fmt.Println(25 / 4)
	fmt.Println(24 / 4)

	fmt.Println(3 / 4)
	fmt.Println(2 / 4)
	fmt.Println(1 / 4)
	fmt.Println(0 / 4)

	fmt.Printf("%04b & %04b = %04b\n", 15, 8, 15&8)
	fmt.Printf("%04b | %04b = %04b\n", 15, 8, 15|8)

	fmt.Printf("3/5 = %d", 3/5)

}

// 二进制转十六进制
func btox(b string) string {
	base, _ := strconv.ParseInt(b, 2, 10)
	return strconv.FormatInt(base, 16)
}

// 十六进制转二进制
func xtob(x string) string {
	base, _ := strconv.ParseInt(x, 16, 10)
	return strconv.FormatInt(base, 2)
}

// 十六进制转二进制
func hexToInt(x string) string {
	base, _ := strconv.ParseInt(x, 16, 10)
	return strconv.FormatInt(base, 10)
}

func assertDefaultIRQSMPAffinityAffectedBitMask(defaultSMPBitMask string, isolatedCPU string) bool {
	var isMatch bool
	defaultSMPBitMaskStr := hexToInt(defaultSMPBitMask)
	isolatedCPUStr := hexToInt(isolatedCPU)

	defaultSMPBitMaskInt, _ := strconv.Atoi(defaultSMPBitMaskStr)
	isolatedCPUInt, _ := strconv.Atoi(isolatedCPUStr)

	if defaultSMPBitMaskInt == isolatedCPUInt {
		isMatch = true
		return isMatch
	} else {
		isMatch = false
		return isMatch
	}
}
