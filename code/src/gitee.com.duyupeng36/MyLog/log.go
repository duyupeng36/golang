package MyLog

// 实现日志调用的多态，不管时文件日志还是控制台日志都能够输出

func Debug(format string, a ...interface{}) {
	/**
	输出日志
	*/
	std.Debug(format, a...)

}

// Trace 输出Trace日志
func Trace(format string, a ...interface{}) {
	std.Trace(format, a...)
}

// Info 输出Info日志
func Info(format string, a ...interface{}) {
	std.Info(format, a...)
}

// Warning 输出Warning日志
func Warning(format string, a ...interface{}) {
	std.Warning(format, a...)
}

// Error 输出Error日志
func Error(format string, a ...interface{}) {
	std.Error(format, a...)
}

// Fatal 输出Fatal日志
func Fatal(format string, a ...interface{}) {
	std.Fatal(format, a...)
}
