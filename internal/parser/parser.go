package parsing

import (
	"log"

	parse "github.com/MukhortovDenis/gopostcss_parser"
)

func Parse(filename string) (*parse.AST, error) {
	ast, err := parse.ParseIntoAST(filename)
	if err != nil {
		return nil, err
	}
	return ast, nil
}

func Create(newFilename string, ast *parse.AST) {
	if err := parse.ParseIntoCSS(ast, newFilename); err != nil {
		log.Fatal(err)
	}
}
