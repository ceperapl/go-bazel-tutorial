# Examples to build Go code

In this tutorial, you will learn the basics of building Golang applications using Bazel.

## Example 1

The first example is very simple. Here's how to build a Go application that don't have external dependencies and consists of several packages.

## Example 2

The second example shows building an application with external dependencies. The Gazelle tool is used to generate BUILD.bazel files. Also, this example shows how to run tests using Bazel.

## Example 3

The third example demonstrates building and testing a service (server and client) that uses gRPC and gRPC-Gateway.

## TODO

- Show test coverage with Bazel

## Issues

- There is a possibility that after installing the bazel on Ubuntu 20.04, it will be necessary to fix a python-related error when building docker images.

```
sudo apt install python-is-python3
```

- After building Docker images we have images created around 51 years ago (see readme for example2 and example3)
