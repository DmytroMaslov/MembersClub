package logger

import "go.uber.org/zap"

var Log *zap.SugaredLogger

func init() {
	logger, _ := zap.NewDevelopment()
	Log = logger.Sugar()
}

func Info(msg string) {
	Log.Infof(msg)
	Log.Sync()
}

func Debag(msg string, data ...interface{}) {
	Log.Debugf(msg, data...)
}
