package creational

import (
	"fmt"
	"sync"
	"testing"
)

// SINGLETON PATTERN
// An unique single instance of a type in the entire program (no duplicates)
// Shared instance and restricted object creating

// SYNC SINGLETON - A COUNTER EXAMPLE

type Singleton interface {
	AddOne() int
}
type singleton struct{
	count int
}

//initializes a pointer to a struct as nil (global variable!)
var instance *singleton

func GetSingleton() *singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func (s *singleton) AddOne() int{
	s.count++
	return s.count
}

//go test -v -run=TestGetInstance .
func TestGetInstance(t *testing.T) {
	counter1 := GetSingleton()
	if counter1 == nil {
		t.Error("expected pointer to Singleton after calling GetInstance(), not nil")
	}
	currentCount := counter1.AddOne()
	if currentCount != 1 {
		t.Errorf("After calling for the first time to count, the count must be1 but it is %d\n", currentCount)
	}
	expectedCounter := counter1
	counter2 := GetSingleton()
	if counter2 != expectedCounter {
		//Test 2 failed
		t.Error("Expected same instance in counter2 but it got a different instance")
	}
	currentCount = counter2.AddOne()
	if currentCount != 2 {
		t.Errorf("After calling 'AddOne' using the second counter, the current count must be 2 but was %d\n", currentCount)
	}
}

// ASYNC - CONCURRENCY SAFE SINGLETONS EXAMPLES
// In Golang we could achieve concurrency safety singletons with:
// 1) init() functions (only applicable if the early initialization of the object is ok, and can be trouble if more packages have same init code)
// 2) sync.Once in sync package
// 3) Mutexes

// 4) Channels

var lock = &sync.Mutex{}


type singleton2 struct{
	count int
}

//initializes a pointer to a struct as nil (global variable!)
var singleInstance *singleton2

func getInstance() *singleton2 {
	//“Check-Lock-Check” approach, still not atomic - many checks
	if singleInstance == nil {
		// locking inside to avoid expensives blocks bottleneck
		lock.Lock()
		defer lock.Unlock()
		// to make sure that if more than one goroutine bypass the first check then only one goroutine will create the instance
		if singleInstance == nil {
			fmt.Println("Creting Single Instance Now")
			singleInstance = &singleton2{}
		} else {
			fmt.Println("Single Instance already created-1")
		}
	} else {
		fmt.Println("Single Instance already created-2")
	}
	return singleInstance
}

func (s *singleton2) AddOne() int{
	s.count++
	return s.count
}

//func main() {
//	for i := 0; i < 100; i++ {
//		go getInstance()
//	}
//	// Scanln is similar to Scan, but stops scanning at a newline and
//	// after the final item there must be a newline or EOF.
//	fmt.Scanln()
//}

// USING SYNC ONCE
var once sync.Once

type single struct {
}

var singleInstanceOnce *single

func getInstanceOnce() *single {
	if singleInstanceOnce == nil {
		once.Do(
			func() {
				fmt.Println("Creting Single Instance Now")
				singleInstanceOnce = &single{}
			})
	} else {
		fmt.Println("Single Instance already created-2")
	}
	return singleInstanceOnce
}