package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type valuesAtT0 struct {
	acceleration float64
	velocity     float64
	displacement float64
}

func main() {
	values := loadValuesAtT0()

	fn := GenDisplaceFn(values.acceleration, values.velocity, values.displacement)

	for i := 0; i <= 10; i++ {
		fmt.Printf("your dispacement at t%d is %f\n", i, fn(float64(i)))
	}
}

func GenDisplaceFn(acceleration float64, velocity float64, displacement float64) func(time float64) float64 {
	return func(time float64) float64 {
		return displacement + (velocity * time) + (math.Pow(time, 2) * acceleration / 2)
	}
}

func loadValuesAtT0() valuesAtT0 {

	var slice []float64
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("enter acceleration, velocity at t0, displacement at t0 separated by spaces: ")
	scanner.Scan()
	line := scanner.Text()
	for _, x := range strings.Fields(line) {
		intVar, err := strconv.ParseFloat(x, 64)
		if err != nil {
			log.Fatal(err)
		}
		slice = append(slice, intVar)
	}
	return valuesAtT0{slice[0], slice[1], slice[2]}
}
