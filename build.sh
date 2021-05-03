#!/bin/bash
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

# Utility for generating protobuf and gRPC artifacts for the Bulk Beacon service.

valid_versions="v1 v2"
PB_DIR="proto/bulkbeacon"
PB_OUT="pkg/bulkbeacon"
PKG_PREFIX="github.com/ns1/pulsar-rum"

function validate_version() {
  for i in $valid_versions; do
    if [[ $i == $1 ]]; then
      echo "Bulk Beacon: ${i} is a valid version."
      return
    fi
  done
  echo "Error: '$1' is not a valid Bulk Beacon version."
  exit 1
}

function build() {
  v="${1}"
  if [ "${1}" == "v1" ]
  then
    proto_file="${PB_DIR}/bulkbeacon.proto"
    proto_pkg="${PB_OUT}"
  else
    # Versions 2 and up follow this convention.
    proto_file="${PB_DIR}/${v}/bulkbeacon.proto"
    proto_pkg="${PB_OUT}/${v}"
  fi
  opt_m="M${proto_file}=${proto_pkg}"
  echo "Bulk Beacon: Building ${proto_file}"
  mkdir -p "${proto_pkg}"
  protoc --go_out=. --go_opt="${opt_m}" --go-grpc_out=. --go-grpc_opt="${opt_m}" "${proto_file}"
  # protoc --go-grpc_out=. --go-grpc_opt="${opt_m}" "${proto_file}"
}

function usage() {
  echo -n "Utility for generating protobuf and gRPC artifacts for the Pulsar Bulk "
  echo "Beacon Service."
  echo "(c) NS1 Inc."
  echo "Usage:"
  echo "$0 <v1|v2>"
}

if [[ $# != 1 ]]; then
  usage
  exit 1
elif [[ $1 == "-h" || $1 == "--help" ]]; then
  usage
  exit 1
fi

validate_version "${1}"
build "${1}"
