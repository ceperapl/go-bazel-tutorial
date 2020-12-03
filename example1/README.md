# Example 1

## Foder structure

```
example1
├── cmd
│   ├── hello1
│   │   ├── BUILD
│   │   └── main.go
│   └── hello2
│       ├── BUILD.bazel
│       └── main.go
├── pkg
│   └── greeting
│       ├── BUILD.bazel
│       └── greeting.go
├── README.md
└── WORKSPACE
```

## How to build

```
bazel build //cmd/hello1:hello
bazel build //cmd/hello2:hello

# cross-compile
bazel build --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //cmd/hello1:hello
bazel build --platforms=@io_bazel_rules_go//go/toolchain:windows_amd64 //cmd/hello1:hello
```

## How to run

```
bazel run //cmd/hello1:hello Siarhei Pazniakou
bazel run //cmd/hello2:hello Siarhei Pazniakou
```

## Visualize the project’s dependencies

```
bazel query --notool_deps --noimplicit_deps "deps(//cmd/hello1:hello)" --output graph

bazel query --notool_deps --noimplicit_deps "deps(//cmd/hello2:hello)" --output graph
```

And use Graphviz http://www.webgraphviz.com/

or

```
sudo apt update && sudo apt install graphviz xdot

xdot <(bazel query --notool_deps --noimplicit_deps "deps(//cmd/hello1:hello)" --output graph)

xdot <(bazel query --notool_deps --noimplicit_deps "deps(//cmd/hello2:hello)" --output graph)
```
