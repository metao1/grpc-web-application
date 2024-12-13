package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"

	"github.com/metao1/creativefabrica/backend/internal/api"
	aps "github.com/metao1/creativefabrica/backend/internal/api/proto"
)

const (
	port = ":50051"
)

func main() {

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	serverRegistrar := grpc.NewServer()
	go func(serverRegistrar *grpc.Server) {
		// Creates a gRPC server
		listener, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}
		log.Println("Server started on port", port)
		log.Println("Reading File path:", os.Getenv("FILE_PATH"))
		filepath := os.Getenv("FILE_PATH")
		service := &api.ActiveCreatorsConfig{FilePath: filepath}
		aps.RegisterCreatorServiceServer(serverRegistrar, service)
		service.Init()
		if err := serverRegistrar.Serve(listener); err != nil {
			log.Fatalf("cannot start server: %s", err)
		}
	}(serverRegistrar)
	<-stop
	log.Println("Server stopped")
	serverRegistrar.GracefulStop()
}
