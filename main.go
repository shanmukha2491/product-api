package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"product-api/handlers"
	"time"
)

func main() {
	fmt.Println("Good Morning!! What would you like?")
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	productHandler := handlers.NewProduct(l)

	serveMultiplexer := http.NewServeMux()
	serveMultiplexer.Handle("/", productHandler)
	server := &http.Server{
		Addr:         ":9090",
		Handler:      serveMultiplexer,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	l.Println("Server Started Successfully")
	server.ListenAndServe()
}
