package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/m3o/m3o-proxy/client"
)

func main() {
	token := os.Getenv("M3O_API_TOKEN")

	if len(token) == 0 {
		fmt.Println("Missing M3O_API_TOKEN environment variable")
		os.Exit(1)
	}

	// logging handler
	handler := handlers.CombinedLoggingHandler(os.Stdout, client.New(token))
	// register the proxy handler
	http.Handle("/", handler)
	// run on port 8080
	http.ListenAndServe(":8080", nil)
}
