package main

import (
	"fmt"
	"sync"
)

type philosopher struct {
	number int
	left   *sync.Mutex
	right  *sync.Mutex
}

func (p philosopher) eat() {
	//wait for authorization by watchDog on channel one or two the first available
	select {
	case <-watchDogChannelOne:
		// release the authorization in the end
		defer func() { watchDogFeedbackOne <- true }()
	case <-watchDogChannelTwo:
		// release the authorization in the end
		defer func() { watchDogFeedbackTwo <- true }()
	}

	p.left.Lock()
	p.right.Lock()

	fmt.Printf("philosopher: starting to eat %d\n", p.number)
	fmt.Printf("philosopher: finishing to eat %d\n", p.number)

	p.left.Unlock()
	p.right.Unlock()

	wg.Done()
}

const philosophersNumber = 5

var watchDogChannelOne = make(chan bool)
var watchDogFeedbackOne = make(chan bool)
var watchDogChannelTwo = make(chan bool)
var watchDogFeedbackTwo = make(chan bool)
var wg sync.WaitGroup

func main() {
	chopsticks := make([]*sync.Mutex, philosophersNumber)
	for i := 0; i < philosophersNumber; i++ {
		chopsticks[i] = new(sync.Mutex)
	}

	philosophers := make([]philosopher, philosophersNumber)
	for i := 0; i < philosophersNumber; i++ {
		philosophers[i] = philosopher{
			number: i,
			left:   chopsticks[i],
			right:  chopsticks[(i+1)%philosophersNumber],
		}
	}

	go watchDog()
	for i := 0; i < 3; i++ {
		for index := range philosophers {
			wg.Add(1)
			go philosophers[index].eat()
		}
		//wait for all philosopher to have eaten one time before proceed with next course
		wg.Wait()
	}
}

// this is the host that give permissions to eat, it has two different channels each one after the permission require a feedback
// to ensure the philosopher is done eating before give a new permission
func watchDog() {
	for {
		fmt.Println("watch dog: giving one authorization to eat on channel one")
		watchDogChannelOne <- true
		fmt.Println("watch dog: giving one authorization to eat on channel two")
		watchDogChannelTwo <- true

		//todo little optimization can be done here, watchdog waits for both philosopher to be done before giving a new permission
		//we should have 2 watch dogs one for each channel but it's against specifications
		<-watchDogFeedbackOne
		<-watchDogFeedbackTwo
	}
}
