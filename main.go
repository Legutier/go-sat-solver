package main

import (
	"fmt"
)

func main() {
	implicationGraph, _ := newimplicationGraph()
	status, _ := implicationGraph.addImpliedNode("x1", true, 0)
	fmt.Println(status)
}
