//package go_csp
package main

import (
"fmt"
"sync"
"time"
)

/*
MUTEXES
Mutex(Mutual Exclusion) is a control mecanism to avoid race condition ( two or more goroutines trying to aces
some memory location).
Go has a very handy tool to help diagnose race conditions
$ go run -race main.go
Comment line 36-37 and check the race detector
The detector works at runtime, so we can have potential race condition that the detector
will not detect, for example one goroutine vs many goroutines, and bad designs that is not
immediately executing race conditions
*/

type Counter struct{
	sync.Mutex
	value int
}

func MutexCounterExample(){
	counter := Counter{}

	for i :=0; i<10; i++{
		go func(i int){
		counter.Lock()
		counter.value++
		defer counter.Unlock()
		}(i)
	}

	time.Sleep(time.Second)
	counter.Lock()
	defer counter.Unlock()
	fmt.Printf("Total counter value %d", counter.value)
}


func main(){
	MutexCounterExample()
}