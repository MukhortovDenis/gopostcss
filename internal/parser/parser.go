package parser

import (
	parse "github.com/MukhortovDenis/gopostcss_parser"
)

func Parse(filename string, newFilename string, config string) (*parse.AST, error) {
	ast, err := parse.ParseIntoAST(filename)
	if err != nil {
		return nil, err
	}
	return ast, nil
}
