package logger

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

var _ logrus.Hook = (&TracingHook{})

type TracingHook struct {
}

func (s TracingHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.DebugLevel,
		logrus.ErrorLevel,
		logrus.FatalLevel,
	}
}

func (s TracingHook) Fire(e *logrus.Entry) error {
	span := trace.SpanFromContext(e.Context)
	if !span.IsRecording() {
		return nil
	}
	span.SetAttributes(attribute.KeyValue{
		Key:   attribute.Key(fmt.Sprintf("%d.logrus.%s", time.Now().UnixMicro(), e.Level.String())),
		Value: attribute.StringValue(e.Message),
	})
	return nil
}
