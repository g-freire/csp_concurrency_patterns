//package go_csp
package main

import (
	"time"
)

/*
GOROUTINES

Goroutines primitives are cheap lightweight unity of works managed by the Go runtime.
They run in the same address space, so access to shared memory must be synchronized.
When we use goroutines the main function has to wait for the scheduled goroutines.
Below, we demonstrate that the main func finishes execution before executing the
scheduled goroutine and then an example using a sleep approach, giving some deterministic time for the scheduled execution
 */

func NotSyncExample(){
	go func(msg string){
		println(msg)
	}("Calling the lambda without sync \n")
}

func SleepExample(){
	go func(msg string){
		println(msg)
	}("Calling the lambda with some delay \n")
	time.Sleep(time.Second)
}


// func main(){
	// NotSyncExample()
	// AnonymousExample()
// }


//////////////////////////////////////////////
//  BEST PRACTICES - GO ROUTINES
//////////////////////////////////////////////
//  DONT CREATE GOROUTINES IN YOUR LIBS -LET CONSUMER CONTROL CONCURRENCY
//  KNOW HOW IT WILL END - AVOIS MEMORY LEAKS ( ZOMBIE LISTENERS)
//  CHECK FOR RACE CONDITIONS AT COMPILE TIME 

//////////////////////////////////////////////
//  SUMMARY
//////////////////////////////////////////////
//  go keyword before function call creates a goroutine
//  when using anonymou functions, pass data as local variable(by local), avoiding closures race conditions
//  we need synchronization to deal with shared memory problems
//  we can use sync.WaitGroup to group the goroutines and wait the completion
//  use GOMAXPROCS to manage the os threads to available cores