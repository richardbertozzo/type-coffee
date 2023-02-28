package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type handlers struct {
	Data []coffee
}

func main() {
	fmt.Println("starting the mock coffee API")

	records, err := openAndReadCsv("external/mock-api-coffee/data/coffee_data.csv")
	if err != nil {
		panic(err)
	}
	coffees := convertCoffeeRecord(records)

	h := handlers{
		Data: coffees,
	}
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/coffee", h.handlerGetCoffee)

	http.ListenAndServe(":3333", r)
}

func openAndReadCsv(fileName string) ([][]string, error) {
	csvFile, err := os.Open(fileName)
	defer csvFile.Close()

	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(csvFile)
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func convertCoffeeRecord(records [][]string) []coffee {
	var coffees []coffee

	for line, row := range records {
		if line > 0 { // omit header line
			id, _ := strconv.Atoi(row[0])

			c := coffee{
				ID:      id,
				Specie:  row[1],
				Owner:   row[2],
				Country: row[3],
			}
			coffees = append(coffees, c)
		}
	}

	return coffees
}

type coffee struct {
	ID      int    `json:"id"`
	Specie  string `json:"specie"`
	Owner   string `json:"owner,omitempty"`
	Country string `json:"country"`
}

func (h *handlers) handlerGetCoffee(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(h.Data)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
