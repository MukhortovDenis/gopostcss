package plugins

import (
	"errors"
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
		plug, err := plugin.Open(all_plugins[i])
		if err != nil {
			return err
		}
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
