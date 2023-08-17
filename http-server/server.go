package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	address    = "127.0.0.1"
	port       = "8000"
	defaultMsg = "Using default address value"
)

func init() {
	envReading()
}

func Start() {
	// Initialize context and cancel function
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	// Trap Ctrl+C and call cancel on the context
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	defer func() {
		signal.Stop(signals)
		cancel()
	}()

	addrPort := address + ":" + port
	server := http.Server{
		Addr:    addrPort,
		Handler: nil,
	}

	go func() {
		select {
		case <-signals:
			cancel()
		case <-ctx.Done():
		}
	}()

	go func() {
		defineHandler(&server)
		startServer(&server, ctx)
	}()

	<-ctx.Done()

	// Additional cleanup or shutdown logic can be placed here
	log.Println("Server shutting down...")
}

func startServer(server *http.Server, ctx context.Context) {
	log.Println("Server is running on", server.Addr)

	defer stopServer(server, ctx)

	err := server.ListenAndServe()
	if err == nil || err == http.ErrServerClosed {
		log.Println("Server is stopped due to error:", err)
	}
}

func stopServer(server *http.Server, ctx context.Context) {
	err := server.Shutdown(ctx)
	if err != nil {
		log.Println("Error shutting down server:", err)
	}
}
