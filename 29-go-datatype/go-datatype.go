package main

import (
	"fmt"
	"strconv"
	"strings"
)

//16进制080对应的2进制为10000000（快速算法一位16进制转换成4位2进制左边补0，即0000 1000 0000）所以对应为编号为8的CPU即CPU7（第8颗）

func main() {
	// hex := "ffffffffff"
	// value, err := strconv.ParseUint(hex, 16, 64) /////(hex, 16, 64)
	// if err != nil {
	// 	fmt.Printf("Conversion failed: %s\n", err)
	// } else {
	// 	fmt.Printf("Hexadecimal '%s' is integer %d (%X)\n",
	// 		hex, value, value)
	// }

	// b := make([]byte, 0)
	// b = append(b, 255)
	// b = append(b, 10)
	// fmt.Println(b)
	a := 4
	fmt.Printf("%04b\n", a)
	fmt.Printf("%04b", a>>4)
	//convertCPUBitMaskToByte("7fffffffffff")
	//cpuBitsMask := convertCPUBitMaskToByte("7fffffffffff")
	//isolatedCPU := convertIsolatedCPURange2CPUList("0,1,3-4,11-16,23-27")
	//isolatedCPU := convertIsolatedCPURange2CPUList("1")
	//assertIsolateCPUCoresAffectedBitMask(cpuBitsMask, isolatedCPU)
	//assertDefaultIRQSMPAffinityAffectedBitMask(cpuBitsMask, isolatedCPU, isolatedCPU)
	//getDefaultSMPAffinityBitMaskbyCPUCores()
}

func assertIsolateCPUCoresAffectedBitMask(cpuBitsMask []byte, isolatedCPU []byte) string {

	fmt.Println(isolatedCPU)
	//Isolated CPU Range, 0,1,3-4,11-16,23-27
	//          27 26 25 24 ---------------------------------3 2 1 0
	//          27%6=3
	//[1111     1111         1111 1111 1111 1111 1111         1111] cpuBitMask
	//[0000     1111         1000 0001 1111 1000 0001         1011] isolatedCPU
	//--------------------------------------------------------------
	//[1111     0000         0111 1110 0000 0111 1110         0100] affinityCPUMask
	//  0         1            2    3   4     5   6             7    cpuBitMaskGroupsIndex
	//            6            5    4   3     2   1             0    isolatedCPUIndex
	//     maxValueOfIsolatedCPUIndex
	var affinityCPUMask string
	totalCPUBitMaskGroups := len(cpuBitsMask)
	totalIsolatedCPUNum := len(isolatedCPU)

	fmt.Printf("Total isolated CPUs is: %v\n", totalIsolatedCPUNum)
	fmt.Printf("The max CPU that isolated is : %v\n", int(isolatedCPU[totalIsolatedCPUNum-1]))

	//The max CPU number is 27, Index is 15
	maxValueOfIsolatedCPUIndex := int(isolatedCPU[totalIsolatedCPUNum-1]) / 4
	fmt.Printf("totalCpuGroupNum is: %v\nmaxCpuGroupIndicator is: %v\n", totalCPUBitMaskGroups, maxValueOfIsolatedCPUIndex)
	maxValueOfCPUBitMaskGroupsIndex := totalCPUBitMaskGroups - 1
	for i := totalIsolatedCPUNum - 1; i >= 0; i-- {
		fmt.Printf("i is %v\n", i)
		isolatedCPUIndex := int(isolatedCPU[i]) / 4
		fmt.Printf("cpuGroupIndex is %v\n", isolatedCPUIndex)

		cpuBitsMaskIndex := maxValueOfCPUBitMaskGroupsIndex - isolatedCPUIndex

		fmt.Printf("cpuBitsMaskIndex is %v\n", cpuBitsMaskIndex)
		fmt.Printf("Before int(cpuBitsMask[cpuBitsMaskIndex]) is %04b\n", int(cpuBitsMask[cpuBitsMaskIndex]))
		fmt.Printf("isolatedCPU is %v\n", isolatedCPU[i])
		fmt.Printf("isolatedCPU  mode is %v\n", isolatedCPU[i]%4)

		// 3 => 1000 2=>0100 1=>0010 0=>0000
		modIsolatedCPUby4 := int(isolatedCPU[i] % 4)
		var isolatedCPUMask int
		switch modIsolatedCPUby4 {
		case 3:
			isolatedCPUMask = 8
		case 2:
			isolatedCPUMask = 4
		case 1:
			isolatedCPUMask = 2
		case 0:
			isolatedCPUMask = 1
		}
		valueOfCPUBitsMaskOnIndex := int(cpuBitsMask[cpuBitsMaskIndex]) ^ isolatedCPUMask
		fmt.Printf(" %04b ^ %04b = %04b\n", cpuBitsMask[cpuBitsMaskIndex], isolatedCPUMask, valueOfCPUBitsMaskOnIndex)
		cpuBitsMask[cpuBitsMaskIndex] = byte(valueOfCPUBitsMaskOnIndex)
	}
	fmt.Printf("cpuBitsMask is: %04b\n", cpuBitsMask)
	cpuBitsMaskStr := fmt.Sprintf("%x", cpuBitsMask)
	affinityCPUMask = strings.ReplaceAll(cpuBitsMaskStr, "0", "")
	fmt.Printf("affinityCPUMask is: %s\n", affinityCPUMask)
	return affinityCPUMask

}

