package main

import (
	"github.com/server_prototype/service"
	"github.com/smallnest/rpcx/server"
)

func main() {
	// start server
	s := server.NewServer()
	s.RegisterName("Arith", new(service.Arith), "")
	s.Serve("tcp", ":8972")

	// register service
	s.Register(new(service.Arith), "")
}
