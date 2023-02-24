package logger

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type contextKey struct {
	name string
}

type MiddleWare struct {
	logger *zap.Logger
}

type logContext struct {
	Logger *zap.Logger
	RequestID string
}

var loggerCtxKey = contextKey {
	name: "loggerContext",
}

func (m *MiddleWare) GetMiddleWare() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := uuid.New().String()
			logger := m.logger.With(zap.String("requestID", requestID))

			ctx := context.WithValue(r.Context(), loggerCtxKey, &logContext{
				Logger: logger,
				RequestID: requestID,
			})
			
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func NewMiddleWare(logger *zap.Logger) *MiddleWare {
	return &MiddleWare{
		logger: logger,
	}
}

func ForContext(ctx context.Context) *zap.Logger {
	logContext, _ := ctx.Value(loggerCtxKey).(*logContext)
	return logContext.Logger
}
