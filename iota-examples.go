package main

import "fmt"

func main() {

	type daysOfTheWeek int

	const (
		monday daysOfTheWeek = iota
		tuesday
		wednesday
		thursday
		friday
		saturday
		sunday
	)

	for i := monday; i <= sunday; i++ {

		switch i {
		case monday:
			fmt.Println(monday)
		case tuesday:
			fmt.Println(tuesday)
		case wednesday:
			fmt.Println(wednesday)
		case thursday:
			fmt.Println(thursday)
		case friday:
			fmt.Println(friday)
		case saturday:
			fmt.Println(saturday)
		case sunday:
			fmt.Println(sunday)
		default:
			fmt.Println("I don't know this day")
		}
	}

}
