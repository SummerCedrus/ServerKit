package logkit

import (
	"path/filepath"
	"path"
	"os"
	"log"
	"fmt"
	"runtime"
)

//prefix define
const (
	DEBUG = "[DEBUG]"
	LOG = "[LOG]"
	ERROR = "[ERROR]"
	WARN = "[WARN]"
)
var LoggerMap map[string]*log.Logger

func init() {
	LoggerMap = make(map[string]*log.Logger, 0)
	initLog("run.log")
	initLog("error.log")
}
func initLog(fileName string){
	logPath, _ := filepath.Abs("")
	logPath, _ = path.Split(logPath)
	logPath = filepath.Join(logPath,"log",fileName)
	file, err := os.OpenFile(logPath,os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if nil != err{
		fmt.Println("open file failed [%s]", err.Error())
	}
	logger := log.New(file, "", log.Ldate|log.Ltime)

	LoggerMap[fileName] = logger
}

func writeLog(fileName string, content string){
	pc, file, line, ok := runtime.Caller(3)
	if !ok {
		file = "unknown"
		line = 0
	}
	funcName := "unknown"
	ffpc := runtime.FuncForPC(pc)
	if nil != ffpc{
		funcName = ffpc.Name()
	}
	traceStr := fmt.Sprintf("[%s() %s:%s]", funcName, file, line) 
	if logger, ok := LoggerMap[fileName]; !ok{
		logger.Printf("%v %v", content, traceStr)
	}
}

func Debug(content string)  {
}

func Log()  {

}

func Error(){

}

