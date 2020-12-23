# Example 2

## Foder structure

```
example2
├── cmd
│   └── server
│       ├── BUILD.bazel
│       └── server.go
├── pkg
│   └── http
│       ├── handlers
│       │   ├── BUILD.bazel
│       │   ├── handlers_test.go
│       │   ├── hello.go
│       │   └── panic.go
│       └── middlewares
│           ├── BUILD.bazel
│           ├── logger.go
│           ├── middlewares_test.go
│           ├── multiple.go
│           └── recover.go
├── BUILD.bazel
├── README.md
├── WORKSPACE
├── deps.bzl
├── go.mod
└── go.sum
```

## Generate BUILD.bazel files using gazelle

```
bazel run //:gazelle
```

## Import external dependencies from go.mod

```
bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies
```

## Build service

```
bazel build //cmd/server:server
```

## Run service

```
bazel run //cmd/server:server
```


## Visualize the project’s dependencies

```
bazel query --notool_deps --noimplicit_deps "deps(//cmd/server:server)" --output graph
```

And use Graphviz http://www.webgraphviz.com/

or

```
xdot <(bazel query --notool_deps --noimplicit_deps "deps(//cmd/server:server)" --output graph)
```

## Test

```
# Test all
bazel test ...

# Separately each test
bazel test //pkg/http/handlers:handlers_test
bazel test //pkg/http/middlewares:middlewares_test
```

## Build Docker image

```
bazel build --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //cmd/server:example2_server_image
bazel run --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //cmd/server:example2_server_image
```

## Check image

```
docker images | grep example2
> pazdniakousiarhei/cmd/server   example2_server_image   1abbe2c0fced   51 years ago   11.2MB
```

## Run in Docker container

```
docker run --rm -p 8080:8080 pazdniakousiarhei/cmd/server:example2_server_image
```

## Curl requests

```
curl --request GET http://localhost:8080/hello/Siarhei%20Pazdniakou
> Hello Siarhei Pazdniakou

curl --request POST http://localhost:8080/panic
> Internal Server Error
```
