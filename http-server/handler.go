package server

import (
	"log"
	"net/http"
)

func defineHandler(server *http.Server) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", baseHandler)
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/hi", hiHandler)

	server.Handler = mux
}

func baseHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("🦫 CAPY")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello World")
	w.Write([]byte("Hello, World!"))
}

func hiHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Hi There")
	w.Write([]byte("Hi there!"))
}

func hitServerHandler() http.Handler {
	counter := 0
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			counter++
			log.Println("Server is getting hit!", counter)
		},
	)
}
