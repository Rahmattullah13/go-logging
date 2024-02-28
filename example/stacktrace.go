package example

import (
	"context"

	"github.com/Rahmattullah13/go-logging/logger"
	"github.com/sirupsen/logrus"
)

func Stacktrace() {
	log := logger.NewLog()
	log.SetLevel(int(logrus.DebugLevel))

	// set report caller to enable stack trace
	log.SetReportCaller(true)

	ctx := context.Background()
	log.Infof(ctx, "info")
	getData(ctx, log)
}

func getData(ctx context.Context, log logger.Logger) {
	log.Errorf(ctx, "error")
}
