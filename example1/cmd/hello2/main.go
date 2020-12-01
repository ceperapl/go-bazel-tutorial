package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ceperapl/go-bazel-tutorial/example1/pkg/greeting"
)

func main() {
	who := "World"
	if len(os.Args) > 1 {
		who = strings.Join(os.Args[1:], " ")
	}
	fmt.Printf(greeting.GetGreeting(who))
}
