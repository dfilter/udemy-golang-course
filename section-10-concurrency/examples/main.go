package main

import (
	"fmt"
	"time"
)

func greet(greeting string, doneChan chan bool) {
	fmt.Println("Hello,", greeting)
	doneChan <- true
}

func slowGreet(greeting string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello,", greeting)
	doneChan <- true // sending the value of true to doneChan channel
	close(doneChan)
}

func main() {
	// dones := make([]chan bool, 4)
	// dones[0] = make(chan bool)
	// dones[1] = make(chan bool)
	// dones[2] = make(chan bool)
	// dones[3] = make(chan bool)

	done := make(chan bool)
	go greet("John how are you today.", done)
	go greet("you're looking kinda generic...", done)
	go slowGreet("that took a while!", done)
	go greet("that wasn't too long was it?", done)

	// for _, done := range dones {
	// 	<-done
	// }

	// IF the same channel is used for multiple concurrent functions one has
	// to get the results the same number of times the channel is used.
	// In this case 4 times.
	// <-done // The program will wait untill some data comes out of done
	// <-done
	// <-done
	// <-done

	for range done {
	}
}
