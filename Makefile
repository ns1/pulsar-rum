# Copyright 2020 NSONE, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http:#www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

BUIlD_BASE_DIR := build
GEN_BASE_DIR := gen
PROTO_BASE_DIR := proto

PB_DIR := proto/bulkbeacon
PB_OUT := pkg/bulkbeacon
GRPC_OUT := pkg/bulkbeacon
PKG_PREFIX := github.com/ns1/pulsar-rum
BUILD := ./build.sh

all: bulkbeacon_v2

.PHONY: bulkbeacon_v1
bulkbeacon_v1: pkg/bulkbeacon/*.pb.go
	mkdir -p $(GEN_BASE_DIR)/bulkbeacon/v1 $(BUIlD_BASE_DIR)
	protoc -I=$(PROTO_BASE_DIR) $(PROTO_BASE_DIR)/bulkbeacon/v1/bulkbeacon.proto --go_out=plugins=grpc:$(GEN_BASE_DIR)
	go build -o $(BUIlD_BASE_DIR)/grpc_example_client cmd/example_client/grpc_example_client.go

.PHONY: bulkbeacon_v2
bulkbeacon_v2: $(PB_OUT)/v2/*.go
$(PB_OUT)/v2/*.go: $(PB_DIR)/v2/*.proto
	$(BUILD) v2

.PHONY: bulkbeacon_v3
bulkbeacon_v3: $(PB_OUT)/v3/*.go
$(PB_OUT)/v3/*.go: $(PB_DIR)/v3/*.proto
	$(BUILD) v3

.PHONY: clean
clean:
	rm -rf $(GEN_BASE_DIR) $(BUIlD_BASE_DIR) $(PB_OUT)/v2 $(PB_OUT)/v3

