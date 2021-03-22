package main

import "fmt"

type Car struct {
	Vehicle
}

type Vehicle struct {
	brand  string //empty string
	wheels int ///0
}
func NewFord() *Car {
	return &Car{
		Vehicle{"Ford",7},
	}
}

func main() {
	//car:=NewFord()
	//fmt.Println(car)

	car:=Car{}
	fmt.Println(car.wheels)


}
