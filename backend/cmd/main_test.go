package main

import (
	"context"
	"os"
	"syscall"
	"testing"

	proto "github.com/metao1/creativefabrica/backend/internal/api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestMain_Success(t *testing.T) {
	tests := []struct {
		name                   string
		filePath               string
		expectedActiveCreators []string
	}{
		{
			name:                   "equal activities for all creators",
			filePath:               "../data/data.json",
			expectedActiveCreators: []string{"NfDynPx@oCsT5hsulDwM.com"},
		},
		{
			name:                   "equal activities for all creators",
			filePath:               "../data/data_unequivalent.json",
			expectedActiveCreators: []string{"mehrdad@gmail.com"},
		},
	}

	// Create a wait group to synchronize the server start and stop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("FILE_PATH", tt.filePath)
			go main()
			runTest(t, tt.expectedActiveCreators)
			// Stop the server
			stopServer()
		})
	}

}

func runTest(t *testing.T, expectedActiveCreators []string) {
	// Create a gRPC client
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()
	apiClient := proto.NewCreatorServiceClient(conn)

	// Send a gRPC request to the server
	request := &proto.TopActiveCreatorsRequest{Limit: 3}
	response, err := apiClient.GetTopActiveCreators(context.Background(), request)
	if err != nil {
		t.Fatalf("Failed to get top active creators: %v", err)
	}

	// Check the response is as expected
	if len(response.Emails) != len(expectedActiveCreators) {
		t.Errorf("Expected %d active creators, got %d", len(expectedActiveCreators), len(response.Emails))
	}
	for i, email := range response.Emails {
		if email != expectedActiveCreators[i] {
			t.Errorf("Expected active creator %s, got %s", expectedActiveCreators[i], email)
		}
	}
}

func stopServer() {
	// Send a SIGINT signal to the server
	syscall.Kill(syscall.Getpid(), syscall.SIGINT)
}