func convertIsolatedCPURange2CPUList(isolatedCPURange string) []byte {

	//Get a separated cpu number list
	cpuList := make([]byte, 0, 8)
	//From [1,2,4-5,12-17,24-28,30-32]
	//To   [1 2 4 5 12 13 14 15 16 17 24 25 26 27 28 30 31 32]

	cpuRangeList := strings.Split(isolatedCPURange, ",")

	for i := 0; i < len(cpuRangeList); i++ {
		//if CPU range is 12-17 which contain "-"

		if strings.Contains(cpuRangeList[i], "-") {

			//Ignore such senario when cpu setting as 45-,-46
			if strings.HasPrefix(cpuRangeList[i], "-") {
				continue
			}
			//startCPU is 12
			//endCPU is 17
			//the CPU range must be two numbers
			cpuRange := strings.Split(cpuRangeList[i], "-")
			endCPU, _ := strconv.Atoi(cpuRange[1])
			startCPU, _ := strconv.Atoi(cpuRange[0])
			for i := 0; i <= endCPU-startCPU; i++ {
				cpus := startCPU + i
				cpuList = append(cpuList, byte(cpus))
			}

		} else {
			cpus, _ := strconv.Atoi(cpuRangeList[i])

			//Ignore 1,2,<no number>
			if len(cpuRangeList[i]) != 0 {
				cpuList = append(cpuList, byte(cpus))
			}
		}
	}
	return cpuList
}

func convertCPUBitMaskToByte(cpuHexMask string) []byte {

	cpuHexMaskChars := []rune(cpuHexMask)
	cpuBitsMask := make([]byte, 0)
	cpuNum := 0

	for i := 0; i < len(cpuHexMaskChars); i++ {

		switch cpuHexMaskChars[i] {
		case 'f':
			cpuBitsMask = append(cpuBitsMask, 15)
			cpuNum = cpuNum + 4
		case '7':
			cpuBitsMask = append(cpuBitsMask, 7)
			cpuNum = cpuNum + 3
		case '3':
			cpuBitsMask = append(cpuBitsMask, 3)
			cpuNum = cpuNum + 2
		case '1':
			cpuBitsMask = append(cpuBitsMask, 1)
			cpuNum = cpuNum + 1
		}

	}

	fmt.Printf("The total CPU number is %v\nThe CPU HexMask is:\n%s\nThe CPU BitsMask is:\n%b\n", cpuNum, cpuHexMask, cpuBitsMask)
	return cpuBitsMask
}

