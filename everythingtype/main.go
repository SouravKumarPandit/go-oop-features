package main

import "time"

type Option func(*Server) //function is a type
type Helper interface{}   // interface is type
type SomeStruct struct{}  // Struct is a type
type TakingParam interface{}

type (
	MultiTypeInterface interface{}
	OtherHelper        struct{}
	Primitive          int64
	SomeFunction       func(*TakingParam)
)

type Server struct {
	host    string
	port    int
	timeout time.Duration
	maxConn int
}

func main() {

}
