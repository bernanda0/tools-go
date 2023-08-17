package server

import (
	"bernanda/learn/calculator"
	"bernanda/learn/hello"
	"fmt"
	"log"
	"net/http"
)

func hitServerHandler() http.Handler {
	counter := 0
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			counter++
			log.Println("Server is getting hit!", counter)
		},
	)
}

func defineHandler(server *http.Server) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", baseHandler)

	//The program
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/calculator", calculatorHandler)

	server.Handler = mux
}

func baseHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("🦫 CAPY")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	opening(hello.RunHello, "Calculator")

}

func calculatorHandler(w http.ResponseWriter, r *http.Request) {
	opening(calculator.RunCalculator, "Calculator")
}

func opening(fun func(), s string) {
	defer closing()
	log.Printf("Running %s Program\n\n", s)
	fun()
}

func closing() {
	fmt.Print("\n")
	log.Println("Done")
}
