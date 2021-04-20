# Copyright 2021 NSONE, Inc.
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

BUILD_BASE_DIR := build
PROTO_BASE_DIR := proto

PB_DIR := proto/bulkbeacon
PB_OUT := pkg/bulkbeacon
GRPC_OUT := pkg/bulkbeacon
PKG_PREFIX := github.com/ns1/pulsar-rum
BUILD := ./build.sh

all: bulkbeacon_v1 bulkbeacon_v2

V1_GO_FILES := $(PB_OUT)/v1/bulkbeacon.pb.go $(PB_OUT)/v1/bulkbeacon_grpc.pb.go
.PHONY: bulkbeacon_v1
bulkbeacon_v1: $(V1_GO_FILES)
$(V1_GO_FILES): $(PB_DIR)/v1/bulkbeacon.proto
	$(BUILD) v1
	mkdir -p $(BUILD_BASE_DIR)
	go build -o $(BUILD_BASE_DIR)/grpc_example_client_v1 cmd/example_client_v1/main.go

V2_GO_FILES := $(PB_OUT)/v2/bulkbeacon.pb.go $(PB_OUT)/v2/bulkbeacon_grpc.pb.go
.PHONY: bulkbeacon_v2
bulkbeacon_v2: $(V2_GO_FILES)
$(V2_GO_FILES): $(PB_DIR)/v2/bulkbeacon.proto
	$(BUILD) v2
	mkdir -p $(BUILD_BASE_DIR)
	go build -o $(BUILD_BASE_DIR)/grpc_example_client_v2 cmd/example_client_v2/main.go

.PHONY: clean
clean:
	rm -rf $(BUILD_BASE_DIR) $(PB_OUT)/v2 $(PB_OUT)/v1

