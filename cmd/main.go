package main

import (
	"github.com/imneov/gosql"
)

func main() {
	mb := gosql.NewMemoryBackend()

	gosql.RunRepl(mb)
}
