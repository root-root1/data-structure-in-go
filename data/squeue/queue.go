package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	messagePassStart = iota
	messageTicketStart
	messagePassEnd
	messageTicketEnd
)

type Queue struct {
	waitPass    int
	waitTicket  int
	playPass    bool
	playTicket  bool
	queuePass   chan int
	queueTicket chan int
	message     chan int
}

func (queue *Queue) New() {
	queue.message = make(chan int)
	queue.queuePass = make(chan int)
	queue.queueTicket = make(chan int)

	go func() {
		var message int
		for {
			select {
			case message = <-queue.message:
				switch message {
				case messagePassStart:
					queue.waitPass += 1
				case messagePassEnd:
					queue.playPass = false

				case messageTicketStart:
					queue.waitTicket += 1
				case messageTicketEnd:
					queue.playTicket = false
				}
				if queue.waitPass > 0 && queue.waitTicket > 0 && !queue.playPass && !queue.playTicket {
					queue.playPass = true
					queue.playTicket = true
					queue.waitTicket -= 1
					queue.waitPass -= 1
					queue.queuePass <- 1
					queue.queueTicket <- 1
				}
			}
		}
	}()
}

func (queue *Queue) StartTicketIssue() {
	queue.message <- messageTicketStart
	<-queue.queueTicket
}

func (queue *Queue) EndTicketIssue() {
	queue.message <- messageTicketEnd
}

func TicketIssue(queue *Queue) {
	for {
		time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)
		queue.StartTicketIssue()
		fmt.Println("Ticket Issue Starts")
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		fmt.Println("Ticket Issue Ends")
		queue.EndTicketIssue()
	}
}

func (queue *Queue) StartPass() {
	queue.message <- messagePassStart
	<-queue.queuePass
}

func (queue *Queue) EndPass() {
	queue.message <- messagePassEnd
}

func passenger(queue *Queue) {
	for {
		time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)
		queue.StartPass()
		fmt.Println("Passenger Starts")
		time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)
		fmt.Println("Passenger Ends")
		queue.EndPass()
	}
}

func main() {
	var queue *Queue = &Queue{}
	queue.New()
	fmt.Println(queue)
	for i := 0; i < 10; i += 1 {
		go passenger(queue)
	}
	for j := 0; j < 5; j += 1 {
		go TicketIssue(queue)
	}
	select {}
}
