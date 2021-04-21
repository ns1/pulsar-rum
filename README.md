Pulsar Real User Monitoring (RUM)
=================================

> This project is in [active development](https://github.com/ns1/community/blob/master/project_status/ACTIVE_DEVELOPMENT.md).

Utilities and sample programs for Pulsar RUM.


Bulk Beacons
------------

A method to send performance data to Pulsar in bulk.  gRPC and HTTP+JSON are
supported.

See comments in the `bulkbeacon.proto` definition file for how to structure messages.
Depending on the version of the Bulk Beacon ingestion, you will have to look at its respective
directory (mentioned below in the `Getting started with gRPC` section).

### Versions

We currently support two formats for gRPC and HTTP+JSON for Bulk Beacons. Each one can be
found in `proto/bulkbeacon/{version}/bulkbeacon.proto`, with `version` taking values
of `v1` and `v2`. 

Version 1 was the first formal release of the Bulk Beacon ingestion, and version 2 
supports a new feature for multiple metrics. For more information, please contact
your Customer Success Representative.

### Getting started with gRPC

1. Copy the desired version of `bulkbeacon.proto` from either 
[https://github.com/ns1/pulsar-rum/tree/master/proto/bulkbeacon/v1](https://github.com/ns1/pulsar-rum/tree/master/proto/bulkbeacon/v1)
   or [https://github.com/ns1/pulsar-rum/tree/master/proto/bulkbeacon/v2](https://github.com/ns1/pulsar-rum/tree/master/proto/bulkbeacon/v1) 
into your project. Use your preferred method of building gRPC clients from that `.proto` file.
2. Use `g.ns1p.net:443` as the service's target address.
3. Enable TLS on your gRPC transport.
4. Add your NS1 API key. Please check the examples to see how to add the authentication key.

See the [v1 example Golang client](https://github.com/ns1/pulsar-rum/blob/master/cmd/example_client_v1/main.go) 
or the [v2 example Golang client](https://github.com/ns1/pulsar-rum/blob/master/cmd/example_client_v2/main.go) 
for more details. Additionally, check out https://grpc.io/ for more information 
regarding dependencies and compiling for other languages.

### Getting started with HTTP+JSON

For HTTP+JSON use the `http(s)://b.ns1p.net/v1/beacon/bulk` endpoint.

N.B. You can use the protocol buffer objects as a basis for the required
JSON payload.  You'll need to use helper libraries for your language to
serialize protocol buffer objects to JSON.  Some information is available on 
[the protocol buffer docs](https://github.com/protocolbuffers/protobuf/blob/master/docs/third_party.md)
page.

Building the examples
---------------------

The examples can be built executing several commands from the `pulsar-rum` directory.

To build both examples, execute:
```sh
$ make
```

If you want to build a specific version, execute:
```shell
$ make bulkbeacon_v1
```
or
```shell
$ make bulkbeacon_v2
```

You can find the binaries on the `build/` directory. 

If unsure, reach out with any question you could have.

Contributing
------------

Pull Requests and issues are welcome. See the [NS1 Contribution Guidelines](https://github.com/ns1/community) 
for more information.


License
-------

Copyright (C) 2021, NSONE, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

