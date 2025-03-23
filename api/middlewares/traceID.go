package middlewares

import (
	"context"
	"sync"
)

var (
	logNo int = 1
	mu    sync.Mutex
)

func newTraceID() int {
	var no int
	mu.Lock()
	no = logNo
	logNo++
	mu.Unlock()
	return no
}

type traceIDtype string

const traceIDKey traceIDtype = "traceID"

func SetTraceID(ctx context.Context, traceID int) context.Context {
	return context.WithValue(ctx, traceIDKey, traceID)
}

func GetTraceID(ctx context.Context) int {
	traceID, ok := ctx.Value(traceIDKey).(int)
	if !ok {
		return 0
	}
	return traceID
}
