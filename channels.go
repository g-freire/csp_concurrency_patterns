//package go_csp
package main

import (
	"fmt"
	"sync"
	"time"
)

/*
CHANNELS

Channels are the second primitive in the language. They are the way
we communicate between process, by sending data through one end and
receiving it at the other (like a pipe). They lead to better concurrent designs
than using mutexes.

By default, channels block the execution of goroutines until something is received
or emitted through the channel ( passive/waiting emitters and listeners)
The channel's buffer is initialized with the specified buffer capacity, if zero,
or the size is omitted, the channel is unbuffered.
*/

func UnbufferedChannel(){
	channel := make(chan string)
	go func(){
		channel <- "hello world!"
	}()

	// We don't need to use a WaitGroup to sync the goroutine as
	// the defaul nature of channels is to block until data is received
	message := <-channel
	fmt.Println(message)
}

// In this example, we will block the goroutine until the receiver is ready
func UnbufferedChannel2(){
	channel := make(chan string)
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	go func(){
		channel <- "Hello world! 1"
		println("Finishing blocked goroutine")
		waitGroup.Done()
	}()

	time.Sleep(3 * time.Second)
	message := <-channel
	fmt.Println(message)
	waitGroup.Wait()
}

// Now, senders don't need to wait until some goroutine picks the data they are sending
// So the goroutine buffers the string into the channel and continues, no needing to waiting
// for the receiver
func BufferedChannel(){
	channel := make(chan string, 1)
	go func(){
		channel <- "Hello world! 1"
		//channel <- "Hello World! 2"  // since it has buffer size 1, this would block the execution
		println("Finishing unblocked goroutine")
	}()

	time.Sleep(3 * time.Second)
	message := <-channel
	fmt.Println(message)

}


func main(){
	//UnbufferedChannel()
	UnbufferedChannel2()
	//BufferedChannel()
}