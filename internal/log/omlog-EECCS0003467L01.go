package log

import (
	"GOKIT_v001/internal/conf"
	"fmt"
	"path/filepath"
	"runtime"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// var Writer *syslog.Writer

// NewZapSugarLogger returns a Go kit log.Logger that sends
// log events to a zap.Logger.

type OMLogger interface {
	Begin(m LogMessage)
	Success(m LogMessage)
	Info(m LogMessage)
	Debug(m LogMessage)
	Warn(m LogMessage)
	Error(m LogMessage)
}

type APITraceLogger struct {
	zlogger zap.Logger
}

var ApiTraceLog OMLogger = NewAPITraceLogger()

func NewAPITraceLogger() APITraceLogger {
	outPath := getOutputPath("log.api.output")
	errPath := getOutputPath("log.api.error")

	cfg := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(zapcore.Level(conf.GetInt("log.api.level.trace"))),
		OutputPaths:      outPath,
		ErrorOutputPaths: errPath,

		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "message",
			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,
			TimeKey:     "time",
			EncodeTime:  zapcore.ISO8601TimeEncoder,
		},
	}
	fmt.Printf("cfg: %v\n", cfg)
	logger, _ := cfg.Build()

	return APITraceLogger{
		zlogger: *logger,
	}
}

func (l APITraceLogger) Begin(m LogMessage) {
	if m.LineOfCode == "" {
		_, filename, line, _ := runtime.Caller(1)
		m.LineOfCode = fmt.Sprintf("%s:%d", filepath.Base(filename), line)
	}
	m.LogTime = time.Now()
	m.setLocalTime()
	m.Step = API_IN
	l.zlogger.Info(m.string())
}

func (l APITraceLogger) Success(m LogMessage) {
	if m.LineOfCode == "" {
		_, filename, line, _ := runtime.Caller(1)
		m.LineOfCode = fmt.Sprintf("%s:%d", filepath.Base(filename), line)
	}
	m.LogTime = time.Now()
	m.setLocalTime()
	m.Step = API_OUT_SUCCESS
	l.zlogger.Info(m.string())
}

func (l APITraceLogger) Info(m LogMessage) {
	if m.LineOfCode == "" {
		_, filename, line, _ := runtime.Caller(1)
		m.LineOfCode = fmt.Sprintf("%s:%d", filepath.Base(filename), line)
	}
	m.LogTime = time.Now()
	m.setLocalTime()
	l.zlogger.Info(m.string())
}

func (l APITraceLogger) Warn(m LogMessage) {
	if m.LineOfCode == "" {
		_, filename, line, _ := runtime.Caller(1)
		m.LineOfCode = fmt.Sprintf("%s:%d", filepath.Base(filename), line)
	}
	m.setTimeMsec()
	l.zlogger.Warn(m.string())
}

func (l APITraceLogger) Error(m LogMessage) {
	if m.LineOfCode == "" {
		_, filename, line, _ := runtime.Caller(1)
		m.LineOfCode = fmt.Sprintf("%s:%d", filepath.Base(filename), line)
	}
	m.setTimeMsec()
	if m.Step <= 0 {
		m.Step = API_OUT_ERROR
	}
	l.zlogger.Error(m.string())
}

func (l APITraceLogger) Debug(m LogMessage) {
	if m.LineOfCode == "" {
		_, filename, line, _ := runtime.Caller(1)
		m.LineOfCode = fmt.Sprintf("%s:%d", filepath.Base(filename), line)
	}
	m.LogTime = time.Now()
	m.setLocalTime()
	l.zlogger.Debug(m.string())
}

// ------------------ APPLICATION TRACELOG ----------------------

type AppTraceLogger struct {
	zlogger zap.Logger
}

var AppTraceLog AppTraceLogger = NewAppTraceLogger()

func NewAppTraceLogger() AppTraceLogger {
	outPath := getOutputPath("log.app.output")
	errPath := getOutputPath("log.app.error")

	cfg := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(zapcore.Level(conf.GetInt("log.app.level.trace"))),
		OutputPaths:      outPath,
		ErrorOutputPaths: errPath,

		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "message",
			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,
			TimeKey:     "time",
			EncodeTime:  zapcore.ISO8601TimeEncoder,
		},
	}
	logger, _ := cfg.Build()

	return AppTraceLogger{
		zlogger: *logger,
	}
}

func (l AppTraceLogger) Info(m LogMessage) {
	m.Step = APP_SUCCESS
	if m.LineOfCode == "" {
		_, filename, line, _ := runtime.Caller(1)
		m.LineOfCode = fmt.Sprintf("%s:%d", filepath.Base(filename), line)
	}
	m.LogTime = time.Now()
	m.setLocalTime()
	m.setTimeMsec()
	l.zlogger.Info(m.toAppString())
}

func (l AppTraceLogger) Warn(m LogMessage) {
	if m.LineOfCode == "" {
		_, filename, line, _ := runtime.Caller(1)
		m.LineOfCode = fmt.Sprintf("%s:%d", filepath.Base(filename), line)
	}
	m.LogTime = time.Now()
	m.setLocalTime()
	m.setTimeMsec()
	l.zlogger.Warn(m.toAppString())
}

func (l AppTraceLogger) Error(m LogMessage) {
	m.Step = APP_ERROR
	if m.LineOfCode == "" {
		_, filename, line, _ := runtime.Caller(1)
		m.LineOfCode = fmt.Sprintf("%s:%d", filepath.Base(filename), line)
	}
	m.LogTime = time.Now()
	m.setLocalTime()
	m.setTimeMsec()
	l.zlogger.Error(m.toAppString())
}

func (l AppTraceLogger) Debug(m LogMessage) {
	if m.LineOfCode == "" {
		_, filename, line, _ := runtime.Caller(1)
		m.LineOfCode = fmt.Sprintf("%s:%d", filepath.Base(filename), line)
	}
	m.LogTime = time.Now()
	m.setLocalTime()
	m.setTimeMsec()
	l.zlogger.Debug(m.toAppString())
}
