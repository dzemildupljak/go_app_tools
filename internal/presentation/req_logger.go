package presentation

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/dzemildupljak/go_app_tools/utils"
)

func httpLoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Generate a trace ID
		traceID := uuid.NewString()

		// Add the trace ID to the request context
		ctx := context.WithValue(r.Context(), utils.TraceIDKey, traceID)

		r = r.WithContext(ctx)

		reqLogger := &utils.RequestLogger{
			StartTime: time.Now(),
			Path:      r.URL.Path,
			Method:    r.Method,
			Query:     make(map[string]string),
			Params:    make(map[string]string),
		}

		// Custom ResponseWriter to capture status code
		rw := &responseWriter{w, http.StatusOK}

		// Process request
		next.ServeHTTP(rw, r)

		// Get path parameters
		vars := mux.Vars(r)
		reqLogger.Params = vars

		// Get query parameters
		query := r.URL.Query()
		for key := range query {
			reqLogger.Query[key] = query.Get(key)
		}

		reqLogger.Status = rw.status
		utils.LogRequest(r.Context(), reqLogger)
	})
}

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}
