package main

import (
    "fmt"
    "github.com/docopt/docopt-go"
)

var usage = `
Usage:
	mem-analysis --port=<port> --prefix=<prefix>
	mem-analysis --version
`

func main() {
    arguments, _ := docopt.ParseDoc(usage)
    fmt.Println(arguments)
}