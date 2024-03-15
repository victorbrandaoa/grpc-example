package main

import (
	"net"
	"server/cmd/svc"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"server/proto/user"
)

var (
	listener net.Listener
	server   *grpc.Server
	logger   *zap.Logger
)

// func serialize() {
// 	myself := &user.User{
// 		Name:      "Victor Brandão",
// 		Password:  "123456",
// 		Birthdate: &timestamppb.Timestamp{Nanos: int32(time.Date(2000, 4, 14, 0, 0, 0, 0, time.UTC).Unix())},
// 		Account:   user.Account_PERSONAL,
// 	}

// 	out, err := proto.Marshal(myself)
// 	if err != nil {
// 		fmt.Println("Erro ao encodar o objeto")
// 		return
// 	}

// 	err = os.WriteFile("1.txt", out, 0777)

// 	if err != nil {
// 		fmt.Println("Erro ao escrever o arquivo")
// 		return
// 	}
// }

// func deserialize() {
// 	in, err := os.ReadFile("1.txt")
// 	if err != nil {
// 		fmt.Println("Erro ao tentar ler o arquivo")
// 		return
// 	}

// 	myself := &user.User{}

// 	err = proto.Unmarshal(in, myself)
// 	if err != nil {
// 		fmt.Println("Erro o fazer o parse")
// 	}

// 	fmt.Println(myself)
// }

func initListener() {
	var err error
	addr := "localhost:50000"

	listener, err = net.Listen("tcp", addr)
	if err != nil {
		logger.Panic("Failed to listen")
	}

	logger.Info("Success")

	return
}

func main() {
	logger, _ = zap.NewProduction()
	defer logger.Sync()

	initListener()

	server = grpc.NewServer()
	user.RegisterUserServiceServer(server, &svc.UserServiceServer{})

	logger.Info("Starting gRPC server...")
	if err := server.Serve(listener); err != nil {
		logger.Panic("Failed to start gRPC server", zap.Error(err))
	}
}
