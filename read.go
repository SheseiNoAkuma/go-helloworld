package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type person struct {
	fname string
	lname string
}

func (p *person) setNames(firstName string, lastName string) {
	p.fname = firstName
	p.lname = lastName

	if len(firstName) > 20 {
		p.fname = firstName[:20]
	}

	if len(lastName) > 20 {
		p.lname = lastName[:20]
	}
}

func main() {

	filePath := getFileName()

	slice := extractPersons(filePath)

	for index, value := range slice {
		fmt.Printf("line %d first name: %s, last name %s\n", index, value.fname, value.lname)
	}
}

func extractPersons(filePath string) []person {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	slice := make([]person, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")

		if len(split) < 2 {
			log.Printf("this line is incorrect %s \n", split)
		} else {
			p := new(person)
			p.setNames(split[0], split[1])
			slice = append(slice, *p)
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return slice
}

func getFileName() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("input your file name: ")
	name, _ := reader.ReadString('\n')
	return strings.TrimSpace(name)
}
