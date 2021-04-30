// 父类Logger实现的方法

package MyLog

import (
	"fmt"
	"time"
)

// 给Logger绑定方法，由于文件日志(FileLogger)和控制台日志(StreamLogger)都需要实现输出不同级别日志的方法，
// 将这些方法绑定给他们的父结构体

// enable 判断是否能够执行，只有当比指定的level日志级别大时才执行输出日志
func (l Logger) enable(lv logLevel) (b bool) {
	b = l.level <= lv
	return
}

// Debug 输出debug日志
func (l Logger) Debug(format string, a ...interface{}) {
	/**
	输出日志
	*/
	l.printLog(DEBUG, format, a...)
}

// Trace 输出Trace日志
func (l Logger) Trace(format string, a ...interface{}) {

	l.printLog(TRACE, format, a...)
}

// Info 输出Info日志
func (l Logger) Info(format string, a ...interface{}) {
	l.printLog(INFO, format, a...)
}

// Warning 输出Warning日志
func (l Logger) Warning(format string, a ...interface{}) {
	l.printLog(WARNING, format, a...)
}

// Error 输出Error日志
func (l Logger) Error(format string, a ...interface{}) {
	l.printLog(ERROR, format, a...)
}

// Fatal 输出Fatal日志
func (l Logger) Fatal(format string, a ...interface{}) {

	l.printLog(FATAL, format, a...)
}

// printLog 输出日志，只需要在本包里面使用
func (l Logger) printLog(lv logLevel, format string, a ...interface{}) {
	if l.enable(lv) {
		now := time.Now()
		nowTimeString := now.Format("2006-01-02 15:04:05")
		funcName, fileName, lineNo := getInfo(3)
		msg := fmt.Sprintf(format, a...)
		fmt.Fprint(l.file, fmt.Sprintf("[time: %s] ", nowTimeString))
		fmt.Fprint(l.file, fmt.Sprintf("[level: %s] ", reverseParseLogLevel(lv)))
		fmt.Fprint(l.file, fmt.Sprintf("[%s:%s:%d] ", fileName, funcName, lineNo))
		fmt.Fprint(l.file, fmt.Sprintf("message: %s\n", msg))
		if lv >= ERROR {
			fmt.Fprint(l.errFile, fmt.Sprintf("[time: %s] ", nowTimeString))
			fmt.Fprint(l.errFile, fmt.Sprintf("[level: %s] ", reverseParseLogLevel(lv)))
			fmt.Fprint(l.errFile, fmt.Sprintf("[%s:%s:%d] ", fileName, funcName, lineNo))
			fmt.Fprint(l.errFile, fmt.Sprintf("message: %s\n", msg))
		}
	}
}
