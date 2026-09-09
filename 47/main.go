package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"

	"github.com/andyfusniak/base58"
)

const (
	NUM_SENDERS = 5
	NUM_REVCS   = 2
	SEND_LOOP   = 1000
)

type token struct {
	id         int64
	name       string
	age        int
	isVerified bool
}

var (
	tokenID      int64
	tokenChannel = make(chan token, 50)
)

func beginSend(wg *sync.WaitGroup) {
	defer wg.Done()

	var sendWg sync.WaitGroup
	sendWg.Add(NUM_SENDERS)

	for id := 0; id < NUM_SENDERS; id++ {
		go func(senderID int) {
			defer sendWg.Done()
			for i := 0; i < SEND_LOOP; i++ {
				t := token{
					id:  atomic.AddInt64(&tokenID, 1),
					age: rand.Int()%100 + 5,
				}
				t.isVerified = (t.age%2 == 0)
				t.name, _ = base58.RandString(8)

				fmt.Printf("[SEND: %d : %d] Name: %s | Age: %d | Verified: %t\n",
					senderID, t.id, t.name, t.age, t.isVerified)

				tokenChannel <- t
			}
		}(id)
	}

	sendWg.Wait()
	close(tokenChannel)
}

func beginRecv(wg *sync.WaitGroup) {
	defer wg.Done()

	var recvWg sync.WaitGroup
	recvWg.Add(NUM_REVCS)

	for id := 0; id < NUM_REVCS; id++ {
		go func(receiverID int) {
			defer recvWg.Done()

			for t := range tokenChannel {
				fmt.Printf("[RECV: %d : %d] Name: %s | Age: %d | Verified: %t\n",
					receiverID, t.id, t.name, t.age, t.isVerified)
			}
		}(id)
	}

	recvWg.Wait()
}

func main() {
	var mainWg sync.WaitGroup

	mainWg.Add(2)
	go beginSend(&mainWg)
	go beginRecv(&mainWg)

	mainWg.Wait()

	fmt.Println("--------------Done--------------")
}
