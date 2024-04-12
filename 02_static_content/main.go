package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {	
	http.HandleFunc("/fprint", func(w http.ResponseWriter, r *http.Request) {
		customerFile, err := os.Open("./customer.csv")
		if err != nil {
			log.Fatal(err)
		}

		defer customerFile.Close()

		data, err := io.ReadAll(customerFile)
		if err != nil {
			log.Fatal(err)
		}

		// csv only contains string data, so safe to transform all byte[] to string
		fmt.Fprint(w, string(data))
		
	})
	
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
