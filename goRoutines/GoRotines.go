package goRoutines

import (
	"time"
	"fmt"
)

// concurrency vs Parallelism
// in go the way we do concurrency is go routines

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}
}

func GoRoutines() {
	// say("hey") prints hey hey hey there there there
	go say("hey") // prints hey there hey there hey there
	say("there")
}