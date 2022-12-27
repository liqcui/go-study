package main

import "fmt"

func funcA() {
	fmt.Println("func a")
}

func funcB() {

	defer func() {
		err := recover()
		fmt.Println(err)
		fmt.Println("Free the database connection ...")
	}()
	panic("Disaster error")
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
