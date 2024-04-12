package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.Handle("/", customHandler("Customer service"))

	var handlerFunc http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.URL.String())
	}

	http.HandleFunc("/url", handlerFunc)

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintln(w, "Customer service")
	// })

	server := http.Server {
		Addr: ":3000",
	}

	go func() {
		log.Fatal(server.ListenAndServeTLS("./http_server/cert.pem", "./http_server/key.pem"))
	}()

	fmt.Println("Server started, press <Enter> to shutdown")
	fmt.Scanln()

	server.Shutdown(context.Background())
	fmt.Println("Server stopped")
}

type customHandler string

func (cs customHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-Powered-By", "gophers")
	http.SetCookie(w, &http.Cookie{
		Name: "session-id",
		Value: "12345",
		Expires: time.Now().Add(24 * time.Hour * 365), // expire in a year
	})
	w.WriteHeader(http.StatusAccepted)

	fmt.Fprintln(w, string(cs))
	fmt.Fprintln(w, r.Header)
}