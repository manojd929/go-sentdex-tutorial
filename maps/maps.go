package maps

import "fmt"

func Maps() {
	// var grades map[string]float32
	// map is just a reference  type 
	// if it needs to have value use make
	grades := make(map[string]float32)
	grades["Timmy"] = 42
	grades["Jessy"] = 92
	grades["Sam"] = 32

	fmt.Println(grades)

	SamsGrade := grades["Timmy"]
	fmt.Println(SamsGrade)

	delete(grades, "Sam")

	fmt.Println(grades)

	for key, value := range grades {
		fmt.Println(key, value)
	}

	// what we wanted more than float32 here
	// what if we want multiple values over there
	// structs can be used
}