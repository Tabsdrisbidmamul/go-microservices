package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	http.HandleFunc("/customers/add", func(w http.ResponseWriter, r *http.Request) {
		var c Customer
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&c)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		log.Print(c)
	})

	http.HandleFunc("/customers", func(w http.ResponseWriter, r *http.Request) {
		customers, err := readCustomers()
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data, err := json.Marshal(customers)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Add("content-type", "application/json")
		w.Write(data)
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

type Customer struct {
	ID 				int					`json:"id"`
	FirstName string			`json:"firstName"`
	LastName  string			`json:"lastName"`
	Address   string			`json:"address"`
}

func (customer *Customer) convertFromJson(data []byte) (Customer, error) {
	var c Customer
	err := json.Unmarshal(data, &customer)
	return c, err
}


func readCustomers() ([]Customer, error) {
	file, err := os.Open("./customer.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	customers := make([]Customer, 0)
	csvReader := csv.NewReader(file)

	for {
		record, err := csvReader.Read()

		if err == io.EOF {
			return customers, nil
		}

		if err != nil {
			return nil, err
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, err
		}

		customer := Customer {
			ID: id,
			FirstName: record[1],
			LastName: record[2],
			Address: record[3],
		}

		customers = append(customers, customer)
	}
}