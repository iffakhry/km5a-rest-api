package main

import (
	"fakhry/unit-test/calculate"
	"fmt"
)

func main() {
	bil1 := -10
	bil2 := -20
	resultAdd := calculate.Addition(bil1, bil2)
	fmt.Println(resultAdd)
}
