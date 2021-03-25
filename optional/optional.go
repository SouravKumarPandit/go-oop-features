package main

import (
	"oopfeatures/souravkumarpandit/optional/server"
	"time"
)

func main() {
	svr := server.New(
		server.WithHost("localhost"),
		server.WithMaxConn(100),
		server.WithPort(8080),
		server.WithTimeout(time.Minute),
	)
	svr.Start()
	time.Sleep(time.Second)
	svr.Stop()
}
