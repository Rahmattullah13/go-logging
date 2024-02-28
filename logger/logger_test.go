package logger

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestGenerateFields(t *testing.T) {
	ctx := context.Background()

	data := map[LogKey]interface{}{
		TRACER_ID:     uuid.New(),
		RESPONSE_TIME: time.Now().Second(),
		RESPONSE_TYPE: "Second",
	}

	ctx = context.WithValue(ctx, DATA, data)

	v := generateFields(ctx, false)

	fmt.Println(v)
}
