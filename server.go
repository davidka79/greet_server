package main

import (
	"fmt"
	"log"
	"net"

	"github.com/davidka79/greetpb"

	"google.golang.org/grpc"
)

type server struct {
}

func (s server) mustEmbedUnimplementedGteetServiceServer() {
	panic("implement me")
}

func main() {
	fmt.Println("Hello World")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("failed %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGteetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed server %v", err)
	}
}
