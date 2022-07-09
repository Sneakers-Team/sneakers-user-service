package main

import (
	"log"
	"msqrd/user-service/pkg/api/userapi"
	"msqrd/user-service/pkg/store"
	"msqrd/user-service/pkg/user"
	"net"

	"google.golang.org/grpc"
)

func main() {
	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()

	cfg := store.NewConfig("sneakers.db")

	store, err := store.NewStore(cfg)
	if err != nil {
		log.Fatal(err)
	}

	srv := &user.GRPCServer{
		Store: store,
	}

	userapi.RegisterUserServiceServer(s, srv)

	log.Println("user service started")

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
