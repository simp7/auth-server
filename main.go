package main

import (
	"auth-server/service"
	"auth-server/tokenManager/jwtManager"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
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
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatalf("error when generating random key: %v", err)
	}

	m, err := jwtManager.ECDSA(key)
	if err != nil {
		log.Fatalf("error when starting manager")
	}

	s := grpc.NewServer(opts...)
	serv := service.NewServer(nil, m)

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
