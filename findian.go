package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter your string:")

	text, _ := reader.ReadString('\n')

	toLower := strings.ToLower(strings.TrimSpace(text))

	startsWithI := strings.HasPrefix(toLower, "i")
	endsWithN := strings.HasSuffix(toLower, "n")
	containsA := strings.Contains(toLower, "a")

	if startsWithI && endsWithN && containsA {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
