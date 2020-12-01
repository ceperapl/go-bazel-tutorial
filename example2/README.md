simple http server
external dependencies
tests
bazel test coverage
gazelle
deps as a graph
update repos from go.mod


## Generate BUILD.bazel files using gazelle

bazel run //:gazelle


## Import external dependencies from go.mod

bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies

## Build service

bazel build //cmd/server:server

## Run service

bazel run //cmd/server:server



## Visualize the project’s dependencies

bazel query --notool_deps --noimplicit_deps "deps(//cmd/server:server)" --output graph

And use Graphviz http://www.webgraphviz.com/

or

sudo apt update && sudo apt install graphviz xdot

xdot <(bazel query --notool_deps --noimplicit_deps "deps(//cmd/server:server)" --output graph)


## Test

bazel test ...

bazel test //pkg/http/handlers:handlers_test
bazel test //pkg/http/handlers:handlers_test
