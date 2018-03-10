package loops

import "fmt"

func Loops() {
	// first loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// second loop
	i := 0
	for i < 10 {
		i += 5
		fmt.Println(i)
	}
	// if statement
	x:=3 
	if x < 5 {
		fmt.Println(x)
	}


}