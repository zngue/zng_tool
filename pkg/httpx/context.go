package httpx

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/zngue/zng_app/app/server/middleware"
)

const (
	RequestIDKey = "X-Request-Id"
	OperationKey = "X-Request-Operation"
	PathKey      = "X-Request-Path"
)

type (
	requestIDKey struct{}
	operationKey struct{}
	pathKey      struct{}
	startTimeKey struct{}
)

func NewServerContext(ctx context.Context, w http.ResponseWriter, r *http.Request, operation, path string) context.Context {
	requestID := r.Header.Get(RequestIDKey)
	if requestID == "" {
		requestID = uuid.NewString()
	}
	ctx = context.WithValue(ctx, requestIDKey{}, requestID)
	ctx = context.WithValue(ctx, operationKey{}, operation)
	ctx = context.WithValue(ctx, pathKey{}, path)
	ctx = context.WithValue(ctx, startTimeKey{}, time.Now())

	w.Header().Set(RequestIDKey, requestID)
	w.Header().Set(OperationKey, operation)
	w.Header().Set(PathKey, path)
	return ctx
}

func RequestIDFromContext(ctx context.Context) string {
	v, _ := ctx.Value(requestIDKey{}).(string)
	return v
}

func OperationFromContext(ctx context.Context) string {
	v, _ := ctx.Value(operationKey{}).(string)
	return v
}

func PathFromContext(ctx context.Context) string {
	v, _ := ctx.Value(pathKey{}).(string)
	return v
}

func StartTimeFromContext(ctx context.Context) time.Time {
	v, _ := ctx.Value(startTimeKey{}).(time.Time)
	return v
}

func requestID(r *http.Request) string {
	if r == nil {
		return ""
	}
	return RequestIDFromContext(r.Context())
}

func Handle[R any](ctx context.Context, handler middleware.RequestHandler[R]) (R, error) {
	registry := middleware.GetRegistry()
	if registry == nil {
		return handler(ctx)
	}
	operation := OperationFromContext(ctx)
	wrapped := registry.Build(operation, func(ctx context.Context) (any, error) {
		return handler(ctx)
	})
	out, err := wrapped(ctx)
	if err != nil {
		var zero R
		return zero, err
	}
	return out.(R), nil
}

func After[Req any, Resp any](ctx context.Context, err error, in *Req, rs *Resp) {
	registry := middleware.GetRegistry()
	if registry == nil {
		return
	}
	operation := OperationFromContext(ctx)
	h := registry.BuildAfter(operation, func(ctx context.Context, err error, in any, rs any) {})
	h(ctx, err, in, rs)
}
