package utils

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

const TraceIDKey = "traceID"

// RequestLogger holds information about the request for logging
type RequestLogger struct {
	StartTime time.Time
	Path      string
	Method    string
	Query     map[string]string
	Params    map[string]string
	Status    int
}

// InitLogger initializes the global logger
func InitLogger() {
	Log = logrus.New()

	// Set JSON formatter for structured logging
	Log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime: "timestamp",
			logrus.FieldKeyMsg:  "message",
		},
	})

	// Initially set output to stdout only
	Log.SetOutput(os.Stdout)

	// Get current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		Log.Warnf("Failed to get working directory: %v", err)
		return
	}

	// Create logs directory in current working directory
	logsDir := filepath.Join(currentDir, "logs")
	if err := os.MkdirAll(logsDir, 0755); err != nil {
		Log.Warnf("Failed to create logs directory: %v", err)
		return
	}

	// Try to create and open the log file
	logFile, err := os.OpenFile(
		filepath.Join(logsDir, "application.log"),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0644,
	)
	if err != nil {
		Log.Warnf("Failed to open log file, continuing with stdout only: %v", err)
		return
	}

	// If we successfully opened the file, use both stdout and file
	Log.SetOutput(io.MultiWriter(os.Stdout, logFile))

	// Add default fields
	Log.AddHook(&ContextHook{})
}

// ContextHook adds file and line number to log entries
type ContextHook struct{}

func (hook *ContextHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook *ContextHook) Fire(entry *logrus.Entry) error {
	// Add environment information
	entry.Data["environment"] = getEnv("APP_ENV", "development")
	entry.Data["service_name"] = getEnv("SERVICE_NAME", "go_app")
	entry.Data["timestamp"] = time.Now().Format(time.RFC3339Nano)

	return nil
}

// LogError logs an error with file, line, and function context
func LogError(ctx context.Context, err error, message string) {
	traceID, _ := ctx.Value(TraceIDKey).(string)

	// Get caller information
	pc, file, line, ok := runtime.Caller(1)
	var funcName string
	if ok {
		funcName = runtime.FuncForPC(pc).Name()
	}

	entry := logrus.Fields{
		"trace_id": traceID,
		"error":    err.Error(),
		"file":     filepath.Base(file),
		"line":     line,
		"function": funcName,
		"type":     "application",
	}

	Log.WithFields(entry).Error(message)
}

// LogRequest logs the complete request cycle
func LogRequest(ctx context.Context, reqLogger *RequestLogger) {
	traceId, _ := ctx.Value(TraceIDKey).(string)

	duration := time.Since(reqLogger.StartTime)
	var durationStr string
	if duration.Milliseconds() < 1 {
		durationStr = fmt.Sprintf("%d µs", duration.Microseconds())
	} else {
		durationStr = fmt.Sprintf("%d ms", duration.Milliseconds())
	}

	statusGroup := fmt.Sprintf("%dxx", reqLogger.Status/100)

	logFields := logrus.Fields{
		"trace_id":     traceId,
		"path":         reqLogger.Path,
		"method":       reqLogger.Method,
		"query_params": reqLogger.Query,
		"path_params":  reqLogger.Params,
		"status":       reqLogger.Status,
		"status_group": statusGroup,
		"duration":     durationStr,
		"type":         "request",
	}

	if reqLogger.Status >= 400 {
		Log.WithFields(logFields).Error("Server error")
	} else {
		Log.WithFields(logFields).Info("Request processed")
	}
}
