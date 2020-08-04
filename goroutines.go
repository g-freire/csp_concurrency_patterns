//package go_csp
package main

import (
	"fmt"
	"sync"
	"time"
)

/*
GOROUTINES SYNCRONIZATION
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

/*
Next we use waitgroups primitives, since it is more cost efficient than sleep (waits only the necessary time),
within a anonymous function to syncronize the main goroutine with the scheduled ones.
Is interesting to note that concurrent application does not guarantee the order of execution,
the OS manages the threads priorities.
 */

func WaitGroupExample() {
	var wait sync.WaitGroup
	goRoutines := 5
	wait.Add(goRoutines) //add one wait entity - same as +1

	for i :=0; i < goRoutines; i++{
		go func(goRoutineID int){
			fmt.Printf("ID:%d: Hello waited goroutine ID: \n", goRoutineID)
			wait.Done() //subtract one wait entity - same as -1
		}(i)
	}
	wait.Wait() // this is probably executed before the goroutines
}

// func main(){
	// NotSyncExample()
	// AnonymousExample()
	// WaitGroupExample()
// }