package plugins

import (
	"errors"
	"path/filepath"
	"plugin"
	"strings"

	parse "github.com/MukhortovDenis/gopostcss_parser"
)

type TODO interface {
	Run(*parse.AST) error
}



func GetPlugins(ast *parse.AST) error {
	all_plugins, err := filepath.Glob("plugins/*.so")
	if err != nil {
		return err
	}
	for i := range all_plugins {
		plug, err := plugin.Open(all_plugins[i])
		if err != nil {
			return err
		}
		symTODO, err := plug.Lookup("TODO")
		if err != nil {
			return err
		}
		var todo TODO
		todo, ok := symTODO.(TODO)
		if !ok {
			return errors.New("unexpected type from module symbol")
		}
		todo.Run(ast)

	}
	return nil
}

func getNamePlug(name string) string {
	slice := strings.Split(name, "/")
	return slice[len(slice)-1]
}
