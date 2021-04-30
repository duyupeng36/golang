package main

import (
	log "gitee.com.duyupeng36/MyLog"
)

func main() {
	//streamLogger := log.New("stream", "Warning")
	//for i := 0; i < 5; i++ {
	//	streamLogger.Debug("debug日志")
	//	streamLogger.Trace("Trace日志")
	//	streamLogger.Warning("Warning日志")
	//	streamLogger.Info("Info日志")
	//	streamLogger.Error("Error日志, err:%v", "错误啊")
	//	streamLogger.Fatal("Fatal日志")
	//}
	//
	//log.SetFileName("test.log")
	//log.SetErrorFileName("test.log.err")
	//log.SetFilePath("./")
	//log.SetMaxSize(1024 * 10)
	//fileLogger := log.New("file", "warning")
	//for i := 0; ;i++ {
	//	fileLogger.Debug("debug日志")
	//	fileLogger.Trace("Trace日志")
	//	fileLogger.Warning("Warning日志")
	//	fileLogger.Info("Info日志")
	//	fileLogger.Error("Error日志")
	//	fileLogger.Fatal("Fatal日志")
	//	time.Sleep(time.Second)
	//}
	//

	//log.SetFileName("test.log")
	//log.SetErrorFileName("test.log.err")
	//log.SetFilePath("./")
	//log.SetMaxSize(1024 * 10)
	//log.SetLoggerType("file")
	//log.SetLevel("info")
	log.Info("你好呀 %s", "啦啦啦")
}
