package channels
// the way to send and receive value with go routines

import "fmt"

func foo(c chan int, someValue int) {
	c <-someValue * 5
}

func Channels() {
	fooVal := make(chan int) // fooVal is a channel of int type, we have to make a channel

	go foo(fooVal, 5)
	go foo(fooVal, 3)

	v1 := <-fooVal
	v2 := <-fooVal

	fmt.Println(v1, v2)
}