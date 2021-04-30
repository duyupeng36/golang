package MyLog

import (
	"fmt"
	"os"
	"path"
	"time"
)

// FileLogger 文件日志句柄
type FileLogger struct {
	Logger
	filePath    string // 日志文件保存的路径
	fileName    string // 日志文件保存的文件名
	errFileName string // 错误日志保存的文件名
	maxFileSize int64  // 文件切割的大小
}

//// enable 判断是否能够执行，只有当比指定的level日志级别大时才执行输出日志
//func (l *Logger) enable(lv logLevel) (b bool){
//	b = l.level <= lv
//	return
//}

// Debug 输出debug日志
func (f *FileLogger) Debug(format string, a ...interface{}) {
	/**
	输出日志
	*/
	f.printLog(DEBUG, format, a...)
}

// Trace 输出Trace日志
func (f *FileLogger) Trace(format string, a ...interface{}) {
	f.printLog(TRACE, format, a...)
}

// Info 输出Info日志
func (f *FileLogger) Info(format string, a ...interface{}) {
	f.printLog(INFO, format, a...)
}

// Warning 输出Warning日志
func (f *FileLogger) Warning(format string, a ...interface{}) {
	f.printLog(WARNING, format, a...)
}

// Error 输出Error日志
func (f *FileLogger) Error(format string, a ...interface{}) {
	f.printLog(ERROR, format, a...)
}

// Fatal 输出Fatal日志
func (f *FileLogger) Fatal(format string, a ...interface{}) {
	f.printLog(FATAL, format, a...)

}

func (f *FileLogger) Close() {
	err := f.file.Close()
	if err != nil {
		fmt.Println("文件关闭错误")
	}
	err = f.errFile.Close()
	if err != nil {
		fmt.Println("文件关闭错误")
	}
}

func (f *FileLogger) printLog(lv logLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		now := time.Now()
		nowTimeString := now.Format("2006-01-02 15:04:05")
		funcName, fileName, lineNo := getInfo(3)
		msg := fmt.Sprintf(format, a...)
		if !f.chickFileSize(f.file) {
			writeLogs(f.file, nowTimeString, lv, fileName, funcName, lineNo, msg)
		} else {
			f.file = f.splitFile(f.file, f.filePath, f.fileName)

			// 记录日志
			writeLogs(f.file, nowTimeString, lv, fileName, funcName, lineNo, msg)
		}

		if lv >= ERROR {
			if !f.chickFileSize(f.errFile) {
				writeLogs(f.errFile, nowTimeString, lv, fileName, funcName, lineNo, msg)
			} else {
				f.errFile = f.splitFile(f.errFile, f.filePath, f.errFileName)
				// 记录日志
				writeLogs(f.errFile, nowTimeString, lv, fileName, funcName, lineNo, msg)
			}

		}
	}
}

func (f *FileLogger) chickFileSize(file *os.File) bool {
	info, err := file.Stat()
	if err != nil {
		return false
	}
	size := info.Size()
	if size >= f.maxFileSize {
		// 当前文件大小如果大于或等于文件的最大值，要切割文件，返回true
		return true
	} else {
		return false
	}

}

func (f *FileLogger) splitFile(file *os.File, filePath, fileName string) *os.File {
	// 关闭文件，重命名
	file.Close()
	nowString := time.Now().Format("0102150405")
	logName := path.Join(filePath, fileName)
	bakName := fmt.Sprintf("%s_%s.bak", logName, nowString)
	os.Rename(logName, bakName) // 重命名

	ret, _ := os.OpenFile(logName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644) // 打开新的文件
	return ret
}
