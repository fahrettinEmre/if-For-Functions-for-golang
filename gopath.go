package main

import (
	"fmt"
	"os"
)

func main() {
	processorArchitecture := os.Getenv("PROCESSOR_ARCHITECTURE")
	gopath := os.Getenv("GOPATH")
	fmt.Println("işlem mimarisi :" + processorArchitecture)
	fmt.Println(gopath)
}
