package main

import (
	"fmt"
)

func main() {

	arrayInt := [...]int{1, 1, 2, 3, 5, 8}

	for index, value := range arrayInt {
		fmt.Printf("looping over array at index %d have value %d\n", index, value)
	}

	arrayString := [...]string{"a", "b", "c", "d", "e", "f"}

	slice1 := arrayString[1:4]

	for index, value := range slice1 {
		fmt.Printf("looping over slice at index %d have value %s\n", index, value)
	}
	println(len(slice1))
	println(cap(slice1))

	fmt.Println("------------------------")
	slice2 := make([]int, 3)
	fmt.Print(slice2, "\n")
	fmt.Println("------------------------")
	slice2 = append(slice2, 99)
	fmt.Print(slice2, "\n")

}
