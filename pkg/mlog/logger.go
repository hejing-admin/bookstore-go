package mlog

import (
	"io"
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type MLogger struct {
	options *LoggerOptions
	logger  *zap.Logger
	level   zapcore.Level
	sugar   *zap.SugaredLogger
}

func New(options *LoggerOptions, zapOptions ...zap.Option) (*MLogger, error) {
	intLevel, err := zapcore.ParseLevel(options.Level)
	if err != nil {
		return nil, err
	}

	mLogger := &MLogger{
		options: options,
		level:   intLevel,
	}

	encoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	debugFunc := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zapcore.DebugLevel && level >= mLogger.level
	})

	infoFunc := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zapcore.InfoLevel && level >= mLogger.level
	})

	errorFunc := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.WarnLevel && level >= mLogger.level
	})

	consoleFunc := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= mLogger.level
	})

	debugWriter := getWriter(options.LogsDir+"/debug.log", options)
	infoWriter := getWriter(options.LogsDir+"/info.log", options)
	errorWriter := getWriter(options.LogsDir+"/error.log", options)

	cores := []zapcore.Core{
		zapcore.NewCore(encoder, zapcore.AddSync(debugWriter), debugFunc),
		zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoFunc),
		zapcore.NewCore(encoder, zapcore.AddSync(errorWriter), errorFunc),
	}
	if options.Console {
		cores = append(cores,
			zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), consoleFunc))
	}

	zapOptions = append(zapOptions, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	mLogger.logger = zap.New(zapcore.NewTee(cores...), zapOptions...)
	mLogger.sugar = mLogger.logger.Sugar()

	return mLogger, nil
}

func (m *MLogger) L() *zap.Logger {
	return m.logger
}

func (m *MLogger) setLevel(level string) error {
	intLevel, err := zapcore.ParseLevel(level)
	if err != nil {
		return err
	}

	m.level = intLevel

	return nil
}

func (m *MLogger) DebugEnable() bool {
	return m.level >= zapcore.DebugLevel
}

func (m *MLogger) InfoEnable() bool {
	return m.level >= zapcore.InfoLevel
}

func (m *MLogger) ErrorEnable() bool {
	return m.level >= zapcore.ErrorLevel
}

func (m *MLogger) Debug(msg string, fields ...zap.Field) {
	m.logger.Debug(msg, fields...)
}

func (m *MLogger) Info(msg string, fields ...zap.Field) {
	m.logger.Info(msg, fields...)
}

func (m *MLogger) Warn(msg string, fields ...zap.Field) {
	m.logger.Warn(msg, fields...)
}

func (m *MLogger) Error(msg string, fields ...zap.Field) {
	m.logger.Error(msg, fields...)
}

func (m *MLogger) Fatal(msg string, fields ...zap.Field) {
	m.logger.Fatal(msg, fields...)
}

func (m *MLogger) Debugf(template string, args ...interface{}) {
	m.sugar.Debugf(template, args...)
}

func (m *MLogger) Infof(template string, args ...interface{}) {
	m.sugar.Infof(template, args...)
}

func (m *MLogger) Warnf(template string, args ...interface{}) {
	m.sugar.Warnf(template, args...)
}

func (m *MLogger) Errorf(template string, args ...interface{}) {
	m.sugar.Errorf(template, args...)
}

func (m *MLogger) Fatalf(template string, args ...interface{}) {
	m.sugar.Fatalf(template, args...)
}

func getWriter(filename string, options *LoggerOptions) io.Writer {
	return &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    options.MaxSize, //megabytes
		MaxBackups: options.MaxBackups,
		MaxAge:     options.MaxAge,
		Compress:   options.Compress,
	}
}
