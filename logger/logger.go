package logger

import (
	"context"
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

type Logger interface {
	Infof(ctx context.Context, format string, msg ...interface{})
	Errorf(ctx context.Context, format string, msg ...interface{})
	build(ctx context.Context) *logrus.Entry
	SetLevel(level int)
	SetReportCaller(bool)
}

type log struct {
	logger *logrus.Logger
	report bool
}

func NewLog() Logger {
	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05.999",
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			fileName := path.Base(frame.File) + ":" + strconv.Itoa(frame.Line)
			return "", fileName
		},
	})
	logger.SetLevel(logrus.DebugLevel)
	return &log{
		logger: logger,
	}
}

func (l *log) SetLevel(level int) {
	l.logger.SetLevel(logrus.Level(level))
}

func (l *log) SetReportCaller(report bool) {
	l.report = report
}

func (l *log) build(ctx context.Context) *logrus.Entry {
	var fields logrus.Fields = generateFields(ctx, l.report)
	entry := l.logger.WithFields(fields)
	return entry
}

func (l *log) Infof(ctx context.Context, format string, msg ...interface{}) {
	l.build(ctx).Infof(format, msg...)
}
func (l *log) Errorf(ctx context.Context, format string, msg ...interface{}) {
	l.build(ctx).Errorf(format, msg...)
}

func generateFields(ctx context.Context, report bool) map[string]interface{} {
	val := ctx.Value(DATA)
	var data map[string]interface{} = map[string]interface{}{}

	if report {
		// get current directory
		path, _ := os.Getwd()

		// get file and line caller
		// caller is referrence to stack.frame
		caller := 3

		// get file and line
		_, file, line, _ := runtime.Caller(caller)

		// get current file
		myFile := strings.Split(file, path)[1]

		// add new field to handle file
		data["file"] = fmt.Sprintf("%s.%v", myFile, line)
	}

	if val == nil {
		return data
	}
	for k, v := range val.(map[LogKey]interface{}) {
		data[strings.ToLower(string(k))] = v
	}

	return data
}
