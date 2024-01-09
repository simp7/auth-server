package main

import (
	"auth-server/service"
	"github.com/simp7/idl/gen/go/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	port      = ":50051"
	secretKey = "secret"
)

func main() {

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	serv := service.NewServer(nil, nil)

	auth.RegisterAuthServer(s, serv)
	reflection.Register(s)

	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}
	if err = s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
