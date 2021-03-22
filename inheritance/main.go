package main

import "fmt"

/*Go has composition*/
type Vehicle struct {
	brand string
	wheels int
}

type Car struct {
	Vehicle
}
type MotorCycle struct {
	Base Vehicle
}


func main() {
	vehicle :=Car{Vehicle{"BMW",2}}
	fmt.Println(vehicle.brand)
	fmt.Println(vehicle.wheels)

}
