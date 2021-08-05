### My first steps with go


some notes taken along the way

### Advantages of GO
 * Run fast (it's compiled)
 * Garbage collector (usually not available in compiled languages)
 * Is OO but object are simplified (no inheritance, no constructor, no generics)
 * Concurrency is build in and very efficient

### Pointers
Yes Go has pointers, `&` gives you the pointer to an object, `*` gives you the data from the pointer

### Enum
There is no such things as Enum, the closest thing is iota

```go
package iotaexample

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
```
### Defer

A defer statement defers the execution of a function until the surrounding function returns.

The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

```go
package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
```
