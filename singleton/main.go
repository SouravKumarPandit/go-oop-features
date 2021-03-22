package main

import (
	"fmt"
	"oopfeatures/souravkumarpandit/singleton/globalvariable1"
	"oopfeatures/souravkumarpandit/singleton/globalvariable2"
	"oopfeatures/souravkumarpandit/singleton/globalvariable3"
)
//https://medium.com/golang-issue/how-singleton-pattern-works-with-golang-2fdd61cd5a7f
func main() {
	obj1 := globalvariable1.NewClass()
	fmt.Printf("memory Address  %p \n",obj1)

	obj2 := globalvariable1.NewClass()
	fmt.Printf("memory Address  %p \n",obj2)

	obj3 := globalvariable2.NewClass()
	fmt.Printf("memory Address  %p \n",obj3)

	obj4 := globalvariable2.NewClass()
	fmt.Printf("memory Address  %p \n",obj4)


	obj5 := globalvariable3.NewClass()
	fmt.Printf("memory Address  %p \n",obj5)

	obj6 := globalvariable3.NewClass()
	fmt.Printf("memory Address  %p \n",obj6)


}
