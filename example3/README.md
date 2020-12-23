# Example 3

## Foder structure

```
├── BUILD.bazel
├── cmd
│   ├── client
│   │   ├── BUILD.bazel
│   │   └── client.go
│   └── server
│       ├── BUILD.bazel
│       └── server.go
├── pkg
│   ├── pb
│   │   ├── BUILD.bazel
│   │   ├── helloworld.pb.go
│   │   ├── helloworld.pb.gw.go
│   │   └── helloworld.proto
│   └── svc
│       ├── BUILD.bazel
│       ├── greeter.go
│       └── greeter_test.go
├── README.md
├── WORKSPACE
├── deps.bzl
├── go.mod
└── go.sum
```

## Generate BUILD.bazel files using Gazelle

```
bazel run //:gazelle
```

## Proto dependencies

https://github.com/bazelbuild/rules_go/blob/master/go/dependencies.rst#proto-dependencies

## gRPC dependencies

https://github.com/bazelbuild/rules_go/blob/master/go/dependencies.rst#grpc-dependencies

## Build gRPC Gateway API

```
bazel build //:api_gateway_grpc
```

Then copy generated files

bazel-bin/api_gateway_grpc/pkg/pb/helloworld.pb.gw.go
bazel-bin/api_gateway_grpc/pkg/pb/helloworld.pb.go

to `pkg/pb`

## gRPC Gateway API + Bazel

https://rules-proto-grpc.github.io/rules_proto_grpc/github.com/grpc-ecosystem/grpc-gateway/

## Import external dependencies from go.mod

```
bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies
```

## Build server & client

```
bazel build //cmd/server:server

bazel build //cmd/client:client
```

## Run server & client

```
bazel run //cmd/server:server

bazel run //cmd/client:client
```

## Test

```
# Test all
bazel test ...

# Separately each test
bazel test //pkg/svc:svc_test
```

## Build Docker images

```
# server
bazel build --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //cmd/server:example3_server_image
bazel run --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //cmd/server:example3_server_image

# client
bazel build --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //cmd/client:example3_client_image
bazel run --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //cmd/client:example3_client_image
```

## Check image

```
docker images | grep example3
> pazdniakousiarhei/cmd/client   example3_client_image   474dfe549c3a   51 years ago   16.3MB
> pazdniakousiarhei/cmd/server   example3_server_image   c63524648e5d   51 years ago   18.7MB
```

## Run in Docker container

```
# server
docker run --rm -p 8080:8080 -p 9090:9090 pazdniakousiarhei/cmd/server:example3_server_image --server.port=9090 --gateway.port=8080

# client
docker run --rm pazdniakousiarhei/cmd/client:example3_client_image --server.address="172.17.0.1" --server.port=9090 --name="Siarhei Pazdniakou"
```

## Curl requests

```
curl --request GET http://localhost:8080/hello/Siarhei%20Pazdniakou
> {"message":"Hello Siarhei Pazdniakou"}

curl --request GET http://localhost:8080/hello/
> {"message":"Hello "}
```