func assertDefaultIRQSMPAffinityAffectedBitMask(cpuBitsMask []byte, isolatedCPU []byte, defaultIRQSMPAffinity []byte) string {

	fmt.Println(isolatedCPU)
	//Isolated CPU Range, 0,1,3-4,11-16,23-27
	//          27 26 25 24 ---------------------------------3 2 1 0
	//          27%6=3
	//[1111     1111         1111 1111 1111 1111 1111         1111] cpuBitMask
	//[0000     1111         1000 0001 1111 1000 0001         1011] isolatedCPU
	//--------------------------------------------------------------
	//[0000     1111         1000 0001 1111 1000 0001         1011] affinityCPUMask
	//  0         1            2    3   4     5   6             7    cpuBitMaskGroupsIndex
	//            6            5    4   3     2   1             0    isolatedCPUIndex
	//     maxValueOfIsolatedCPUIndex

	//[1111     1111         1111 1111 1111 1111 1111         1111] cpuBitMask
	//[0000     0000         0000 0000 0000 0000 0000         0010] isolatedCPU
	//--------------------------------------------------------------
	//[0000     0000         0000 0000 0000 0000 0000         0010] affinityCPUMask
	//  0         1            2    3   4     5   6             7    cpuBitMaskGroupsIndex
	//                                                          0    isolatedCPUIndex
	//                                                maxValueOfIsolatedCPUIndex

	var affinityCPUMask string
	totalCPUBitMaskGroups := len(cpuBitsMask)
	totalIsolatedCPUNum := len(isolatedCPU)

	fmt.Printf("Total isolated CPUs is: %v\n", totalIsolatedCPUNum)
	fmt.Printf("The max CPU that isolated is : %v\n", int(isolatedCPU[totalIsolatedCPUNum-1]))
	isolatedCPUMaskGroup := make([]byte, 0, 8)

	for i := 0; i < totalCPUBitMaskGroups; i++ {
		//Initial all bits to zero of isolatedCPUMask first
		isolatedCPUMaskGroup = append(isolatedCPUMaskGroup, byte(int(cpuBitsMask[i])&0))
	}

	fmt.Printf("The initial isolatedCPUMask is %04b\n", isolatedCPUMaskGroup)

	maxValueOfCPUBitMaskGroupsIndex := totalCPUBitMaskGroups - 1
	for i := totalIsolatedCPUNum - 1; i >= 0; i-- {

		isolatedCPUIndex := int(isolatedCPU[i]) / 4

		cpuBitsMaskIndex := maxValueOfCPUBitMaskGroupsIndex - isolatedCPUIndex

		// 3 => 1000 2=>0100 1=>0010 0=>0000
		modIsolatedCPUby4 := int(isolatedCPU[i] % 4)
		var isolatedCPUMask int
		switch modIsolatedCPUby4 {
		case 3:
			isolatedCPUMask = 8
		case 2:
			isolatedCPUMask = 4
		case 1:
			isolatedCPUMask = 2
		case 0:
			isolatedCPUMask = 1
		}

		fmt.Printf("%04b | %04b = %04b\n", isolatedCPUMaskGroup[cpuBitsMaskIndex], isolatedCPUMask, int(isolatedCPUMaskGroup[cpuBitsMaskIndex])|isolatedCPUMask)
		valueOfCPUBitsMaskOnIndex := int(isolatedCPUMaskGroup[cpuBitsMaskIndex]) | isolatedCPUMask
		isolatedCPUMaskGroup[cpuBitsMaskIndex] = byte(valueOfCPUBitsMaskOnIndex)

	}

	//Remove additional 0 in the isolatedCPUMaskGroup
	fmt.Printf("cpuBitsMask is: %04b\n", isolatedCPUMaskGroup)
	cpuBitsMaskStr := fmt.Sprintf("%x", isolatedCPUMaskGroup)
	cpuBitsMaskRune := []rune(cpuBitsMaskStr)
	bitsMaskChars := make([]byte, 0, 2)

	for i := 1; i < len(cpuBitsMaskRune); i = i + 2 {
		bitsMaskChars = append(bitsMaskChars, byte(cpuBitsMaskRune[i]))
	}
	affinityCPUMask = fmt.Sprintf("%s\n", bitsMaskChars)
	fmt.Printf("affinityCPUMask is: %s", affinityCPUMask)
	return affinityCPUMask

}

func getDefaultSMPAffinityBitMaskbyCPUCores() {
	cpuCores := 20
	cpuHexMask := make([]byte, 0, 2)
	if cpuCores%4 != 0 {

		modCPUCoresby4 := int(cpuCores % 4)
		var cpuCoresMask int
		switch modCPUCoresby4 {
		case 3:
			cpuCoresMask = 7
		case 2:
			cpuCoresMask = 3
		case 1:
			cpuCoresMask = 1
		}
		cpuHexMask = append(cpuHexMask, byte(cpuCoresMask))
	}

	for i := 0; i < cpuCores/4; i++ {
		cpuHexMask = append(cpuHexMask, 15)
	}

	cpuHexMaskStr := fmt.Sprintf("%x", cpuHexMask)
	cpuHexMaskFmt := strings.ReplaceAll(cpuHexMaskStr, "0", "")
	fmt.Printf("%s", cpuHexMaskFmt)

}
