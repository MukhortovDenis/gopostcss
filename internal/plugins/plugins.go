package plugins

import (
	"errors"
	"fmt"
	"path/filepath"
	"plugin"

	parse "github.com/MukhortovDenis/gopostcss_parser"
)

type TODO interface {
	Run(*parse.AST) error
}

func GetPlugins(ast *parse.AST) error {
	all_plugins, err := filepath.Glob("./plugins/*.so")
	if err != nil {
		return err
	}
	if len(all_plugins) == 0 {
		return errors.New("no plugs")
	}
	for i := range all_plugins {
		fmt.Println(all_plugins[i])
		plug, err := plugin.Open(all_plugins[i])
		if err != nil {
			return err
		}
		fmt.Println(all_plugins[i])
		symRun, err := plug.Lookup("Run")
		if err != nil {
			return err
		}
		run, ok := symRun.(func(ast *parse.AST) error)
		if !ok {
			return errors.New("unexpected type from module symbol")
		}
		if err = run(ast); err != nil {
			return err
		}

	}
	return nil
}
