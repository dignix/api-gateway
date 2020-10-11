package routes

import (
	"log"

	"github.com/gorilla/mux"
	"github.com/iam-solutions/api-gateway/internal/diagnostics"

	restV1 "github.com/iam-solutions/api-gateway/internal/app/http/rest/v1"
)

func MapURLPathsToHandlers(r *mux.Router, l *log.Logger)  {
	r.HandleFunc("/healthz", diagnostics.LivenessHandler(l)).Methods("GET")
	r.HandleFunc("/readyz", diagnostics.ReadinessHandler(l)).Methods("GET")

	r.HandleFunc("/v1/users", restV1.HandleRequestAndRedirect).Methods("GET")
	r.HandleFunc("/v1/users/{id}", restV1.HandleRequestAndRedirect).Methods("GET")

	//r.HandleFunc("/oauth/client", restV1.HandleRequestAndRedirect).Methods("POST")
	//r.HandleFunc("/oauth/token", restV1.HandleRequestAndRedirect).Methods("POST")
	//r.HandleFunc("/oauth/refresh", restV1.HandleRequestAndRedirect).Methods("POST")
	//
	//r.HandleFunc("/v1/orders", restV1.HandleRequestAndRedirect).Methods("GET", "POST")
	//r.HandleFunc("/v1/orders/{id}", restV1.HandleRequestAndRedirect).Methods("GET", "PATCH", "DELETE")
	//
	//r.HandleFunc("/v1/customers", restV1.HandleRequestAndRedirect).Methods("POST")
	//r.HandleFunc("/v1/customers/{id}", restV1.HandleRequestAndRedirect).Methods("GET", "PATCH", "DELETE")
}
