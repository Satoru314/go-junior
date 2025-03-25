package common

import (
	"context"
	"net/http"
)

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

type userNameKey struct{}

func GetuserName(ctx context.Context) string {
	id := ctx.Value(userNameKey{})

	if usernameStr, ok := id.(string); ok {
		return usernameStr
	}

	return ""
}

func SetuserName(req *http.Request, name string) *http.Request {
	ctx := req.Context()
	ctx = context.WithValue(ctx, userNameKey{}, name)
	req = req.WithContext(ctx)
	return req
}
