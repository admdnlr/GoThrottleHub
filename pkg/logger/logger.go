package logger

import (
	"log"
	"os"
)

// LogLevel türü, log seviyelerini tanımlar.
type LogLevel int

const (
	// LogLevel tipleri tanımlanır.
	LogLevelDebug LogLevel = iota
	LogLevelInfo
	LogLevelWarning
	LogLevelError
)

// Logger struct'ı, loglama işlemlerinin yapılandırmasını tutar.
type Logger struct {
	level    LogLevel
	debugLog *log.Logger
	infoLog  *log.Logger
	warnLog  *log.Logger
	errorLog *log.Logger
}

// NewLogger, yeni bir Logger örneği oluşturur ve döndürür.
func NewLogger(level LogLevel) *Logger {
	return &Logger{
		level:    level,
		debugLog: log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile),
		infoLog:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime),
		warnLog:  log.New(os.Stderr, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLog: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

// Debug, debug mesajlarını loglar.
func (l *Logger) Debug(v ...interface{}) {
	if l.level <= LogLevelDebug {
		l.debugLog.Println(v...)
	}
}

// Info, info mesajlarını loglar.
func (l *Logger) Info(v ...interface{}) {
	if l.level <= LogLevelInfo {
		l.infoLog.Println(v...)
	}
}

// Warn, warn mesajlarını loglar.
func (l *Logger) Warn(v ...interface{}) {
	if l.level <= LogLevelWarning {
		l.warnLog.Println(v...)
	}
}

// Error, error mesajlarını loglar.
func (l *Logger) Error(v ...interface{}) {
	if l.level <= LogLevelError {
		l.errorLog.Println(v...)
	}
}
