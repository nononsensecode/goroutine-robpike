package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	str  string
	wait chan bool
}

func main() {
	joe := boring("Joe")
	ann := boring("Ann")
	combined := fanIn(joe, ann)

	for i := 0; i < 5; i++ {
		msg1 := <-combined
		fmt.Println(msg1.str)
		msg2 := <-combined
		fmt.Println(msg2.str)
		msg1.wait <- true
		msg2.wait <- true
	}
}

func fanIn(input1, input2 <-chan Message) <-chan Message {
	c := make(chan Message)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

func boring(msg string) <-chan Message {
	ch := make(chan Message)
	waitForIt := make(chan bool)
	go func() {
		for i := 0; ; i++ {
			ch <- Message{fmt.Sprintf("%s %d", msg, i), waitForIt}
			time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
			<-waitForIt
		}
	}()
	return ch
}
