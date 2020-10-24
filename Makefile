BUIlD_BASE_DIR := build
GEN_BASE_DIR := gen
PROTO_BASE_DIR := proto

all: examples

.PHONY: examples
examples:
	mkdir -p $(GEN_BASE_DIR)/bulkbeacon/v1 $(BUIlD_BASE_DIR)
	protoc -I=$(PROTO_BASE_DIR) $(PROTO_BASE_DIR)/bulkbeacon/v1/bulkbeacon.proto --go_out=plugins=grpc:$(GEN_BASE_DIR)
	go build -o $(BUIlD_BASE_DIR)/grpc_example_client cmd/example_client/grpc_example_client.go

.PHONY: clean
clean:
	rm -rf $(GEN_BASE_DIR) $(BUIlD_BASE_DIR)

