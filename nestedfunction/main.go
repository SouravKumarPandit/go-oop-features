package main

import (
	"fmt"
	"oopfeatures/souravkumarpandit/nestedfunction/messanger"
)

func main() {
messanger.NumberValidator(func(msg string) {
	fmt.Println(msg)
})
}

