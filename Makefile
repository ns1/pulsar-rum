#
#
EXE=grpc_example_client
GO_FILES := cmd/example_client/grpc_example_client.go
PB_DIR := internal/bulkbeacon
PROTOC_GO_FILES := $(PB_DIR)/bulkbeacon.pb.go
PROTOC_FILES := $(PB_DIR)/bulkbeacon.proto

all: protoc build

.PHONY: protoc
protoc: $(PROTOC_GO_FILES)
$(PROTOC_GO_FILES): $(PROTOC_FILES)
	protoc -I $(PB_DIR) $(PROTOC_FILES) --go_out=plugins=grpc:internal/bulkbeacon

build: $(PROTOC_GO_FILES) $(GO_FILES)
	go build -o build/$(EXE) $(GO_FILES)

.PHONY: clean
clean:
	rm -rf $(PROTOC_GO_FILES) build/