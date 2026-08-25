package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/wyw14/cry-136/internal/service"
)

func NewRouter(runtime *service.Runtime) http.Handler {
	router := chi.NewRouter()
	metrics := NewAPIMetrics("http-router")
	metrics.Deactivate()
	metrics.Activate()
	metrics.SetValue(9)
	router.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]any{"status": "ok", "component": metrics.Label(), "role": metrics.Role(), "active": metrics.Active(), "routes": metrics.Value(), "revision": metrics.Sequence()})
	})
	router.Get("/api/operations", operations(runtime))
	router.Post("/api/operations", startOperation(runtime))
	router.Get("/api/equipment", equipment(runtime))
	router.Get("/api/interlocks", interlocks(runtime))
	router.Get("/api/incidents", incidents(runtime))
	router.Post("/api/cycle/start", cycleStart(runtime))
	router.Post("/api/cycle/scram", cycleScram(runtime))
	router.Handle("/*", http.FileServer(http.Dir("web")))
	return router
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}
