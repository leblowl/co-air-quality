package app

import (
	"context"
	"co-air-quality.api/src/app/route"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	Handler *mux.Router
	Server *http.Server
}

func (a *App) Init(opts map[string]string) {
	var httpAddress = opts["http-address"]

	if httpAddress == "" {
		httpAddress = ":8080"
	}

	log.Println("Server options:")
	log.Println("  HTTP Address: " + httpAddress)

	a.Handler = route.LoadRoutes()
	a.Server = &http.Server{
		Addr: httpAddress,
		Handler: a.Handler,
	}
}

func (a *App) Start() {
	log.Println("Starting server...")

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := a.Server.ListenAndServe();
			err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Println("Server started.")

	<-done
	log.Println("Stopping server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := a.Server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed:%+v", err)
	}
	log.Print("Server stopped.")
}
