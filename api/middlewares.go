package api

import (
	"log/slog"
	"net/http"
)

func requestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Received request",
			slog.String("ip", r.RemoteAddr),
			slog.String("proto", r.Proto),
			slog.String("method", r.Method),
			slog.String("uri", r.RequestURI),
		)
		next.ServeHTTP(w, r)
	})

}

func recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				if err, ok := err.(error); ok {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
			}
		}()

		next.ServeHTTP(w, r)
	})
}
