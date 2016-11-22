package dotquote


type otherLogger interface {
	Log(args ...interface{})
	Logf(format string, args ...interface{})
}


type otherLoggerAdaptor struct {
	OtherLogger otherLogger
}


func (logger otherLoggerAdaptor) Debug(args ...interface{}) {
	logger.OtherLogger.Log(args...)
}
func (logger otherLoggerAdaptor) Debugf(format string, args ...interface{}) {
	logger.OtherLogger.Logf(format, args...)
}

func (logger otherLoggerAdaptor) Error(args ...interface{}) {
	logger.OtherLogger.Log(args...)
}
func (logger otherLoggerAdaptor) Errorf(format string, args ...interface{}) {
	logger.OtherLogger.Logf(format, args...)
}

func (logger otherLoggerAdaptor) Trace(args ...interface{}) {
	logger.OtherLogger.Log(args...)
}
func (logger otherLoggerAdaptor) Tracef(format string, args ...interface{}) {
	logger.OtherLogger.Logf(format, args...)
}

func (logger otherLoggerAdaptor) Warn(args ...interface{}) {
	logger.OtherLogger.Log(args...)
}
func (logger otherLoggerAdaptor) Warnf(format string, args ...interface{}) {
	logger.OtherLogger.Logf(format, args...)
}
