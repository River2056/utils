package main

import (
	"flag"
	"plugin"
	"utils/common"
)

func loadPlugin(pluginPath, funcName string) func() {
    pluginModule, err := plugin.Open(pluginPath)
    common.CheckError(err)
    pluginFunc, err := pluginModule.Lookup(funcName)
    common.CheckError(err)
    return pluginFunc.(func())
}

func main() {
    pluginModule := flag.String("p", "", "provide plugin name")
    pluginFunc := flag.String("f", "", "function to run")
    flag.Parse()

    fun := loadPlugin(*pluginModule, *pluginFunc)
    fun()
}
