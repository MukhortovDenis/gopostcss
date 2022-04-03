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

// func Test_getNamePlug(t *testing.T){
// 	t.Run("Проверка получения имени плагина", func(t *testing.T) {
// 		name := getNamePlug("https://github.com/MukhortovDenis/gopostcss_parser")
// 		if name != "gopostcss_parser"{
// 			t.Errorf("smthg gone wrong: %v", name)
// 		}
// 	})
// }
