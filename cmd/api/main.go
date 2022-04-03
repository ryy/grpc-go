package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"user/pb"
	"user/service"
)

func main() {
	port := 50051
	listenPort, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterRockPaperScissorsServiceServer(server, service.NewRockPaperScissorsService())

	reflection.Register(server)

	log.Printf("Server is running!")
	server.Serve(listenPort)
}