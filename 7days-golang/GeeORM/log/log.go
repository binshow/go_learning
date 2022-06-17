package log

import (
	"io/ioutil"
	"log"
	"os"
	"sync"
)

// -------------------------------------------
// @file          : log.go
// @author        : binshow
// @time          : 2022/6/16 12:15 AM
// @description   :	实现一个简单的log库
// -------------------------------------------


//这个简易的 log 库具备以下特性：
//
//支持日志分级（Info、Error、Disabled 三级）。
//不同层级日志显示时使用不同的颜色区分。
//显示打印日志代码对应的文件名和行号。

var (
	errorLog = log.New(os.Stdout, "\033[31m[error]\033[0m ", log.LstdFlags|log.Lshortfile)
	infoLog  = log.New(os.Stdout, "\033[34m[info ]\033[0m ", log.LstdFlags|log.Lshortfile)
	loggers  = []*log.Logger{errorLog, infoLog}
	mu       sync.Mutex
)

// log methods
var (
	Error  = errorLog.Println
	Errorf = errorLog.Printf
	Info   = infoLog.Println
	Infof  = infoLog.Printf
)

// log levels
const (
	InfoLevel = iota
	ErrorLevel
	Disabled
)

// SetLevel controls log level
func SetLevel(level int) {
	mu.Lock()
	defer mu.Unlock()

	for _, logger := range loggers {
		logger.SetOutput(os.Stdout)
	}

	if ErrorLevel < level {
		errorLog.SetOutput(ioutil.Discard)
	}
	if InfoLevel < level {
		infoLog.SetOutput(ioutil.Discard)
	}
}