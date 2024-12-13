## Description

This project is a backend application that provides an gRPC API backend for managing products. 
There is a service file definition for retrieving top creators for products.

The project uses Go as the programming language and includes the following packages:

## Project Structure

The project structure is as follows:

- `backend/`: Contains the backend code for project.
  -`data/`: Contains the data files used in  project.
  - `cmd/`: Contains the main entry point for  application.
  - `internal/`: Contains the internal packages of project.
    - `api/`: Contains the API-related code.
        - `proto/`: contains protobuf files for gRPC API definition            
    - `file/`: Contains the file related utilities

## Dependencies

The project uses the following external libraries:

- Go: The programming language used for the backend implementation.
- Google Protocol Buffers: Used for defining and serializing structured data.
- gRPC: Used for building high-performance, scalable, and distributed systems.

## Setup

To set up the project, follow these steps:

1. Install Go: Make sure you have Go installed on your machine. You can download it from the official Go website.

2. run the following command to run the test
```
        go test ./...
```

3. run the following command to run the gRPC API gateway

```
        FILE_PATH=backend/data/data.json go run backend/cmd/main.go
```

##### (Optional)
A service grpc go files are already generated out of `service.proto` protobuf
If any thing is made to this file, we need to recreate them using below command.

```
        make tools
        make proto
```

Sample output:

```
Compiling protobuf files...
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=.  --go-grpc_opt=paths=source_relative backend/internal/api/proto/*.proto
Done.
```

The position of the json sample file is given as environment variable named `FILE_PATH`

The server as a gRPC app, starts on port `50051` TCP.