# Determine the root of the Git repository
REPO_ROOT := $(shell git rev-parse --show-toplevel)

# Define the directories and files relative to the repository root
PROTO_DIR := $(REPO_ROOT)/src/proto
PROTO_FILES := $(wildcard $(PROTO_DIR)/*.proto)
OUT_DIR := $(REPO_ROOT)/src/
OPENAPI_OUT_DIR := $(REPO_ROOT)/src/app/src/gen
FRONTEND_API_OUT_DIR := $(REPO_ROOT)/src/app/src/gen/api

# Define the paths to the protoc plugins
PROTOC_GEN_GO := $(shell go env GOPATH)/bin/protoc-gen-go
PROTOC_GEN_GO_GRPC := $(shell go env GOPATH)/bin/protoc-gen-go-grpc
PROTOC_GEN_OPENAPI := $(shell go env GOPATH)/bin/protoc-gen-openapi

# Rule to generate Go code from .proto files
.PHONY: all
all: $(PROTO_FILES)
	protoc -I=$(PROTO_DIR) \
		--plugin=protoc-gen-go=$(PROTOC_GEN_GO) \
		--go_out=$(OUT_DIR) \
		--plugin=protoc-gen-go-grpc=$(PROTOC_GEN_GO_GRPC) \
		--go-grpc_out=$(OUT_DIR) \
		--plugin=protoc-gen-openapi=$(PROTOC_GEN_OPENAPI) \
		--openapi_out=$(OPENAPI_OUT_DIR) \
		$(PROTO_FILES)



# Clean up generated files
.PHONY: clean
clean:
	rm -rf $(OUT_DIR)/*.pb.go

