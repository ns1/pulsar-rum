// Copyright 2020 NSONE, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "pulsar-rum/gen/bulkbeacon/v1"
)

var beacons = &pb.Beacons{
	Beacons: []*pb.Beacon{
		{
			Appid: "__APPID__", // FIXME: Your AppID here.
			Measurements: []*pb.Measurement{
				{
					Attribution: &pb.Attribution{
						Jobid: "__JOBID__", // FIXME: Your JobID here.
						Location: &pb.Location{
							GeoCountry: "GB",
							Asn:        2856,
						},
						DeviceType: pb.DeviceType_DESKTOP,
					},
					Payloads: []*pb.Payload{
						{
							StatusCode: 200,
							DataTtl:    7200,
							Value: &pb.Payload_Simple{
								Simple: &pb.SimpleLatency{ValueMs: 50},
							},
						},
					},
				},
			},
		},
	},
}

// auth it's a simple structure to manage the authentication. It implements
// grpc.PerRPCCredentials interface.
type auth struct {
	key string
}

// GetRequestMetadata sets the authentication key into the metadata map. The
// required map key is "x-ns1-key".
func (a auth) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {
	m := map[string]string{
		"x-ns1-key": a.key,
	}
	return m, nil
}

// RequireTransportSecurity must return true if we are using TLS.
func (a auth) RequireTransportSecurity() bool {
	return true
}

func main() {

	address := "g.ns1p.net:443"

	// Debug if needed
	marshaler := jsonpb.Marshaler{}
	m, _ := marshaler.MarshalToString(beacons)
	fmt.Println(string(m))

	fmt.Printf("Go version: %s\n", runtime.Version())

	// Setup authentication
	auth := auth{key: "__YOUR_NS1_API_KEY__"}

	// Set up gRPC connection
	log.Println("dialing")
	conn, err := grpc.Dial(address,
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})),
		grpc.WithBlock(),
		grpc.WithPerRPCCredentials(auth))
	log.Println("dialed")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create client
	c := pb.NewPulsarDataIngestionClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Send beacons
	r, err := c.Ingest(ctx, beacons) // beacons defined / generated above
	if err != nil {
		log.Printf("Error sending beacons: %v", err)
	} else {
		log.Printf("%d datapoints processed (%d failed)", r.Processed, r.Failed)
	}
}
