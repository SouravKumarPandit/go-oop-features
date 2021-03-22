package bpack

import (
	"fmt"
	"oopfeatures/souravkumarpandit/protectedpackage/apackage/bpackage/internal/apack"
)

func Bar() {
	fmt.Println("Hello from Bar")
	apack.Bar()
}