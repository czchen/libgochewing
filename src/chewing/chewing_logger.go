package chewing

type ChewingLogger interface {
    Printf(format string, v ...interface{})
}

type ChewingDefaultLogger struct {}
func (this *ChewingDefaultLogger) Printf(format string, v ...interface{}) {}
