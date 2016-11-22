package dotquote


type internalDiscardLogger struct{}

func (internalDiscardLogger) Debug(...interface{}) {}
func (internalDiscardLogger) Debugf(string, ...interface{}) {}

func (internalDiscardLogger) Error(...interface{}) {}
func (internalDiscardLogger) Errorf(string, ...interface{}) {}

func (internalDiscardLogger) Trace(...interface{}) {}
func (internalDiscardLogger) Tracef(string, ...interface{}) {}

func (internalDiscardLogger) Warn(...interface{}) {}
func (internalDiscardLogger) Warnf(string, ...interface{}) {}

