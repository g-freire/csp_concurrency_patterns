// package barrier

package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
	"io/ioutil"
	"net/http"
	"time"
)

/*
BARRIER PATTERN
Objective: put up a barrier so that nobody passes until we have
all the results we need.
Uses:  1) Compose the value of a type with the data coming from one or more Goroutines
	   2) Control the correctness of any of those incoming data pipes so that no inconsistent
       data is returned. We don't want a partially filled result because one
       of the pipes has returned an error.

//UNIT TESTS
$	go test -run=TestBarrier/Correct_endpoints -v pattern_barrier_test.go
$ 	go test -run=TestBarrier/One_endpoint_incorrect -v pattern_barrier_test.go
$ 	go test -run=TestBarrier/Very_short_timeout -v pattern_barrier_test.go
*/


func TestBarrier(t *testing.T) {
	t.Run("Correct endpoints", func(t *testing.T) {
		endpoints := []string{
			"http://httpbin.org/headers",
			"http://httpbin.org/user-agent",
		}
		result := captureBarrierOutput(endpoints...)

		if  !strings.Contains(result, "Accept-Encoding") ||
			!strings.Contains(result, "User-Agent") ||
			!strings.Contains(result, "user-agent") {

			t.Fail()
		}
		t.Log(result)
	})
	t.Run("One endpoint incorrect", func(t *testing.T) {
		endpoints := []string{"http://malformed-url",
			"http://httpbin.org/User-Agent"}
		result := captureBarrierOutput(endpoints...)

		if !strings.Contains(result, "ERROR") {
			t.Fail()
		}
		t.Log(result)
	})
	t.Run("Very short timeout", func(t *testing.T) {
		endpoints := []string{"http://httpbin.org/headers",
			"http://httpbin.org/User-Agent"}
		timeoutMilliseconds = 1
		result := captureBarrierOutput(endpoints...)
		if !strings.Contains(result, "Timeout") {
			t.Fail()
		}
		t.Log(result)
	})
}

func barrier(endpoints ...string) {
	requestNumber := len(endpoints)
	in := make(chan barrierResp, requestNumber)
	defer close(in)
	responses := make([]barrierResp, requestNumber)
	for _, endpoint := range endpoints {
		go makeRequest(in, endpoint)
	}
	var hasError bool
	for i := 0; i < requestNumber; i++ {
		resp := <-in
		if resp.Err != nil {
			fmt.Println("ERROR: ", resp.Err)
			hasError = true
		}
		responses[i] = resp
	}
	if !hasError {
		for _, resp := range responses {
			fmt.Println(resp.Resp)
		}
	}
}

func makeRequest(out chan<- barrierResp, url string) {
	res := barrierResp{}
	client := http.Client{
		Timeout: time.Duration(time.Duration(timeoutMilliseconds) *
			time.Millisecond),
	}
	resp, err := client.Get(url)
	if err != nil {
		res.Err = err
		out <- res
		return
	}
	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		res.Err = err
		out <- res
		return
	}
	res.Resp = string(byt)
	out <- res
}

func captureBarrierOutput(endpoints ...string) string {
	reader, writer, _ := os.Pipe()
	os.Stdout = writer
	out := make(chan string)

	// copies reader input to a byte buffer
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, reader)
		out <- buf.String()
	}()

	barrier(endpoints...)
	writer.Close()
	temp := <-out
	return temp
}

var timeoutMilliseconds int = 5000
type barrierResp struct {
	Err error
	Resp string
}

//func main(){
//	WaitGroupExample()
//}