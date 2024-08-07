package main

import (
	"context"
	"go/grpc/invoicer"
	"log"
	"net"

	"google.golang.org/grpc"
)

type MyInvoice struct {
	invoicer.UnimplementedInvoicerServer
}

func (o *MyInvoice) Create(context context.Context, _ *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	return &invoicer.CreateResponse{
		Pdf:  []byte("test"),
		Docx: []byte("test"),
	}, nil
}

func main() {
	listner, err := net.Listen("tcp", "8089")
	if err != nil {
		log.Fatal("Invalid Address")
	}

	serverRigister := grpc.NewServer()
	service := &MyInvoice{}
	invoicer.RegisterInvoicerServer(serverRigister, service)

	serverRigister.Serve(listner)
}
