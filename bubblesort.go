package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	slice := loadSlicex()
	BubbleSortx(slice)
	fmt.Print(slice)
}

func BubbleSortx(slice []int) {

	hasSwap := false
	for {
		for index, value := range slice {
			if index+1 < len(slice) && value > slice[index+1] {
				Swapx(slice, index)
				hasSwap = true
			}
		}

		if hasSwap {
			hasSwap = false
		} else {
			break
		}
	}
}

func Swapx(slice []int, index int) {
	tmpValue := slice[index]
	slice[index] = slice[index+1]
	slice[index+1] = tmpValue
}

func loadSlicex() []int {

	var slice []int
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("enter your list separated by space: ")
	scanner.Scan()
	line := scanner.Text()
	for _, x := range strings.Fields(line) {
		intVar, err := strconv.Atoi(x)
		if err != nil {
			log.Fatal(err)
		}
		slice = append(slice, intVar)
	}
	return slice
}
