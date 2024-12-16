# Variables
PROTOC_VERSION = 26.1
PROTOC_DIR = /usr/local/bin
PROTO_SRC_DIR = backend/internal/api/proto
PROTO_OUT_DIR = .
PROTOC = protoc
PROTOC_INSTALLED = $(PROTOC) >/dev/null
PROTOC_GEN_GO = ${HOME}/go/bin/protoc-gen-go
FILE_PATH=backend/data/data.json

ARCH := $(shell uname -m)
ifeq ($(ARCH),arm64)
    PROTOC_ARCH := aarch_64
else
    PROTOC_ARCH := x86_64
endif

tools:  ## fetch and install all required tools
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/golang/protobuf/protoc-gen-go@latest
	go install github.com/mwitkow/go-proto-validators/protoc-gen-govalidators@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

fmt:    ## format the go source files
	go fmt ./...
	goimports -w .

# Targets
.PHONY: proto
proto:
	 @echo "Compiling protobuf files..."
	$(PROTOC) --go_out=$(PROTO_OUT_DIR) --go_opt=paths=source_relative \
	--go-grpc_out=$(PROTO_OUT_DIR) --go-grpc_opt=paths=source_relative \
	$(PROTO_SRC_DIR)/*.proto
	@echo "Done."

.PHONY: clean
clean:
	@echo "Cleaning generated files..."
	rm -rf $(PROTO_OUT_DIR)/*.go
	@echo "Done."

.PHONY: run
run:
	@echo "Running server..."
	@FILE_PATH=$(FILE_PATH) go run backend/cmd/main.go

.PHONY: test
test:
	@echo "Running tests..."
	@go test ./...
	@echo "Done."