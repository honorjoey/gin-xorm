package log

import (
	"fmt"
	"os"

	//"reflect"
	"time"
)

// LogLevel defines a log level
type LogLevel int

// enumerate all LogLevels
const (
	// !nashtsai! following level also match syslog.Priority value
	LOG_DEBUG LogLevel = iota
	LOG_INFO
	LOG_WARNING
	LOG_ERR
	LOG_OFF
	LOG_UNKNOWN
)

type Logger interface {
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})
	Error(v ...interface{})
	Errorf(format string, v ...interface{})
	Info(v ...interface{})
	Infof(format string, v ...interface{})
	Warn(v ...interface{})
	Warnf(format string, v ...interface{})

	//Level() LogLevel
	//SetLevel(l LogLevel)
}

type Log struct {
	Tag string
	Logger
}

const (
	timeFormat = "2006/01/02 15:04:05.000000"
)

func (l Log) Debug(v ...interface{}) {
	l.logPrintln("Debug", v...)
}

func (l Log) Debugf(format string, v ...interface{}) {
	l.logPrintf("Debug", format, v...)
}

func (l Log) Error(v ...interface{}) {
	l.logPrintln("Error", v...)
}

func (l Log) Errorf(format string, v ...interface{}) {
	l.logPrintf("Error", format, v)
}

func (l Log) Info(v ...interface{}) {
	l.logPrintln("Info", v...)
}

func (l Log) Infof(format string, v ...interface{}) {
	l.logPrintf("Info", format, v...)
}

func (l Log) Warn(v ...interface{}) {
	l.logPrintln("Warn", v...)
}

func (l Log) Warnf(format string, v ...interface{}) {
	l.logPrintf("Warn", format, v...)
}

func (l Log) logPrintln(level string, v ...interface{}) {
	fmt.Fprintf(os.Stdout,"[%s] [%s] %s ",level, l.Tag, time.Now().Format(timeFormat))
	fmt.Fprintln(os.Stdout, v...)
}

func (l Log) logPrintf(level, format string, v ...interface{}) {
	fmt.Fprintf(os.Stdout, "[%s] [%s] %s ",level, l.Tag, time.Now().Format(timeFormat))
	fmt.Fprintf(os.Stdout, format, v...)
}

/*func Level() LogLevel {

}

func SetLevel(l LogLevel) {

}*/
