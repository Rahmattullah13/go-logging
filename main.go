package main

import (
	"context"
	"time"

	"github.com/Rahmattullah13/go-logging/example"
	"github.com/Rahmattullah13/go-logging/logger"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func main() {
	base()
	example.StacktraceWithData()
}

func base() {
	log := logger.NewLog()

	log.SetLevel(int(logrus.ErrorLevel))
	log.SetReportCaller(true)
	ctx := context.Background()

	data := map[logger.LogKey]interface{}{
		logger.TRACER_ID:     uuid.New(),
		logger.RESPONSE_TIME: time.Now().Second(),
		logger.RESPONSE_TYPE: "Second",
	}

	ctx = context.WithValue(ctx, logger.DATA, data)

	// this will not see because
	// log level is ERROR
	log.Infof(ctx, "Info %s", "info message")

	// you can see this log in stdout
	log.Errorf(ctx, "Error %s", "error message")

	log.SetLevel(int(logrus.DebugLevel))

	// remove file caller
	log.SetReportCaller(false)
	log.Infof(ctx, "info %s", "info message")
}
