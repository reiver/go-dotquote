package dotquote


type Logger interface {
	Debug(...interface{})
	Debugf(string, ...interface{})

	Trace(...interface{})
	Tracef(string, ...interface{})
}
