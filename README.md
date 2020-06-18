# pulsar-rum

The bulkbeacon proto code can be compiled (from the `pulsar-rum` directory) with:
`protoc -I bulkbeacon bulkbeacon/bulkbeacon.proto --go_out=plugins=grpc:bulkbeacon`

See https://grpc.io/ for more examples & details regarding dependencies and compiling for other languages.