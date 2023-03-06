package main

import (
	"context"
	"database/sql"
	"encoding/csv"
	"flag"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/richardbertozzo/type-coffee/coffee/service"
	"github.com/richardbertozzo/type-coffee/infra/database"
)

func main() {
	flagFile := flag.String("FILE", "./data/coffee_data.csv", "Path to the coffee dataset CSV file")
	dbURL := flag.String("DATABASE_URL", "", "Database URL value")

	flag.Parse()

	if dbURL == nil {
		log.Fatal("DATABASE_URL ENV is required")
	}

	dbPool, err := database.NewConnection(context.Background(), *dbURL)
	if err != nil {
		log.Fatal(err)
	}
	queries := service.New(dbPool)

	err = run(queries, *flagFile)
	if err != nil {
		log.Fatal(err)
	}
}

func run(queries *service.Queries, filePath string) error {
	f, csvReader, err := CSVReader(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	isHeaderRow := true
	recordsInserted := 0
	start := time.Now()
	for {
		record, err := csvReader.Read()
		if err != nil {
			// no more rows in the file
			if err == io.EOF {
				break
			}
			return err
		}

		// skip the header row (first row)
		if isHeaderRow {
			isHeaderRow = false
			continue
		}

		coffee := createCoffeeFromCSVRow(record)
		id, err := queries.InsertCoffee(context.Background(), service.InsertCoffeeParams{
			Specie:          coffee.Specie,
			Owner:           coffee.Owner,
			CountryOfOrigin: coffee.CountryOfOrigin,
			Company:         coffee.Company.String,
			Aroma:           coffee.Aroma,
			Flavor:          coffee.Flavor,
			Aftertaste:      coffee.Aftertaste,
			Acidity:         coffee.Acidity,
			Body:            coffee.Body,
			Sweetness:       coffee.Sweetness,
		})
		log.Printf("coffee inserted, id: %s - specie %s", id.String(), coffee.Specie)
		recordsInserted++
	}

	log.Printf("finished processing rows csv file, total: %d - duration %v", recordsInserted, time.Now().Sub(start))
	return nil
}

func parseStrToFloat(strValue string) float32 {
	f, err := strconv.ParseFloat(strValue, 32)
	if err != nil {
		log.Printf("error in parse to float 32: %s", err)
	}

	return float32(f)
}

func createCoffeeFromCSVRow(row []string) service.Coffee {
	return service.Coffee{
		Specie:          row[1],
		Owner:           row[2],
		CountryOfOrigin: row[3],
		Company: sql.NullString{
			String: row[5],
			Valid:  true,
		},
		Aroma:      parseStrToFloat(row[10]),
		Flavor:     parseStrToFloat(row[11]),
		Aftertaste: parseStrToFloat(row[12]),
		Acidity:    parseStrToFloat(row[13]),
		Body:       parseStrToFloat(row[14]),
		Sweetness:  parseStrToFloat(row[19]),
	}
}

func CSVReader(pathFile string) (*os.File, *csv.Reader, error) {
	file, err := os.Open(pathFile)
	if err != nil {
		return nil, nil, err
	}

	return file, csv.NewReader(file), nil
}
