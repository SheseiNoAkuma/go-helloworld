package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("initializing the program.... (digit x to end the program)")
	reader := bufio.NewReader(os.Stdin)
	slice := make([]int, 0)

	for {
		fmt.Print("digit your number then press enter: ")
		text, _ := reader.ReadString('\n')
		input := strings.ToLower(strings.TrimSpace(text))
		if input == "x" {
			break
		}

		asInt, err := strconv.Atoi(input)
		if err != nil {
			fmt.Printf("your input is not a valid int %s\n\n", err)
			continue
		}

		slice = append(slice, asInt)
		sort.Ints(slice)
		fmt.Print(slice, "\n")
	}
	fmt.Println("program completed")
}
