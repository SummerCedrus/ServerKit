package main

import ("fmt"
	"path/filepath"
	"plugin")

func main() {
	rootPath,err := filepath.Abs("testplugin.so")
	if nil != err{
		panic(err)
	}
	p, err := plugin.Open(rootPath)
	if nil != err{
		fmt.Printf("err1: %v \n", err)
		return
	}
	f, err := p.Lookup("Hello")
	if nil != err{
		fmt.Printf("err2: %v \n", err)
		return
	}
	f.(func())()
}