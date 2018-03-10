package panicAndRecover

import (
	"time"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func cleanup() {
	defer wg.Done() // makes more sense to be here even without defer but at the end of func
	if r := recover(); r != nil {
		fmt.Println("Recovered in cleanup", r)
	}
}

func say(s string) {
	defer cleanup()
	for i := 0; i < 3; i++ {
		time.Sleep(time.Millisecond * 100)
		fmt.Println(s)
		if i == 2 {
			panic("Oh dear, i 2")
		}
	}
}

func PanicAndRecover() {
	wg.Add(1)
	go say("hey")
	wg.Add(1)
	go say("There")
	wg.Wait()
}