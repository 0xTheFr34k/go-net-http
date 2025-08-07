package middleware

import (
	"log"
	"net/http"
	"runtime/debug"
	"sync"
	"time"
)

var Wg sync.WaitGroup

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func WithWaitGroup(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Wg.Add(1)
		defer Wg.Done()
		h.ServeHTTP(w, r)
	})
}

func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

type Middleware func(http.Handler) http.Handler

func CreateStack(xs ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for _, x := range xs {
			next = x(next)
		}
		return next
	}
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("AuthMiddleware called")
		next.ServeHTTP(w, r)
	})
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		originalURI := r.RequestURI

		wrapped := &wrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapped, r)

		log.Println(r.Method, wrapped.statusCode, originalURI, time.Since(start))
	})
}

func RecoveryMiddlerware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				msg := "Caught panic: %v, Stack trace %s"
				log.Printf(msg, err, string(debug.Stack()))
				http.Error(w, "", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
