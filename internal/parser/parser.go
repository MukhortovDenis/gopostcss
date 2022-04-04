package parsing

import (
	"log"

	parse "github.com/MukhortovDenis/gopostcss_parser"
)

// Parse is func where using gopostcss_parser, create AST
func Parse(filename string) (*parse.AST, error) {
	ast, err := parse.ParseIntoAST(filename)
	if err != nil {
		return nil, err
	}
	return ast, nil
}

// Create is func where using gopostcss_parser, create new CSS file
func Create(newFilename string, ast *parse.AST) {
	if err := parse.ParseIntoCSS(ast, newFilename); err != nil {
		log.Fatal(err)
	}
}
