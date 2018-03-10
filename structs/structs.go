package structs

import (
	"fmt"	
)

const unsignedSixteenBitMax float64 = 65535
const mileToKilometer float64 = 1.60934

type car struct {
	gasPedal uint16
	brakePedal uint16
	steeringWheel int16
	topSpeedKmh float64
}

func (c car) kmhValueReceiverMethod() float64 {
	return float64(c.gasPedal) * (c.topSpeedKmh/unsignedSixteenBitMax)
}

func (c car) mphValueReceiverMethod() float64 {
	return float64(c.gasPedal) * (c.topSpeedKmh/unsignedSixteenBitMax/mileToKilometer)
}

func (c  *car) kmhPointerReceiverMethod() float64 {
	c.topSpeedKmh = 400.0
	return float64(c.gasPedal) * (c.topSpeedKmh/unsignedSixteenBitMax)
}

func (c *car) mphPointerReceiverMethod() float64 {
	return float64(c.gasPedal) * (c.topSpeedKmh/unsignedSixteenBitMax/mileToKilometer)
}

func Structs() {
	a_car := car { 60000, 0, 12562, 225.0 }
	// fmt.Println(a)
	fmt.Println(a_car.kmhValueReceiverMethod(), " kmh")
	fmt.Println(a_car.mphValueReceiverMethod(), " mph")

	fmt.Println(a_car.kmhPointerReceiverMethod(), " kmh")
	fmt.Println(a_car.mphValueReceiverMethod(), " mph")
}