package logger

type LogWriter interface {
	Write(data interface{}) error
}

type Logger struct {
	WriterList []LogWriter
}

func (l *Logger) RegisterWriter(writer LogWriter) {
	l.WriterList = append(l.WriterList, writer)
}

func (l *Logger) Log(data interface{}) {
	for _, writer := range l.WriterList {
		writer.Write(data)
	}
}

func NewLogger() *Logger {
	return &Logger{}
}
