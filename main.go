package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Customer service")
	})

	server := http.Server {
		Addr: ":3000",
	}

	go func() {
		log.Fatal(server.ListenAndServeTLS("./cert.pem", "key.pem"))
	}()

	fmt.Println("Server started, press <Enter> to shutdown")
	fmt.Scanln()

	server.Shutdown(context.Background())
	fmt.Println("Server stopped")
}