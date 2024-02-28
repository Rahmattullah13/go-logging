package example

import (
	"context"
	"time"

	"github.com/Rahmattullah13/go-logging/logger"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func StacktraceWithData() {
	log := logger.NewLog()
	log.SetLevel(int(logrus.DebugLevel))

	// set report caller to enable stack trace
	log.SetReportCaller(true)

	ctx := context.Background()

	data := map[logger.LogKey]interface{}{
		logger.TRACER_ID:              uuid.New(),
		logger.RESPONSE_TIME:          time.Now().Second(),
		logger.RESPONSE_TYPE:          "Second",
		logger.LogKey("custom_field"): "hello_world!",
	}

	ctx = context.WithValue(ctx, logger.DATA, data)

	log.Infof(ctx, "info")
	logError(ctx, log)
}

func logError(ctx context.Context, log logger.Logger) {
	log.Errorf(ctx, "error")
}
