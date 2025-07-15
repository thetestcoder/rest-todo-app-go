package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type responseWriterWrapper struct {
	http.ResponseWriter
	buffer []byte
}

type ResponseWrapper struct {
	Data      interface{} `json:"data"`
	Status    string      `json:"status"`
	Error     string      `json:"error" omitempty:"true"`
	Timestamp time.Time   `json:"timestamp"`
}

func (w *responseWriterWrapper) Write(b []byte) (int, error) {
	w.buffer = append(w.buffer, b...)
	return len(b), nil
}

func RestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		wrapper := &responseWriterWrapper{ResponseWriter: writer, buffer: []byte{}}

		writer.Header().Set("Content-Type", "application/json")
		writer.Header().Set("Access-Control-Allow-Origin", "*")

		next.ServeHTTP(wrapper, request)

		if len(wrapper.buffer) > 0 {
			var originalData interface{}

			if err := json.Unmarshal(wrapper.buffer, &originalData); err == nil {
				response := ResponseWrapper{
					Status:    "success",
					Data:      originalData,
					Timestamp: time.Now().UTC(),
				}

				if newData, err := json.Marshal(response); err == nil {
					writer.Write(newData)
					return
				} else {
					log.Println(err)
				}
			}
		}

		writer.Write(wrapper.buffer)
	})
}
