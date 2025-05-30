# Go RPC

gRPC is a modern open source high performance Remote Procedure Call (RPC) framework that can run in any environment. It can efficiently connect services in and across data centers with pluggable support for load balancing, tracing, health checking and authentication. It is also applicable in last mile of distributed computing to connect devices, mobile applications and browsers to backend services.
Uses Http2 with latest network protocol.
A simple Go RPC server and client using [grpc](https://grpc.io/)
## Components

- Message Descriptor
- Message Implementation
- Parsing and Serialization

## Types if RPC

- Unary: One request, one response
- Client Streaming: Many requests, one response
- Server Streaming: One request, many responses
- Full Duplex: Many requests, many responses

## Where to use gRPC

- Microservices
- Distributed systems
- Remote procedure calls
- Integrations and APIs

## Installation

First install protobuff

```bash
go get -u google.golang.org/protobuf
go get -u google.golang.org/protobuf/proto
go install google.golang.org/protobuf/cmd/protoc-gen-go
```

Install gRPC

```bash
go get -u google.golang.org/grpc

```


