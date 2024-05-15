package main

import (
	"log"
	"net"
	"stakeholders-service/server"
	"stakeholders-service/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func startServer(JwtService *service.JWTService) {
	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	server.RegisterStakeholdersMicroserviceServer(grpcServer, &server.StakeholdersMicroservice{
		JwtService: JwtService,
	})

	listener, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("gRPC server listening on port :8082")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}

func main() {
	jwtService := &service.JWTService{}

	startServer(jwtService)
}
