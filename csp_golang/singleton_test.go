//package go_csp
package main

import (
	"fmt"
	"testing"
	"time"
)

// SINGLETON PATTERN USING CHANNELS AND MUTEXES

// go test -v singleton_test.go

func TestSingletonUnit(t *testing.T) {
	singleton := GetInstance()
	singleton2 := GetInstance()
	n := 5000
	for i := 0; i < n; i++ { // this loop is scheduling the execution of the goroutines
		go singleton.AddOne()
		go singleton2.AddOne()
	}
	fmt.Printf("Before loop, current count is %d\n", singleton.GetCount())
	var val int
	for val != n*2 {
		val = singleton.GetCount()
		time.Sleep(10 * time.Millisecond)
	}
	singleton.Stop()
}

type singleton struct {}

var instance singleton

func GetInstance() *singleton {
	return &instance
}

func (s *singleton) AddOne() {
	addCh <- true
}

func (s *singleton) GetCount() int {
	resCh := make(chan int)
	defer close(resCh)
	getCountCh <- resCh
	return <-resCh
}

func (s *singleton) Stop() {
	quitCh <- true
	close(addCh)
	close(getCountCh)
	close(quitCh)
}

var addCh chan bool = make(chan bool)
var getCountCh chan chan int = make(chan chan int)
var quitCh chan bool = make(chan bool)

// the init() function in any package will get executed on program execution

func init() {
	var count int
	go func(addCh <-chan bool, getCountCh <-chan chan int, quitCh <-chan
	bool) {
		for {
			select {
			case <-addCh:
				count++
			case ch := <-getCountCh:
				ch <- count
			case <-quitCh:
				return
			}
		}
	}(addCh, getCountCh, quitCh)
}

/*
EQUIVALENT SINGLETON DATA STRUCTURE USING MUTEX
RWMutex stands for READ and WRITE (eg. don't need wait if n are reading the same address)
*/

/*
type singleton struct {
	count int
	sync.RWMutex
}

var instance singleton
func GetInstance() *singleton {
	return &instance
}

func (s *singleton) AddOne() {
	s.Lock()
	defer s.Unlock()
	s.count++
}

func (s *singleton) GetCount()int {
	s.RLock()
	defer s.RUnlock()
	return s.count
}
*/