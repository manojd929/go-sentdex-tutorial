package goRoutineSynchronization

import (
	"time"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done()
}

func GoRoutineSynchronization() {
	/*
	go say("hey")
	go say("there")
	time.Sleep(time.Second)
	*/

	wg.Add(1)
	go say("hey")
	wg.Add(1)
	go say("There")
	wg.Wait()

	wg.Add(1)
	go say("Hi")
	wg.Wait()

}