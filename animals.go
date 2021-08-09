package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal1 struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal1) Eat() string {
	return a.food
}

func (a Animal1) Move() string {
	return a.locomotion
}

func (a Animal1) Speak() string {
	return a.noise
}

type command1 struct {
	animal string
	action string
}

func main() {
	zoo := map[string]Animal1{
		"cow":   {food: "grass", locomotion: "walk", noise: "moo"},
		"bird":  {food: "worms", locomotion: "fly", noise: "peep"},
		"snake": {food: "mice", locomotion: "slither", noise: "hsss"},
	}

	for {
		command := askForCommandx()

		if animal, ok := zoo[command.animal]; ok {
			switch command.action {
			case "eat":
				fmt.Printf("the %s eats %s\n", command.animal, animal.Eat())
			case "move":
				fmt.Printf("the %s moves by %s\n", command.animal, animal.Move())
			case "speak":
				fmt.Printf("the %s sound like %s\n", command.animal, animal.Speak())
			default:
				fmt.Printf("this action is unknown %s\n", command.action)
			}
		} else {
			fmt.Printf("this animal is unknown %s\n", command.animal)
		}
	}
}

func askForCommandx() command1 {

	var slice []string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("enter the animal and the action separated by spaces: ")
	scanner.Scan()
	line := scanner.Text()
	for _, x := range strings.Fields(line) {
		slice = append(slice, x)
	}
	return command1{slice[0], slice[1]}
}
