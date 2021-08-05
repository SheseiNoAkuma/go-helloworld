package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	name := requestInput("enter a name: ")
	address := requestInput("enter the address: ")

	personMap := map[string]string{"name": name, "address": address}
	marshal, _ := json.Marshal(personMap)

	fmt.Println(string(marshal))
}

func requestInput(label string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(label)
	name, _ := reader.ReadString('\n')
	return strings.TrimSpace(name)
}
