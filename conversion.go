package main

import (
	"fmt"
)

func main() {

	var x int8 = 1
	var y int16 = 10000

	x = int8(y)

	fmt.Print(x)

}
