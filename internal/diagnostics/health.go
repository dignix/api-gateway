package diagnostics

import (
	"log"
	"net/http"
)

func ReadinessHandler(l *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l.Println("Readiness probe")

		w.WriteHeader(http.StatusOK)
	}
}

func LivenessHandler(l *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l.Println("Liveness probe")

		w.WriteHeader(http.StatusOK)
	}
}
