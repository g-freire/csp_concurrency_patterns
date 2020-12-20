package main

import "fmt"

func main() {
	var data int
	go func(){
		// A trying to access the variable data
		data++
	}()
	// B trying to access the variable data
	if data ==0 {
		fmt.Printf("the value is %v. \n", data)
	}
}
//	RACE CONDITION FIRST EXAMPLE

//	Here, A and B are both trying to access the variable data, but there is no
//	guarantee what order this might happen in. There are three possible outcomes
//	to running this code:

//	Nothing is printed. In this case, A was executed before B.

//	“The value is 0” is printed. In this case, B were executed
//	before A.

//	“the value is 1” is printed. In this case, B was executed before A. But A was executed
//  before the print line


//CONCEPTS
// critical section: region of shared resources
// atomic: indivisible, uninterrupted
