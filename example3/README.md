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
