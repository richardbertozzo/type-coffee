package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/richardbertozzo/type-coffee/coffee/api"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.NewClient("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCoffeeServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	r, err := c.GetBestTypeCoffee(ctx, &pb.GetBestTypeCoffeeRequest{Characteristics: []pb.Characteristic{pb.Characteristic_FLAVOR, pb.Characteristic_BODY}})
	if err != nil {
		log.Fatalf("could not get best coffee: %v", err)
	}
	log.Printf("Result: %s", r.BestCoffee)
}
