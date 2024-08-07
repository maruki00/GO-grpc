package main

import (
	"go/grpc/invoicer/go/grpc"
	"log"
	"net"
)

func main() {
	listner, err := net.Listen("tcp", "8089")
	if err != nil {
		log.Fatal("Invalid Address")
	}

	serverRigister := grpc.NewServer()
	grpc.Invoicer.RegisterInvoicerServer(serverRigister)
}
