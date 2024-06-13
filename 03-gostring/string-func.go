package main

import (
	"fmt"
	"strings"
)

func ProcessNum() {
	str_a := `
	18
	18
	78
	80
	`
	aaa := strings.ReplaceAll(str_a, "\n", ",")

	for i, v := range aaa {
		fmt.Printf("%v,%c", i, v)
	}
}
