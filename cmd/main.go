package main

import (
	"log"

	"github.com/MukhortovDenis/gopostcss/internal/cli"
	parsing "github.com/MukhortovDenis/gopostcss/internal/parser"
	"github.com/MukhortovDenis/gopostcss/internal/plugins"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	filename, newFilename, _, err := cli.InitFlags()
	if err != nil {
		log.Fatal(err)
	}
	ast, err := parsing.Parse(filename)
	if err != nil {
		log.Fatal(err)
	}
	if err := plugins.GetPlugins(ast); err != nil {
		log.Fatal(err)
	}
	parsing.Create(newFilename, ast)
	return nil
}
