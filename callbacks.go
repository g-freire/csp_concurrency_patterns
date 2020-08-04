//package go_csp
package main

import (
	"fmt"
	"strings"
	"sync"
)

/*
CALLBACKS

A callback is an anonymous function that will be executed
within the context of a different function. They are good for async models
but can easy turn into complexity, aka, callback hell.
*/

var wait sync.WaitGroup

func toUpperAsync(word string, f func(string)){
	go func() {
		f(strings.ToUpper(word))
	}()
}

func CallbackExample(){
	wait.Add(1)
	toUpperAsync("Hello callback",
					func(word string){
						fmt.Printf("Callback func result: %s \n", word)
						wait.Done()
						})
	println("waiting async response ...")
	wait.Wait()
}


func main(){
	CallbackExample()
}
