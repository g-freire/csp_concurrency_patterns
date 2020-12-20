package main

import (
	"fmt"
	"time"
)

func produce(p Producer) {
	var i int64 = 0
	for {
		msg := Msg{Id: p.Id, Counter: i}
		p.Channel <- msg
		i++
		time.Sleep(p.Delay)
	}
}

func read(out chan Msg) {
	for x := range out {
		fmt.Println(" GOROUTINE ID: ", x.Id,
					    " CURRENT COUNT:", x.Counter)
	}
}

type Msg struct {
	Id      string
	Counter int64
}

type Producer struct {
	Id      string
	Channel chan Msg
	Delay   time.Duration
}

func main() {
	inputChannel := make(chan Msg)
	p1 := Producer{Id: "A", Delay: 100 * time.Millisecond, Channel: inputChannel}
	p2 := Producer{Id: "B", Delay: 250 * time.Millisecond, Channel: inputChannel}
	go produce(p1)
	go produce(p2)

	singleConsumerCh := make(chan Msg)
	go read(singleConsumerCh)

	for i := range inputChannel {
		singleConsumerCh <- i
	}
}
