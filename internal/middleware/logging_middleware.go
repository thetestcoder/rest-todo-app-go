package middleware

import (
	"log"
	"net/http"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", request.Method, request.URL.Path)

		rw := &responseWriter{ResponseWriter: writer, statusCode: http.StatusOK}
		next.ServeHTTP(rw, request)

		log.Printf("Completed %s %s with %d in %v",
			request.Method,
			request.URL.Path,
			rw.statusCode,
			time.Since(start),
		)
	})
}
