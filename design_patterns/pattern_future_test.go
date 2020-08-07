//package go_csp
package main

import (
	"sync"
	"testing"
	"time"
)
/* FUTURE PATTERN (PROMISE)

We will define each possible behavior of an action before executing them in
different Goroutines. The idea here is to achieve a fire-and-forget that handles all
possible results in an action.
This is a kind of lazy programming, where a Future could be calling to itself indefinitely or
just until some rule is satisfied. The idea is to define the behavior in advance and let the
future resolve the possible solutions.
*/

func TestStringOrError_Execute(t *testing.T) {
	future := &MaybeString{}

	// SUCESS - FAIL -  EXECUTE FUTURE
	t.Run("Success result", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(1)
		go timeout(t, &wg)
		future.Success(func(s string) {
			t.Log(s)
			wg.Done()
		}).Fail(func(e error) {
			t.Fail()
			wg.Done()
		})

		future.Execute(func() (string, error) {
			return "Hello World!", nil
		})
		wg.Wait()
	})

	//t.Run("Failed result", func(t *testing.T) {
	//	var wg sync.WaitGroup
	//	wg.Add(1)
	//	future.Success(func(s string) {
	//		t.Fail()
	//		wg.Done()
	//	}).Fail(func(e error) {
	//		t.Log(e.Error())
	//		wg.Done()
	//	})
	//	future.Execute(func() (string, error) {
	//		return "", errors.New("Error ocurred")
	//	})
	//	wg.Wait()
	//})
}

type SuccessFunc func(string)
type FailFunc func(error)
type ExecuteStringFunc func() (string, error)

type MaybeString struct {
	successFunc SuccessFunc
	failFunc FailFunc
}

func (s *MaybeString) Success(f SuccessFunc) *MaybeString {
	s.successFunc = f
	return s
}

func (s *MaybeString) Fail(f FailFunc) *MaybeString {
	s.failFunc = f
	return s
}

func (s *MaybeString) Execute(f ExecuteStringFunc) *MaybeString {
	return nil
}


// util function to prevent deadlock exception - aka waiting forever for the not implemented callback
func timeout(t *testing.T, wg *sync.WaitGroup) {
	time.Sleep(time.Second)
	t.Log("Timeout!")
	t.Fail()
	wg.Done()
}

