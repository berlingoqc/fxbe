package main

import (
	"os"

	"git.wquintal.ca/berlingoqc/fxbe/api"
	"git.wquintal.ca/berlingoqc/fxbe/files"
)

func main() {
	api.Context[files.RootKey] = os.Getenv("HOME")
	api.StartWebServer()
}
