Pulsar Real User Monitoring (RUM)
=================================

> This project is in [active development](https://github.com/ns1/community/blob/master/project_status/ACTIVE_DEVELOPMENT.md).

Utilities for working with Pulsar RUM.


Bulk Beacons
------------

A method to send performance data to Pulsar in bulk.  gRPC and HTTP+JSON are
supported.  See comments in the `bulkbeacon.proto` file for how to structure
messages.

### Getting started with gRPC

1. Copy the desired version of 
[bulkbeacon.proto](https://github.com/ns1/pulsar-routemap/proto/bulkbeacon)
into your project.  Use your preferred method of building gRPC clients from that 
`.proto` file.
2. Use `g.ns1p.net:443` as the service's target address.

See the [example Golang client](https://github.com/ns1/pulsar-rum/blob/master/cmd/example_client/grpc_example_client.go) 
for more details.  Additionally, check out https://grpc.io/ for more examples & details
regarding dependencies and compiling for other languages.

### Getting started with HTTP+JSON

For HTTP+JSON use the `b.ns1p.net/v1/bulk/beacon` endpoint.

N.B. You can use the protocol buffer objects as a basis for the required
JSON payload.  You'll need to use helper libraries for your language to
serialize protocol buffer objects to JSON.  Some information is available on 
[the protocol buffer docs](https://github.com/protocolbuffers/protobuf/blob/master/docs/third_party.md)
page.


Building the examples
---------------------

The examples can be built executing (from the `pulsar-rum` directory) with:

```sh
$ make examples
```


Contributing
------------

Pull Requests and issues are welcome. See the [NS1 Contribution Guidelines](https://github.com/ns1/community) 
for more information.


License
-------

Copyright (C) 2020, NSONE, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

