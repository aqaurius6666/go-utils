package logger

import (
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func InitLogger(appCtx *cli.Context) *logrus.Logger {
	logger := logrus.New()
	host, _ := os.Hostname()
	logLevel := appCtx.String("log-level")
	level, err := logrus.ParseLevel(logLevel)
	logger.Out = os.Stdout
	if logPath := appCtx.String("log-file-path"); logPath != "" {
		file, err := os.Create(logPath)
		if err != nil {
			logger.Fatalf("File path: %s", logPath)
			return logger
		}
		logger.Out = file
	}

	if err != nil {
		logger.Fatalf("Unknown log-level type: %s", logLevel)
		return logger
	}
	logger.Level = level
	logFormat := appCtx.String("log-format")
	switch strings.ToLower(logFormat) {
	case "json":
		logger.Formatter = &logrus.JSONFormatter{
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime:  "timestamp",
				logrus.FieldKeyLevel: "severity",
				logrus.FieldKeyMsg:   "message",
				"version":            appCtx.String("runtime-version"),
				"host":               host,
			},
			TimestampFormat: time.RFC3339Nano,
		}
	case "plain":
		logger.Formatter = &logrus.TextFormatter{}
	default:
		logger.Fatalf("Unknown log-format type: %s", logFormat)
		return logger
	}

	if appCtx.String("otel-address") != "" {
		logger.AddHook(&TracingHook{})
	}
	return logger
}

func InitLoggerWithoutCLIContext() *logrus.Logger {
	logger := logrus.New()
	host, _ := os.Hostname()
	logger.Level = logrus.DebugLevel
	logger.Formatter = &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "severity",
			logrus.FieldKeyMsg:   "message",
			"version":            "0.0.1-debugger",
			"host":               host,
		},
		TimestampFormat: time.RFC3339Nano,
	}
	logger.Out = os.Stdout
	return logger
}
