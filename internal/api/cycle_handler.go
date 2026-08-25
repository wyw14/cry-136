package api

import (
	"net/http"

	"github.com/wyw14/cry-136/internal/service"
)

func cycleStart(runtime *service.Runtime) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, runtime.StartCycle())
	}
}

func cycleScram(runtime *service.Runtime) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, runtime.TriggerScram("operator request"))
	}
}
