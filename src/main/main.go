package main

import (
	"os"
	"fmt"
	"log"
	"path/filepath"
	"path"
)

func main(){
	file, err := os.Create("test.log")
	if nil != err{
		fmt.Println("create file failed")
	}
	logger := log.New(file, "", log.Ldate)
	logPath, _ := filepath.Abs("")
	logPath, fileName := path.Split(logPath)
	logger.Println(logPath)
	logger.Println(fileName)
}
