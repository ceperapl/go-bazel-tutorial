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
