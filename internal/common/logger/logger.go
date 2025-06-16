// logger/ctx_logrus.go
package logger

import (
	"context"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"go.temporal.io/sdk/activity"
)

type ctxKey string

const (
	// context keys for storing values
	requestIDKey      ctxKey = "logger.request_id"
	idempotencyKeyKey ctxKey = "logger.idempotency_key"
)

// Option to override default behavior
type Option func(o *options)

type options struct {
	idempKey string
}

// WithIdempotencyKey allows explicitly setting an idempotency key
func WithIdempotencyKey(key string) Option {
	return func(o *options) {
		o.idempKey = key
	}
}

// WithRequestAndIdempHTTP extracts or generates IDs in an HTTP/Gin context
// and returns a standard context.Context carrying both values.
func WithRequestAndIdempHTTP(c *gin.Context, opts ...Option) context.Context {
	// apply options
	o := &options{}
	for _, opt := range opts {
		opt(o)
	}

	// base context
	ctx := c.Request.Context()

	// 1) requestID from Gin middleware header or generate new
	reqID := requestid.Get(c)
	if reqID == "" {
		reqID = uuid.NewString()
	}

	// 2) idempotencyKey: use supplied, else generate new
	idem := o.idempKey
	if idem == "" {
		idem = uuid.NewString()
	}

	// inject into context
	ctx = context.WithValue(ctx, requestIDKey, reqID)
	ctx = context.WithValue(ctx, idempotencyKeyKey, idem)
	return ctx
}

// WithRequestAndIdempActivity extracts or generates IDs in a Temporal activity context
// and returns a context.Context carrying both values.
func WithRequestAndIdempActivity(ctx context.Context, opts ...Option) context.Context {
	// apply options
	o := &options{}
	for _, opt := range opts {
		opt(o)
	}

	// 1) requestID from workflow ID or generate new
	var reqID string
	if info := activity.GetInfo(ctx); info.WorkflowExecution.ID != "" {
		reqID = info.WorkflowExecution.ID
	}
	if reqID == "" {
		reqID = uuid.NewString()
	}

	// 2) idempotencyKey: use supplied, else from RunID, else generate new
	idem := o.idempKey
	if idem == "" && activity.GetInfo(ctx).WorkflowExecution.RunID != "" {
		idem = activity.GetInfo(ctx).WorkflowExecution.RunID
	}
	if idem == "" {
		idem = uuid.NewString()
	}

	// inject into context
	ctx = context.WithValue(ctx, requestIDKey, reqID)
	ctx = context.WithValue(ctx, idempotencyKeyKey, idem)
	return ctx
}

// LoggerFromContext returns a *logrus.Entry pre-populated with request_id and idempotency_key
func LoggerFromContext(ctx context.Context) *logrus.Entry {
	reqID, _ := ctx.Value(requestIDKey).(string)
	idem, _ := ctx.Value(idempotencyKeyKey).(string)
	return logrus.WithFields(logrus.Fields{
		"request_id":      reqID,
		"idempotency_key": idem,
	})
}

// RequestIDFrom returns the request_id stored in context
func RequestIDFrom(ctx context.Context) string {
	if v, _ := ctx.Value(requestIDKey).(string); v != "" {
		return v
	}
	return ""
}

// IdempotencyKeyFrom returns the idempotency_key stored in context
func IdempotencyKeyFrom(ctx context.Context) string {
	if v, _ := ctx.Value(idempotencyKeyKey).(string); v != "" {
		return v
	}
	return ""
}
