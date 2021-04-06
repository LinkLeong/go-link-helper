package logging

import (
	"fmt"
	"github.com/a624669980/go-link-helper/file"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

//定义一个int的别名
type Level int

var (
	F                  *os.File
	DefaultPrefix      = ""
	DefaultCallerDepth = 2
	logger             *log.Logger
	logPrefix          = ""
	levelFlags         = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

//iota在const关键字出现时将被重置为0(const内部的第一行之前)，const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。
const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
	FATAL
)

//初始化日志相关内容
func Setup( logPath string) {
	var err error
	var filePath string
	filePath = getLogFilePath()
	if len(logPath)>0 {
		filePath = logPath
	}
	fileName := getLogFileName()
	F, err =file.MustOpen(fileName, filePath)
	if err != nil {
		log.Fatalf("logging.Setup err: %v", err)
	}

	logger = log.New(F, DefaultPrefix, log.LstdFlags)

}

func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v)
}

func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v)
}

func Warn(v ...interface{}) {
	setPrefix(WARN)
	logger.Println(v)
}

func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v)
}

func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Println(v)
}

func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
}



func getLogFilePath() string {
	fmt.Printf("%s%s", "runtime/", "logs/")
	return fmt.Sprintf("%s%s", "runtime/", "logs/")
}

//获取全部长度
func getLogFileFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", "log", time.Now().Format("2006-01-02"), "log")
	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

func mkDir() {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+"/"+getLogFilePath(), os.ModePerm)
	if err != nil {
		//主动触发错误
		panic(err)
	}
}

func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		"log",
		time.Now().Format("2006-01-02"),
		"log",
	)
}

func openLogFile(filePath string) *os.File {
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		mkDir()
	case os.IsPermission(err):
		log.Fatalf("Permission:%v", err)
	}
	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("fail to openfile :%v", err)

	}
	return handle
}
