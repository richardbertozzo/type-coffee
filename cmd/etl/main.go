package main

import (
	"context"
	"encoding/csv"
	"errors"
	"flag"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
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

	run(ctx, queries, *flagFile)
}

func run(ctx context.Context, queries *provider.Queries, filePath string) {
	f, reader, err := csvReader(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = f.Close()
	}()

	isHeaderRow := true
	var recordsInserted atomic.Uint32
	start := time.Now()

	const numWorkers = 10
	done := make(chan bool)
	jobs := make(chan job, numWorkers)
	results := make(chan result, numWorkers)

	go func() {
		var wg sync.WaitGroup
		wg.Add(numWorkers)
		for w := 1; w <= numWorkers; w++ {
			go workerInsert(ctx, &wg, queries, jobs, results)
		}
		wg.Wait()
		close(results)
	}()

	go func() {
		for {
			record, err := reader.Read()
			if err != nil {
				// no more rows in the file
				if errors.Is(err, io.EOF) {
					break
				}

				log.Fatalf("error reading record row: %v", err)
				return
			}

			// skip the header row (first row)
			if isHeaderRow {
				isHeaderRow = false
				continue
			}

			jobs <- job{
				record: record,
			}
		}

		close(jobs)
	}()

	go func() {
		for r := range results {
			recordsInserted.Add(1)
			log.Printf("coffee inserted, id: %s - specie %s", r.ID, r.Specie)
		}
		done <- true
	}()

	<-done
	close(done)
	log.Printf("finished processing rows csv file, total: %d - duration %v", recordsInserted.Load(), time.Now().Sub(start).Minutes())
}

type job struct {
	id     string
	record []string
}

type result struct {
	ID     string
	Specie string
}

func workerInsert(ctx context.Context, wg *sync.WaitGroup, queries *provider.Queries, jobChan chan job, results chan<- result) {
	for j := range jobChan {
		coffee := createCoffeeFromCSVRow(j.record)
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

		u, err := uuid.FromBytes(id.Bytes[:])
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("coffee inserted, id: %s - specie %s", u.String(), coffee.Specie)

		results <- result{
			ID:     u.String(),
			Specie: coffee.Specie,
		}
	}

	wg.Done()
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
