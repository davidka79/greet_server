package main

import (
	"fmt"
	"log"
	"net"

	"github.com/davidka79/go-grp-study/greetpb"

	"google.golang.org/grpc"
)

type server struct {
}

func main() {
	fmt.Println("Hello World")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("failed %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, srv)
}
