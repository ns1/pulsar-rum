// Copyright 2021 NSONE, Inc.
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
	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"pulsar-rum/pkg/bulkbeacon"
	"runtime"
	"time"

	pb "pulsar-rum/pkg/bulkbeacon/bulkbeacon_v2"
)

var (
	appID   = "__appID__"   // FIXME: Your AppID here.
	jobID   = "__jobID__"   // FIXME: Your JobID here.
	authKey = "__authKey__" // FIXME: Your NS1 API key here.
)

var beacons = &pb.Beacons{
	Metrics: []*pb.Metrics{
		{
			Attribution: &pb.Attribution{
				Jobid: jobID,
				Appid: appID,
				Location: &pb.Attribution_GeoAsn{
					// GEO+ASN
					GeoAsn: &pb.GeoAsn{
						GeoCountry: "US",
						GeoSubdiv:  "NY",
						Asn:        2856,
					},
				},
			},
			Payloads: []*pb.Payload{
				{
					Metric: &pb.Payload_Latency{
						Latency: &pb.LatencyMetric{
							Value: 50,
						},
					},
				},
				{
					Metric: &pb.Payload_PerfScore{
						PerfScore: &pb.PerformanceScoreMetric{
							Value: 55,
							Meta: &pb.StaticMetricMetadata{
								Ttl: 3200,
							},
						},
					},
				},
				{
					Metric: &pb.Payload_Avail{
						Avail: &pb.AvailabilityMetric{
							StatusCode: 200,
						},
					},
				},
				{
					Metric: &pb.Payload_AvailScore{
						AvailScore: &pb.AvailabilityScoreMetric{
							Value: 1.0,
							Meta: &pb.StaticMetricMetadata{
								Ttl: 7200,
							},
						},
					},
				},
			},
		},
		{
			Attribution: &pb.Attribution{
				Appid: appID,
				Jobid: jobID,
				Location: &pb.Attribution_Ip{
					// IP Address instead of GEO+ASN.
					Ip: &pb.IPAddress{
						Address: "1.2.3.4",
					},
				},
			},
			Payloads: []*pb.Payload{
				{
					Metric: &pb.Payload_PerfScore{
						// Static value without metadata.
						PerfScore: &pb.PerformanceScoreMetric{
							Value: 70,
						},
					},
				},
				{
					Metric: &pb.Payload_Avail{
						Avail: &pb.AvailabilityMetric{
							StatusCode: 200,
						},
					},
				},
			},
		},
		{
			Attribution: &pb.Attribution{
				Appid: appID,
				Jobid: jobID,
				Location: &pb.Attribution_GeoAsn{
					GeoAsn: &pb.GeoAsn{
						// No subdiv or ASN
						GeoCountry: "US",
					},
				},
			},
			Payloads: []*pb.Payload{
				{
					Metric: &pb.Payload_Latency{
						Latency: &pb.LatencyMetric{
							Value: 65,
						},
					},
				},
				{
					// Static value for avail.
					Metric: &pb.Payload_AvailScore{
						AvailScore: &pb.AvailabilityScoreMetric{
							Value: 0.98,
							// no static meta, TTL used comes from the job config.
						},
					},
				},
			},
		},
	},
}

func main() {
	address := "g.ns1p.net:443"

	// Debug if needed
	marshaler := jsonpb.Marshaler{}
	m, _ := marshaler.MarshalToString(beacons)
	fmt.Println(string(m))

	fmt.Printf("Go version: %s\n", runtime.Version())

	// Setup authentication
	auth := bulkbeacon.NewAuth(authKey)

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
	var r *pb.IngestResult
	r, err = c.Ingest(ctx, beacons) // beacons defined / generated above
	if err != nil {
		log.Printf("Error sending beacons: %v", err)
	} else {
		log.Printf("%d datapoints processed (%d payloads failed)", r.Processed, r.Failed)
	}
}
