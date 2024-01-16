package main

import (
	"auth-server/service"
	"auth-server/storage/inMemory"
	"auth-server/storage/postgresql"
	"auth-server/tokenManager/jwtManager"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"github.com/simp7/idl/gen/go/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	_ "github.com/lib/pq"
)

const (
	port = ":50051"
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
	userStorage, err := postgresql.Storage(postgresql.DBInfo{
		Host:     "localhost",
		Port:     5432,
		User:     "gopher",
		Password: "pass1234",
		Database: "auth",
	})
	tokenStorage := inMemory.Storage()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = userStorage.Close(); err != nil {
			log.Fatalf("error when closing userStorage: %v", err)
		}
	}()
	serv := service.NewServer(userStorage, tokenStorage, m)

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
