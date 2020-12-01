hello-world app
without external deps
deps as a graph



## Build

bazel build //cmd/hello1:hello
bazel build //cmd/hello2:hello

bazel run //cmd/hello1:hello Siarhei Pazniakou
bazel run //cmd/hello2:hello Siarhei Pazniakou



## Visualize the project’s dependencies

bazel query --notool_deps --noimplicit_deps "deps(//cmd/hello1:hello)" --output graph

bazel query --notool_deps --noimplicit_deps "deps(//cmd/hello2:hello)" --output graph

And use Graphviz http://www.webgraphviz.com/

or

sudo apt update && sudo apt install graphviz xdot

xdot <(bazel query --notool_deps --noimplicit_deps "deps(//cmd/hello1:hello)" --output graph)

xdot <(bazel query --notool_deps --noimplicit_deps "deps(//cmd/hello2:hello)" --output graph)

