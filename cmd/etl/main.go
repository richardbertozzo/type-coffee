package main

import (
	"context"
	"encoding/csv"
	"flag"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/richardbertozzo/type-coffee/coffee/provider"
	"github.com/richardbertozzo/type-coffee/infra/database"
)

func main() {
	flagFile := flag.String("FILE", "./data/coffee_data.csv", "Path to the coffee dataset CSV file")
	dbURL := flag.String("DATABASE_URL", "", "Database URL value")

	flag.Parse()

	if dbURL == nil {
		log.Fatal("DATABASE_URL ENV is required")
	}

	ctx, cancelFn := context.WithTimeout(context.Background(), time.Minute)
	defer cancelFn()

	dbPool, err := database.NewConnection(ctx, *dbURL)
	if err != nil {
		log.Fatal(err)
	}

	if err = dbPool.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	// open sql table schema file and apply it
	sqlSchema, err := os.ReadFile("./configs/db/schema.sql")
	if err != nil {
		log.Fatalf("error opening SQL schema file: %v", err)
	}
	if _, err = dbPool.Exec(ctx, string(sqlSchema)); err != nil {
		log.Fatal(err)
	}

	queries := provider.New(dbPool)

	err = run(ctx, queries, *flagFile)
	if err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context, queries *provider.Queries, filePath string) error {
	f, reader, err := csvReader(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = f.Close()
	}()

	isHeaderRow := true
	recordsInserted := 0
	start := time.Now()

	for {
		record, err := reader.Read()
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
		id, err := queries.InsertCoffee(ctx, provider.InsertCoffeeParams{
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
		if err != nil {
			log.Fatal(err)
		}

		u, err := uuid.FromBytes(id.Bytes[0:16])
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("coffee inserted, id: %s - specie %s", u.String(), coffee.Specie)
		recordsInserted++
	}

	log.Printf("finished processing rows csv file, total: %d - duration %v", recordsInserted, time.Now().Sub(start).Minutes())
	return nil
}

func parseStrToFloat(strValue string) float32 {
	f, err := strconv.ParseFloat(strValue, 32)
	if err != nil {
		log.Printf("error in parse to float 32: %s", err)
	}

	return float32(f)
}

func createCoffeeFromCSVRow(row []string) provider.Coffee {
	return provider.Coffee{
		Specie:          row[1],
		Owner:           row[2],
		CountryOfOrigin: row[3],
		Company: pgtype.Text{
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

func csvReader(pathFile string) (*os.File, *csv.Reader, error) {
	file, err := os.Open(pathFile)
	if err != nil {
		return nil, nil, err
	}

	return file, csv.NewReader(file), nil
}
