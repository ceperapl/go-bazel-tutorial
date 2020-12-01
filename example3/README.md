protobuf + grpc
grpc gateway (REST)
tests
shell commands
build docker images
external dependencies
gazelle
deps as a graph
update repos from go.mod


## Setup Gazelle

Copy code to WORKSPACE from here https://github.com/bazelbuild/bazel-gazelle#setup

Add this code to root BUILD.bazel

```
load("@bazel_gazelle//:def.bzl", "gazelle")

gazelle(
    name = "gazelle",
    prefix = "github.com/ceperapl/go-bazel-tutorial/example2",
)
```

## Generate BUILD.bazel files using Gazelle

bazel run //:gazelle

## Proto dependencies

https://github.com/bazelbuild/rules_go/blob/master/go/dependencies.rst#proto-dependencies

## gRPC dependencies

https://github.com/bazelbuild/rules_go/blob/master/go/dependencies.rst#grpc-dependencies

## Build gRPC Gateway API

bazel build //:api_gateway_grpc

Then copy generated files

bazel-bin/api_gateway_grpc/pkg/pb/helloworld.pb.gw.go
bazel-bin/api_gateway_grpc/pkg/pb/helloworld.pb.go

to pkg/pb

## gRPC Gateway API + Bazel

https://rules-proto-grpc.github.io/rules_proto_grpc/github.com/grpc-ecosystem/grpc-gateway/

## Import external dependencies from go.mod

bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies

## Build server/client

bazel build //cmd/server:server

bazel build //cmd/client:client

## Run server/client

bazel run //cmd/server:server

bazel run //cmd/client:client

## Visualize the project’s dependencies

bazel query --notool_deps --noimplicit_deps "deps(//cmd/server:server)" --output graph

And use Graphviz http://www.webgraphviz.com/

or

sudo apt update && sudo apt install graphviz xdot

xdot <(bazel query --notool_deps --noimplicit_deps "deps(//cmd/server:server)" --output graph)

