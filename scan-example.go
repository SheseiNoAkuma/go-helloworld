package main

import "fmt"

func main() {

	var apples int

	fmt.Println("how may apples?")

	scan, err := fmt.Scan(&apples)

	fmt.Println(scan)
	fmt.Println(err)

	fmt.Printf("apples %d\n", apples)

}
