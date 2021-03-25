package main

import "fmt"

const (
	CarType = 45 + iota
	MotorCycleType
	CycleType
	OtherType
)

/*
const (
	CarType        = 75
	MotorCycleType = 75*(iota+1)
	CycleType = 75*(iota)
	OtherType
)
*/
func main() {
	/*can be used as enum*/
	fmt.Println(CarType)
	fmt.Println(MotorCycleType)
	fmt.Println(CycleType)
	fmt.Println(OtherType)

}
