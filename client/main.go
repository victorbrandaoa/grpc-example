package main

import (
	"client/proto/user"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}
	client := user.NewUserServiceClient(conn)

	personGet, _ := client.CreateUser(context.Background(), &user.CreateUserRequest{Name: "Victor"})

	log.Printf("person: %v", personGet)
}
