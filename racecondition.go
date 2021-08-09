package main

import (
	"fmt"
	"math/rand"
	"time"
)

// the two routines in this function have a race condition, it means the output of the program is not deterministic but depends on the order in witch
// the instructions in routines are executed, you could have different possible output, for example:

// if routine one is executed first you will se in log:
//routine one done: 1
//routine two done: 0
//main process completed: 0

// if routine two is executed first you will se in log:
//routine two done: -1
//routine one done: 0
//main process completed: 0

// but be careful there could also be other cases where one routine update the value of x and is suspended before printing

// NOTE the sleep it's not needed to have a race condition, I introduced it to make the problem more evident
func main() {
	rand.Seed(time.Now().UnixNano())
	x := 0

	channel := make(chan int, 2)

	// go routine one
	go func() {
		n := rand.Intn(500) // n will be between 0 and 500
		time.Sleep(time.Duration(n) * time.Millisecond)
		x++
		fmt.Printf("routine one done: %d\n", x)
		channel <- -x
	}()

	// go routine two
	go func() {
		n := rand.Intn(500) // n will be between 0 and 500
		time.Sleep(time.Duration(n) * time.Millisecond)
		x--
		fmt.Printf("routine two done: %d\n", x)
		channel <- -x
	}()

	raceX := <-channel
	fmt.Printf("main process completed: %d\n", raceX)
}
