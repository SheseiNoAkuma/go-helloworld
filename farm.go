package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat() string
	Move() string
	Speak() string
}

type Cow struct{}

func (a Cow) Eat() string   { return "grass" }
func (a Cow) Move() string  { return "walk" }
func (a Cow) Speak() string { return "moo" }

type Bird struct{}

func (a Bird) Eat() string   { return "worms" }
func (a Bird) Move() string  { return "fly" }
func (a Bird) Speak() string { return "peep" }

type Snake struct{}

func (a Snake) Eat() string   { return "mice" }
func (a Snake) Move() string  { return "slither" }
func (a Snake) Speak() string { return "hsss" }

type command struct {
	action     string
	name       string
	parameter3 string
}

func (c command) isEmpty() bool {
	return c.action != "" && c.name != "" && c.parameter3 != ""
}

func usage() {
	fmt.Println("commands must be formed by 3 strings separated by space:")
	fmt.Println(" 1) the action to perform (newanimal/query)")
	fmt.Println(" 2) the name of the animal to create / query")
	fmt.Println(" 3.1) for newanimal: the spice of the animal (cow/bird/snake)")
	fmt.Println(" 3.2) for query: the information you want to gather (eat/move/speak)")
}

var farm = map[string]Animal{}

func main() {
	for {
		command := askForCommand()
		if !command.isEmpty() {
			usage()
			continue
		}

		switch command.action {
		case "newanimal":
			if newAnimal(command.name, command.parameter3) {
				fmt.Printf("animal %s of spiece %s has been created\n", command.name, command.parameter3)
			} else {
				usage()
			}
		case "query":
			if !queryAnimal(command.name, command.parameter3) {
				usage()
			}
		default:
			fmt.Printf("%s is not a valid action\n", command.action)
			usage()
		}
	}
}

func queryAnimal(name string, information string) bool {
	if animal, ok := farm[name]; ok {
		switch information {
		case "eat":
			fmt.Printf("%s eats %s\n", name, animal.Eat())
		case "move":
			fmt.Printf("%s moves by %s\n", name, animal.Move())
		case "speak":
			fmt.Printf("%s sound like %s\n", name, animal.Speak())
		default:
			fmt.Printf("animals in the farm does not have this information %s\n", information)
			return false
		}
	} else {
		fmt.Printf("the farm dos not have an animal called %s --> %s\n", name, farm)
	}
	return true
}

func newAnimal(name string, spice string) bool {
	switch spice {
	case "cow":
		farm[name] = Cow{}
	case "bird":
		farm[name] = Bird{}
	case "snake":
		farm[name] = Snake{}
	default:
		fmt.Printf("%s is not a valid spiece\n", spice)
		return false
	}
	return true
}

func askForCommand() command {

	var slice []string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("enter your command: ")
	scanner.Scan()
	line := scanner.Text()
	for _, x := range strings.Fields(line) {
		slice = append(slice, x)
	}
	if len(slice) < 3 {
		fmt.Println("your command is invalid")
		return command{}
	}

	return command{slice[0], slice[1], slice[2]}
}
