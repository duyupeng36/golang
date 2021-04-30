// 该日志包公用的属性和函数

package MyLog

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"
)

// Logger 日志句柄父类
type Logger struct {
	level   logLevel
	file    *os.File
	errFile *os.File
}

// LoggerInterface 定义一个接口
type LoggerInterface interface {
	Debug(format string, a ...interface{})
	Trace(format string, a ...interface{})
	Info(format string, a ...interface{})
	Warning(format string, a ...interface{})
	Error(format string, a ...interface{})
	Fatal(format string, a ...interface{})
	Close()
}

// 配置文件日志时使用的变量
var (
	logFilePath   string = "./"
	logFileName   string
	errorFileName string
	maxSize       int64  = 10 * 1024 * 1024
	loggerType    string = "stream"
	level         string = "debug"
)
var std LoggerInterface

func init() {
	// 读取配置文件，并设置loggerType和level
	f, _ := os.Open("./settings.json")
	ret, _ := ioutil.ReadAll(f)
	var c struct {
		LogFilePath   string
		LogFileName   string
		ErrorFileName string
		MaxSize       int64
		LoggerType    string
		Level         string
	}
	json.Unmarshal(ret, &c)

	loggerType = c.LoggerType
	level = c.Level
	logFilePath = c.LogFilePath
	logFileName = c.LogFileName
	errorFileName = c.ErrorFileName
	maxSize = c.MaxSize
	std = New(loggerType, level)
}

// 日志等级
type logLevel uint8

const (
	UNKNOWN logLevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

// parseLogLevel 日志等级解析函数
func parseLogLevel(level string) (logLevel, error) {
	level = strings.ToUpper(level)
	switch level {
	case "DEBUG":
		return DEBUG, nil
	case "TRACE":
		return TRACE, nil
	case "INFO":
		return INFO, nil
	case "WARNING":
		return WARNING, nil
	case "ERROR":
		return ERROR, nil
	case "FATAL":
		return FATAL, nil
	default:
		err := errors.New("日志等级解析失败")
		return UNKNOWN, err
	}
}

// reverseParseLogLevel 解析日志等级对应的字符串
func reverseParseLogLevel(lv logLevel) (level string) {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

// 使用工厂模式来创建新的Logger

// Factory 工厂结构体
type Factory struct{}

func (f *Factory) New(loggerType string, level string) (logger LoggerInterface) {
	loggerType = strings.ToUpper(loggerType)
	lv, _ := parseLogLevel(level)
	switch loggerType {
	case "STREAM":
		var streamLogger = StreamLogger{
			Logger{
				level:   lv,
				file:    os.Stdout,
				errFile: os.Stderr,
			},
		}
		logger = streamLogger

	case "FILE":
		fullFilePath := path.Join(logFilePath, logFileName)
		file, _ := os.OpenFile(fullFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		fullErrFilePath := path.Join(logFilePath, errorFileName)
		errFile, _ := os.OpenFile(fullErrFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		var fileLogger = &FileLogger{
			Logger: Logger{
				level:   lv,
				file:    file,
				errFile: errFile,
			},
			filePath:    logFilePath,
			fileName:    logFileName,
			errFileName: errorFileName,
			maxFileSize: maxSize,
		}
		logger = fileLogger
	}
	return
}

// New 工厂的多态实现
func New(loggerType string, level string) (logger LoggerInterface) {
	factory := &Factory{} // 新建工厂
	logger = factory.New(loggerType, level)
	return
}

// 获取产生日志的文件名、行号和函数名
func getInfo(n int) (funcName, fileName string, lineNo int) {
	var pc uintptr
	var file string
	var line int
	var ok bool
	pc, file, line, ok = runtime.Caller(n)
	if !ok {
		fmt.Println("runtime.Caller() failed")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	lineNo = line
	return
}

//// printLog 输出日志，只需要在本包里面使用
//func printLog(file io.Writer, lv logLevel, format string, a ...interface{}) {
//	now := time.Now()
//	nowTimeString := now.Format("2006-01-02 15:04:05")
//	funcName, logFileName, lineNo := getInfo(3)
//	msg := fmt.Sprintf(format, a...)
//	fmt.Fprint(file, fmt.Sprintf("[time: %s] ", nowTimeString))
//	fmt.Fprint(file, fmt.Sprintf("[level: %s] ", reverseParseLogLevel(lv)))
//	fmt.Fprint(file, fmt.Sprintf("[%s:%s:%d] ", logFileName, funcName, lineNo))
//	fmt.Fprint(file, fmt.Sprintf("message: %s\n", msg))
//
//}

func writeLogs(file *os.File, nowTimeString string, lv logLevel, fileName, funcName string, lineNo int, msg string) {
	fmt.Fprint(file, fmt.Sprintf("[time: %s] ", nowTimeString))
	fmt.Fprint(file, fmt.Sprintf("[level: %s] ", reverseParseLogLevel(lv)))
	fmt.Fprint(file, fmt.Sprintf("[%s:%s:%d] ", fileName, funcName, lineNo))
	fmt.Fprint(file, fmt.Sprintf("message: %s\n", msg))
}

// SetFilePath 设置文件路径
func SetFilePath(path string) {
	logFilePath = path
}

// SetFileName 设置日志文件名
func SetFileName(name string) {
	logFileName = name
}

// SetErrorFileName 设置错误文件名
func SetErrorFileName(errName string) {
	errorFileName = errName
}

// SetMaxSize 配置文件最大的拆分大小
func SetMaxSize(size int64) {
	maxSize = size
}

// SetLoggerType 设置日志类型（stream和file可选）
func SetLoggerType(Type string) {
	loggerType = Type
}

// SetLevel 设置日志等级 （Debug Trace Info Warning Error Fatal可选）
func SetLevel(lv string) {
	level = lv
}
