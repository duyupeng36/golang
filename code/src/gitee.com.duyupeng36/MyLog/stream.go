//

package MyLog

import (
	"fmt"
	"time"
)

// StreamLogger 控制台日志句柄
type StreamLogger struct {
	Logger
}

//// enable 判断是否能够执行
//func (s StreamLogger) enable(lv logLevel) (b bool){
//	b = s.level <= lv
//	return
//}
//
//// Debug 输出debug日志
//func (s StreamLogger) Debug(msg string)  {
//	if s.enable(DEBUG){
//		fmt.Println(msg)
//	}
//}
//
//// Trace 输出Trace日志
//func (s StreamLogger)Trace(msg string) {
//	if s.enable(TRACE) {
//		fmt.Println(msg)
//	}
//}
//
//// Info 输出Info日志
//func (s StreamLogger)Info(msg string) {
//	if s.enable(INFO) {
//		fmt.Println(msg)
//	}
//}
//
//// Warning 输出Warning日志
//func (s StreamLogger)Warning(msg string) {
//	if s.enable(WARNING) {
//		fmt.Println(msg)
//	}
//
//}
//
//// Error 输出Error日志
//func (s StreamLogger)Error(msg string) {
//	if s.enable(ERROR) {
//		fmt.Println(msg)
//	}
//}
//
//// Fatal 输出Fatal日志
//func (s StreamLogger)Fatal(msg string) {
//	if s.enable(FATAL){
//		fmt.Println(msg)
//	}
//}

func (s StreamLogger) Close() {

}

// printLog 输出日志，只需要在本包里面使用
func (s StreamLogger) printLog(lv logLevel, format string, a ...interface{}) {
	if s.enable(lv) {
		now := time.Now()
		nowTimeString := now.Format("2006-01-02 15:04:05")
		funcName, fileName, lineNo := getInfo(3)
		msg := fmt.Sprintf(format, a...)
		fmt.Fprint(s.file, fmt.Sprintf("[time: %s] ", nowTimeString))
		fmt.Fprint(s.file, fmt.Sprintf("[level: %s] ", reverseParseLogLevel(lv)))
		fmt.Fprint(s.file, fmt.Sprintf("[%s:%s:%d] ", fileName, funcName, lineNo))
		fmt.Fprint(s.file, fmt.Sprintf("message: %s\n", msg))
		if lv >= ERROR {
			fmt.Fprint(s.errFile, fmt.Sprintf("[time: %s] ", nowTimeString))
			fmt.Fprint(s.errFile, fmt.Sprintf("[level: %s] ", reverseParseLogLevel(lv)))
			fmt.Fprint(s.errFile, fmt.Sprintf("[%s:%s:%d] ", fileName, funcName, lineNo))
			fmt.Fprint(s.errFile, fmt.Sprintf("message: %s\n", msg))
		}
	}
}
