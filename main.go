package main

import (
	repository "go-grpc/internal/Repository"
	internal_grpc "go-grpc/internal/grpc"
	internal_services "go-grpc/internal/services"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	listner, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatal("Invalid Address : ", err)
	}
	defer listner.Close()

	grpcServer := grpc.NewServer()
	svc := &internal_services.PersonService{
		Repo: repository.NewRepository(),
	}
	internal_grpc.RegisterPersonServiceServer(grpcServer, svc)
	grpcServer.Serve(listner)

}
