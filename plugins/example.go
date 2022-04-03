package main

import (
	"errors"
	"fmt"

	parse "github.com/MukhortovDenis/gopostcss_parser"
)

func Run(ast *parse.AST) error {
	if ast.Tokens == nil {
		return errors.New("не получен ast")
	}
	fmt.Println("Плагин работает")
	return nil
}
