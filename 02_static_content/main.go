package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {	
	http.HandleFunc("/fprint", func(w http.ResponseWriter, r *http.Request) {
		customerFile, err := os.Open("./02_static_content/customer.csv")
		if err != nil {
			log.Fatal(err)
		}

		defer customerFile.Close()

		// transforms byte array to string data
		// data, err := io.ReadAll(customerFile)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// // csv only contains string data, so safe to transform all byte[] to string
		// fmt.Fprint(w, string(data))

		// stream byte array to http
		io.Copy(w, customerFile);
		
	})

	http.HandleFunc("/serveFile", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./02_static_content/customer.csv")
	})

	http.HandleFunc("/serveContent", func(w http.ResponseWriter, r *http.Request) {
		customerFile, err := os.Open("./02_static_content/customer.csv")
		if err != nil {
			log.Fatal(err)
		}

		defer customerFile.Close()

		 http.ServeContent(w, r, "customerdata.csv", time.Now(), customerFile)
	})
	
	http.Handle(
		"/files/", 
		http.StripPrefix(
			"/files/", 
			http.FileServer(http.Dir("./02_static_content")),
		),
	)

	server := http.Server {
		Addr: ":3000",
	}

	go func() {
		log.Fatal(server.ListenAndServe())
	}()

	fmt.Println("Server started, press <Enter> to shutdown")
	fmt.Scanln()

	server.Shutdown(context.Background())
	fmt.Println("Server stopped")
}
