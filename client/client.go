package main

import (
	"context"
	"fmt"
	"log"

	"github.com/server_prototype/service"
	"github.com/smallnest/rpcx/client"
)

var addr string = "192.168.1.192"

func main() {
	AsyncCall()
	SyncCall()
}

func AsyncCall() {
	d, e := client.NewPeer2PeerDiscovery("tcp@"+addr, "")
	if e != nil {
		fmt.Printf("client.NewPeer2PeerDiscovery, e = %v", e)
		return
	}
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	args := &service.Args{
		A: 10,
		B: 20,
	}

	reply := &service.Reply{}
	call, err := xclient.Go(context.Background(), "Mul", args, reply, nil)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	replyCall := <-call.Done
	if replyCall.Error != nil {
		log.Fatalf("failed to call: %v", replyCall.Error)
	} else {
		log.Printf("%d * %d = %d", args.A, args.B, reply.C)
	}
}

func SyncCall() {
	// #1
	d, e := client.NewPeer2PeerDiscovery("tcp@"+addr, "")
	if e != nil {
		fmt.Printf("client.NewPeer2PeerDiscovery, e = %v\n", e)
		return
	}
	// #2
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	// #3
	args := &service.Args{
		A: 10,
		B: 20,
	}

	// #4
	reply := &service.Reply{}

	// #5
	err := xclient.Call(context.Background(), "Mul", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Printf("%d * %d = %d", args.A, args.B, reply.C)
}
