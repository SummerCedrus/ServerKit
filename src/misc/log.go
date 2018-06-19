package misc

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
	run = "run"
	err = "error"
	DEBUG = "\x1b[034;1m[DEBUG]\x1b[034;0m"
	LOG = "\x1b[034;1m[  LOG]\x1b[034;0m"
	ERROR = "\x1b[031;1m[ERROR]\x1b[031;0m"
	WARN = "\x1b[033;1m[ WARN]\x1b[033;0m"
)
var LoggerMap map[string]*log.Logger
var fileNameMap map[string]string

func init() {
	LoggerMap = make(map[string]*log.Logger, 0)
	fileNameMap = make(map[string]string, 0)
}
func InitLog(typ string, fileName string) *log.Logger{
	logPath, _ := filepath.Abs("")
	logPath, _ = path.Split(logPath)
	fileName = fileName + "_" + typ + ".log"
	logPath = filepath.Join(logPath,"log",fileName)
	file, err := os.OpenFile(logPath,os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if nil != err{
		fmt.Println("open file failed [%s]", err.Error())
		panic("init log failed!")
		return nil
	}

	logger := log.New(file, "", log.LstdFlags)

	LoggerMap[typ] = logger

	fileNameMap[typ] = fileName

	if typ == run{
		log.SetOutput(file)
	}

	return logger
}

func getFileName(typ string) string{
	if name, ok := fileNameMap[typ]; ok{
		return name
	}
	return typ
}
func getLogger(typ string) *log.Logger{
	if logger, ok := LoggerMap[typ]; ok{
		return logger
	}
	fileName := getFileName(typ)
	return InitLog(typ, fileName)
}
func formatLog(tag string, content string) string{
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
	traceStr := fmt.Sprintf("[%s() %s:%d]", funcName, file, line)

	return fmt.Sprintf("%s%s %s", tag, content, traceStr)
}

func Logf(format string, v ...interface{}){
	commonf(LOG, format, v ...)
}

func Log(v ...interface{}){
	common(LOG, v ...)
}

func Debugf(format string, v ...interface{}){
	commonf(DEBUG, format, v ...)
}

func Debug(v ...interface{}){
	common(DEBUG, v ...)
}

func Warnf(format string, v ...interface{}){
	commonf(WARN, format, v ...)
}

func Warn(v ...interface{}){
	common(WARN, v ...)
}

func Errorf(format string, v ...interface{}){
	s := commonf(ERROR, format, v ...)
	logger := getLogger(err)
	logger.Println(s)
}

func Error(v ...interface{}){
	s := common(ERROR, v ...)
	logger := getLogger(err)
	logger.Println(s)
}

func common(tag string, v ...interface{}) string{
	s := fmt.Sprint(v ...)
	//fmt.Println(s)
	s = formatLog(tag, s)
	log.Println(s)

	return s
}

func commonf(tag string, format string, v ...interface{}) string{
	s := fmt.Sprintf(format, v ...)
	s = formatLog(tag, s)
	log.Println(s)

	return s
}