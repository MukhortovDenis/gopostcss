package plugins

import (
	"testing"

	parse "github.com/MukhortovDenis/gopostcss_parser"
)

func Test_GetPlugins(t *testing.T) {
	t.Run("Проверка скачавания плагинов", func(t *testing.T) {
		ast := parse.AST{}
		if err := GetPlugins(&ast); err != nil {
			t.Error(err)
		}
	})
}