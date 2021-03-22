package main

import "fmt"

/*
const (
	CarType = iota
	MotorCycleType
	OtherType
)*/

const (
	CarType        = 75
	MotorCycleType = 75*(iota+1)
	OtherType
)

func main() {
	/*can be used as enum*/
	fmt.Println(CarType)
	fmt.Println(MotorCycleType)
	fmt.Println(OtherType)

}
