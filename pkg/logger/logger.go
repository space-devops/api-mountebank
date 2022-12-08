package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
)

var myLogger *zap.Logger
var logLevel zap.AtomicLevel

type LogExtraInfo struct {
	Key   string
	Value interface{}
}

func InitLogger(loggerFile string) {
	logLevel = zap.NewAtomicLevel()
	logLevel.SetLevel(zapcore.DebugLevel)

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	fileEncoder := zapcore.NewJSONEncoder(encoderCfg)
	consoleEncoder := zapcore.NewConsoleEncoder(encoderCfg)

	logFile, err := os.OpenFile(loggerFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		msg := fmt.Sprintf("Error while creating logger file: %v", err)
		log.Fatal(msg)
	}

	writer := zapcore.AddSync(logFile)

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writer, logLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), logLevel),
	)
	myLogger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
}

func SetLogLevel(ll string) {
	switch ll {
	case "debug":
		logLevel.SetLevel(zap.DebugLevel)
		break
	case "info":
		logLevel.SetLevel(zap.InfoLevel)
		break
	case "warn":
		logLevel.SetLevel(zap.WarnLevel)
		break
	case "error":
		logLevel.SetLevel(zap.ErrorLevel)
		break
	case "fatal":
		logLevel.SetLevel(zap.FatalLevel)
		break
	case "panic":
		logLevel.SetLevel(zap.PanicLevel)
		break
	default:
		logLevel.SetLevel(zap.DebugLevel)
	}
}

func LogDebug(message string, correlationId string, extraInfo ...LogExtraInfo) {
	logEntry(zapcore.DebugLevel, message, correlationId, extraInfo...)
}

func LogInfo(message string, correlationId string, extraInfo ...LogExtraInfo) {
	logEntry(zapcore.InfoLevel, message, correlationId, extraInfo...)
}

func LogWarn(message string, correlationId string, extraInfo ...LogExtraInfo) {
	logEntry(zapcore.WarnLevel, message, correlationId, extraInfo...)
}

func LogError(message string, correlationId string, extraInfo ...LogExtraInfo) {
	logEntry(zapcore.ErrorLevel, message, correlationId, extraInfo...)
}

func LogFatal(message string, correlationId string, extraInfo ...LogExtraInfo) {
	logEntry(zapcore.FatalLevel, message, correlationId, extraInfo...)
}

func LogPanic(message string, correlationId string, extraInfo ...LogExtraInfo) {
	logEntry(zapcore.PanicLevel, message, correlationId, extraInfo...)
}

func logEntry(lvl zapcore.Level, message string, correlationId string, extraInfo ...LogExtraInfo) {
	defer myLogger.Sync()

	if len(extraInfo) > 0 {
		for _, info := range extraInfo {
			myLogger.Log(lvl, message,
				zap.String("correlationId", correlationId),
				zap.Any(info.Key, info.Value),
			)
		}
	} else {
		myLogger.Log(lvl, message,
			zap.String("correlationId", correlationId),
		)
	}
}
