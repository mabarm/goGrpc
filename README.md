# goGrpc

Grpc use in go

# Steps

1. Install go
2. Install protobuff
3. Install golang grpc and set path

# protoc --go_out=. --go-grpc_out=. proto/greet.proto


About protobuff- 
Protocol Buffers (Protobuf) is a language- and platform-neutral data serialization format developed by Google. 
It uses .proto files to define structured data (messages), which are then compiled into efficient binary formats for fast transmission and storage. 
Protobuf is compact, faster than JSON or XML, supports cross-language compatibility, and ensures backward and forward compatibility for evolving data models. 
It’s widely used in scenarios like gRPC for service communication.