package main

import (
	"fmt"
	"sync"
)

func inc(x *int, wg *sync.WaitGroup) {
	*x += 1
	wg.Done()
}

func update(x *int, wg *sync.WaitGroup) {
	*x = 4
	wg.Done()
}

/* Depending on when x is read/updated in the two go routines, program can output 1,4,5
 *	case 1: x is read in inc() => x is updated to 4 in update() => inc() increments/writes back 1 to x
 *	case 4: x is read/incremented/updated to 1 in inc() => x is updated to 4 in update()
 *  case 5: x is updated to 4 in update() => x is read/incremented/updated to 5 in inc()
 */
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	x := 0
	go inc(&x, &wg)
	go update(&x, &wg)
	wg.Wait()
	fmt.Println("x is ", x)
}
