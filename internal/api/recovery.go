package api

import (
	"net/http"

	"github.com/wyw14/cry-136/internal/service"
)

func operations(runtime *service.Runtime) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) { writeJSON(w, http.StatusOK, runtime.Operations()) }
}

func equipment(runtime *service.Runtime) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) { writeJSON(w, http.StatusOK, runtime.Equipment()) }
}

func interlocks(runtime *service.Runtime) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) { writeJSON(w, http.StatusOK, runtime.Interlocks()) }
}

func incidents(runtime *service.Runtime) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) { writeJSON(w, http.StatusOK, runtime.Incidents()) }
}

func startOperation(runtime *service.Runtime) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) { writeJSON(w, http.StatusCreated, runtime.CreateOperation("operator-cycle")) }
}
