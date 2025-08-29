package mlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *MLogger

func DefaultOptions() *LoggerOptions {
	return &LoggerOptions{
		Level:      "info",
		LogsDir:    "logs",
		MaxSize:    10,
		MaxAge:     5,
		MaxBackups: 5,
		Compress:   false,
		Console:    true,
	}
}

func Init(options *LoggerOptions) error {
	var err error
	logger, err = New(options, zap.AddCallerSkip(2))
	return err
}

func DebugEnable() bool {
	return logger.level >= zapcore.DebugLevel
}

func InfoEnable() bool {
	return logger.level >= zapcore.InfoLevel
}

func ErrorEnable() bool {
	return logger.level >= zapcore.ErrorLevel
}

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}

func Debugf(template string, args ...interface{}) {
	logger.Debugf(template, args...)
}

func Infof(template string, args ...interface{}) {
	logger.Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	logger.Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	logger.Errorf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	logger.Fatalf(template, args...)
}

func GetLogger() *zap.Logger {
	return logger.logger
}
