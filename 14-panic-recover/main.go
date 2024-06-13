package main

import (
	"fmt"
)

func funcA() {
	fmt.Println("func a")
}

func funcB() {

	defer func() {
		err := recover()
		fmt.Println(err)
		fmt.Println("Free the database connection ...")
	}()

	timeDurationSec := 20
	if timeDurationSec <= 30 {
		//errors.New("specify the value of timeDurationSec great than 30")
		panic("ERROR: specify the value of timeDurationSec great than 30 for CheckAllNodepoolReadyByHostedClusterName")
	}
	fmt.Println("func b")
}

func funcC() {
	fmt.Println("func c")
}

func main() {
	funcA()
	funcB()
	funcC()
}
