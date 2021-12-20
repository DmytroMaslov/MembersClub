package log

import "go.uber.org/zap"

type CustomLogger struct {
	log *zap.SugaredLogger
}

var log CustomLogger

func init() {
	logger, _ := zap.NewDevelopment()
	log.log = logger.Sugar()
}
func Get() CustomLogger {
	return log
}
func (log CustomLogger) Info(msg string, data ...interface{}) {
	log.log.Infof(msg, data...)
	log.log.Sync()
}

func (log CustomLogger) Debug(msg string, data ...interface{}) {
	log.log.Debugf(msg, data)
	log.log.Sync()
}

func (log CustomLogger) Error(msg string, data ...interface{}) {
	log.log.Errorf(msg, data)
	log.log.Sync()
}
