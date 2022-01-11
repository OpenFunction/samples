package main

import (
	"reflect"

	"github.com/OpenFunction/functions-framework-go/plugin"
	pluginCustom "main.go/userfunction/plugins/plugin-custom"
)

func getLocalPlugins() map[string]plugin.Plugin {
	nilPlugins := map[string]plugin.Plugin{}
	localPlugins := map[string]plugin.Plugin{
		pluginCustom.Name: pluginCustom.New(),
	}

	if reflect.DeepEqual(localPlugins, nilPlugins) {
		return nil
	} else {
		return localPlugins
	}
}
