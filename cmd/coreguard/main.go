package main

import (
	"log"
	"net/http"
	"os"

	"github.com/wyw14/cry-136/internal/api"
	"github.com/wyw14/cry-136/internal/service"
)

func main() {
	address := os.Getenv("COREGUARD_ADDR")
	if address == "" {
		address = "127.0.0.1:21236"
	}
	runtime := service.NewRuntime(os.Getenv("COREGUARD_DATA"))
	server := &http.Server{Addr: address, Handler: api.NewRouter(runtime)}
	log.Printf("coreguard listening on %s", address)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
