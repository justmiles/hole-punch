package main

import "github.com/justmiles/hole-punch/cmd"

// version of cli. Written during build.
var version = "development"

func main() {
	cmd.Execute(version)
}
