package main

import (
	"fmt"
	"grpc-example/proto/user"
	"os"

	"time"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func serialize() {
	myself := &user.User{
		Name:      "Victor Brandão",
		Password:  "123456",
		Birthdate: &timestamppb.Timestamp{Nanos: int32(time.Date(2000, 4, 14, 0, 0, 0, 0, time.UTC).Unix())},
		Account:   user.Account_PERSONAL,
	}

	out, err := proto.Marshal(myself)
	if err != nil {
		fmt.Println("Erro ao encodar o objeto")
		return
	}

	err = os.WriteFile("1.txt", out, 0777)

	if err != nil {
		fmt.Println("Erro ao escrever o arquivo")
		return
	}
}

func deserialize() {
	in, err := os.ReadFile("1.txt")
	if err != nil {
		fmt.Println("Erro ao tentar ler o arquivo")
		return
	}

	myself := &user.User{}

	err = proto.Unmarshal(in, myself)
	if err != nil {
		fmt.Println("Erro o fazer o parse")
	}

	fmt.Println(myself)
}

func main() {
	serialize()
	deserialize()
}
