package main

import "fmt"

func main() {

	fmt.Println(MypPerson{"Bob", 20})
	fmt.Println(MypPerson{age: 30, name: "Alice"})
	fmt.Println(MypPerson{name: "Fred"})
	fmt.Println(&MypPerson{name: "Ann", age: 40})
}

type MypPerson struct {
	name string
	age  int
}
