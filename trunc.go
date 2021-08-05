package main

import "fmt"

func main() {

	var inputNumber float64

	_, err := fmt.Scan(&inputNumber)

	if err != nil {
		fmt.Printf("an error occurred %s", err)
	} else {
		fmt.Println(int(inputNumber))
	}

}
