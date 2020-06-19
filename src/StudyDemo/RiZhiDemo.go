package main

type LogWriter interface {
	Write(data interface{}) error
}
type Logger struct {
	writerLisr []LogWriter
}

func (l *Logger) RegisterWriter(writer LogWriter){
	l.writerLisr = append(l.writerLisr,writer)
}

func (l *Logger) Log(data interface{}) {
	for _,writer := range l.writerLisr{
		writer.Write(data)
	}
}

func NewLogger()*Logger  {
	return &Logger{}
}
