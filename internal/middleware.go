package golog

import (
	"fmt"
	"net/http"
	"time"
)

type CustomResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewCustomResponseWriter(w http.ResponseWriter) *CustomResponseWriter {
	return &CustomResponseWriter{
		ResponseWriter: w,
		statusCode:     200,
	}
}

func (c *CustomResponseWriter) WriteHeader(statusCode int) {
	c.statusCode = statusCode
	c.ResponseWriter.WriteHeader(statusCode)
}
func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		cw := NewCustomResponseWriter(w)
		next.ServeHTTP(cw, r)

		duration := time.Since(start)
		durationStr := formatDuration(duration)

		fmt.Printf("%s[%s]%s %s%s%s %s%s%s %s%d%s %s%s%s\n",
			Gray, time.Now().Format("15:04:05"), Reset,
			getMethodColor(r.Method), r.Method, Reset,
			Bold, r.URL.Path, Reset,
			getStatusColor(cw.statusCode), cw.statusCode, Reset,
			getSpeedColor(duration), durationStr, Reset,
		)
	})
}
