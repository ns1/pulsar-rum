#!/bin/bash
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

# Utility for generating protobuf and gRPC artifacts for the Bulk Beacon service
# from versions 2 and up.

valid_versions="v2 v3"
PB_DIR="proto/bulkbeacon"
PB_OUT="pkg/bulkbeacon"
PKG_PREFIX="github.com/ns1/pulsar-rum"

function validate_version() {
    for i in $valid_versions
    do
      if [[ $i == $1 ]]
      then
        echo "Building bulkbeacon_${i}"
        return
      fi
    done
    echo "Error: '$1' is not a valid Bulk Beacon version."
    exit 1
}

function build() {
  v="${1}"
  mkdir -p "${PB_OUT}/${v}"
	protoc --proto_path="${PB_DIR}/${v}" --go_out=. --go-grpc_out=. \
	--go_opt="module=${PKG_PREFIX}" --go-grpc_opt="module=${PKG_PREFIX}" \
      "${PB_DIR}/${v}/bulkbeacon_${v}.proto"
}

function usage() {
  echo -n "Utility for generating protobuf and gRPC artifacts for the Pulsar Bulk "
  echo "Beacon "
  echo "service from versions 2 and up."
  echo "(c) NS1 Inc."
  echo "Usage:"
  echo "$0 <v2|v3>"
}

if [[ $# != 1 ]]
then
  usage
  exit 1
elif [[ $1 == "-h" || $1 == "--help" ]]
then
  usage
  exit 1
fi

validate_version "${1}"
build "${1}"

