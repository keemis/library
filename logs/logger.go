package logs

import (
	"fmt"
	"path"
	"runtime"
	"strconv"
	"strings"
)

// Logger po
type Logger struct {
	traceID       string
	funcCallDepth int
}

// Option 配置选项
type Option func(*Logger)

// SetTraceID 设置TraceID
func SetTraceID(traceID string) Option {
	return func(log *Logger) {
		log.traceID = traceID
	}
}

// SetFuncCallDepth 设置堆栈深度
func SetFuncCallDepth(depth int) Option {
	return func(log *Logger) {
		log.funcCallDepth = depth
	}
}

// New logger
func New(options ...Option) Logger {
	logger := Logger{
		funcCallDepth: 3,
	}
	for _, option := range options {
		option(&logger)
	}
	if logger.traceID == "" {
		logger.traceID = NewTraceID()
	}
	return logger
}

// Critical [C] 2
func (u Logger) Critical(f interface{}, v ...interface{}) {
	logger.Critical(u.formatLog(f, v...))
}

// Error [E] 3
func (u Logger) Error(f interface{}, v ...interface{}) {
	logger.Error(u.formatLog(f, v...))
}

// Warn [W] 4
func (u Logger) Warn(f interface{}, v ...interface{}) {
	logger.Warn(u.formatLog(f, v...))
}

// Notice [N] 5
func (u Logger) Notice(f interface{}, v ...interface{}) {
	logger.Notice(u.formatLog(f, v...))
}

// Info [I] 6
func (u Logger) Info(f interface{}, v ...interface{}) {
	logger.Info(u.formatLog(f, v...))
}

// Debug [D] 7
func (u Logger) Debug(f interface{}, v ...interface{}) {
	logger.Debug(u.formatLog(f, v...))
}

// formatLog format log
func (u Logger) formatLog(f interface{}, v ...interface{}) string {
	str := fmt.Sprintf("[%v] [%v] %v", u.funcCall(), u.traceID, u.format(f, v...))
	str = strings.Replace(str, "\r\n", " ", -1)
	str = strings.Replace(str, "\n", " ", -1)
	str = strings.Replace(str, "\t", "", -1)
	return str
}

// funcCall return function call stack
func (u Logger) funcCall() string {
	_, file, line, ok := runtime.Caller(u.funcCallDepth)
	if !ok {
		file = "???"
		line = 0
	}
	_, filename := path.Split(file)
	return filename + ":" + strconv.Itoa(line)
}

// format row
func (u Logger) format(f interface{}, v ...interface{}) string {
	var msg string
	switch f.(type) {
	case string:
		msg = f.(string)
		if len(v) == 0 {
			return msg
		}
		if strings.Contains(msg, "%") && !strings.Contains(msg, "%%") {
		} else {
			msg += strings.Repeat(" %v", len(v))
		}
	default:
		msg = fmt.Sprint(f)
		if len(v) == 0 {
			return msg
		}
		msg += strings.Repeat(" %v", len(v))
	}
	return fmt.Sprintf(msg, v...)
}
