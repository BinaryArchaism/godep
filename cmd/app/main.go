package main

import (
	"os"

	"github.com/BinaryArchaism/godep/internal"
)

func main() {
	internal.ParseArgs(os.Args[1:])
}
