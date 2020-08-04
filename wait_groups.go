//package go_csp
package main

import (
	"fmt"
	"sync"
)

/*
WAIT GROUPS

Wait groups  are more cost efficient than sleep (waits only the necessary time),
Next example show goroutines within a anonymous function using waiting groups
to syncronize the main goroutine with the scheduled ones.
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
func main(){
	WaitGroupExample()
}