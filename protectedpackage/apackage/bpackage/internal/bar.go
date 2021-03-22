package somepack

import (
	"fmt"
	"oopfeatures/souravkumarpandit/protectedpackage/apackage/bpackage/internal/apack"
	"oopfeatures/souravkumarpandit/protectedpackage/apackage/bpackage/internal/bpack"
)

func Bar() {
	fmt.Println("Hello from Bar")
	apack.Bar()
	bpack.Bar()
}