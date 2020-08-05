//package go_csp
package main

import (
	"time"
)

/*
SELECT STATEMENT

Used to handle more than one channel input within a goroutine.

*/


func sendString(ch chan<- string, s string) {
	ch <- s
}

func receiver(helloCh <-chan string, goodbyeCh <-chan string, quitCh chan<- bool) {
	for {
		select {
		case msg := <-helloCh:
			println(msg)
		case msg := <-goodbyeCh:
			println(msg)
		case <-time.After(time.Second * 2):
			println("Nothing received in 2 seconds. Exiting")
			quitCh <- true
			break
		}
	}
}

func SelectChannel() {
	helloCh := make(chan string, 1)
	goodbyeCh := make(chan string, 1)
	quitCh := make(chan bool)

	go receiver(helloCh, goodbyeCh, quitCh) // handle many incoming channels in the same goroutine
	go sendString(helloCh, "hello!")
	time.Sleep(time.Second)
	go sendString(goodbyeCh, "goodbye!")
	time.Sleep(time.Second)
	<-quitCh
}

func RangeOverChannel(){
	ch := make(chan int)
	go func() {
		ch <- 1
		time.Sleep(time.Second)
		ch <- 2
		close(ch)
	}()
	for v := range ch {
		println(v)
	}
}


func main(){
	//SelectChannel()
	RangeOverChannel()
}