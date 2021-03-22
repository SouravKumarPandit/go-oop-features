package main

import "fmt"

type Vehicle struct {
	brand  string
	wheels int
}

func (v *Vehicle) ComputeSpeed() int {

	return v.wheels * 5
}

type Car struct {
	Vehicle
}
type MotorCycle struct {
	Base Vehicle
}

type Speeder interface {
	GetSpeed() int
}

func (c *Car) GetSpeed() int {
	return c.wheels * 3
}

func (c *MotorCycle) GetSpeed() int {
	return c.Base.wheels * 5
}

func main() {
	vehicle := Car{Vehicle{"BeCon", 5}}
	vehicle1 := MotorCycle{Base: Vehicle{"BeCon", 5}}
	fmt.Println(vehicle.GetSpeed())
	fmt.Println(vehicle1.GetSpeed())
}
