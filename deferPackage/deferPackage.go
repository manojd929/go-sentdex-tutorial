package deferPackage

import (
	"time"
	"fmt"
	"sync"
)

/* what if say throws error
 so it will be waiting forever

 in go we have defer
 defer statement wont eun until surrounding function completes
 or panics out
 */
var wg sync.WaitGroup

func say(s string) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}
}

func DeferPackage() {
	wg.Add(1)
	go say("hey")
	wg.Add(1)
	go say("There")
	wg.Wait()
}

/* 
func foo() {
	// will run only foo completes
	defer fmt.Println("done! ")
	// what if we add another - first in last out order
	defer fmt.Println("Are we done ?")
	fmt.Println("Doing some stuff ")
}

func DeferPackage() {
	foo()
}
*/