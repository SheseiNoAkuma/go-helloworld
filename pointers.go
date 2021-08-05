package main

import (
	"fmt"
)

func main() {

	x := 5

	//var  myPointer = new(int)
	//*myPointer = x

	var myPointer = &x

	fmt.Println(x)
	fmt.Println(*myPointer)

	fmt.Printf("is my pointer equal? %t", *myPointer == x)

}
