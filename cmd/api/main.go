package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	openapiMiddleware "github.com/oapi-codegen/nethttp-middleware"
	"github.com/richardbertozzo/type-coffee/internal/config"
	"google.golang.org/grpc"

	"github.com/richardbertozzo/type-coffee/coffee"
	pb "github.com/richardbertozzo/type-coffee/coffee/api"
	"github.com/richardbertozzo/type-coffee/coffee/handler"
	"github.com/richardbertozzo/type-coffee/coffee/provider"
	"github.com/richardbertozzo/type-coffee/coffee/usecase"
	"github.com/richardbertozzo/type-coffee/infra/database"
)

func main() {
	cfg := config.LoadConfig()

	var db coffee.Service
	if cfg.DBCfg.DBUrl != "" {
		log.Println("database mode service enabled")
		dbURL := database.BuildURL(cfg.DBCfg.DBUrl, cfg.DBCfg.DBDatabase, cfg.DBCfg.DBDatabase, cfg.DBCfg.DBPassword)

		dbPool, err := database.NewConnection(context.Background(), dbURL)
		if err != nil {
			log.Fatal(err)
		}
		db = provider.NewDatabase(dbPool)
	}

	geminiCli, err := provider.NewGeminiClient(cfg.GeminiAPIKey)
	if err != nil {
		log.Fatal(err)
	}

	uc := usecase.New(geminiCli, db)

	if cfg.GRPCEnabled {
		grpcH := handler.NewGrpcHandler(uc)

		lis, err := net.Listen("tcp", cfg.GRPCPort)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		var opts []grpc.ServerOption
		s := grpc.NewServer(opts...)
		pb.RegisterCoffeeServiceServer(s, grpcH)

		log.Printf("gRPC server listening on %v \n", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve gRPC: %v", err)
		}
	} else {
		swagger, err := coffee.GetSwagger()
		if err != nil {
			log.Fatal(err)
		}

		swagger.Servers = nil

		httpH := handler.NewHTTPHandler(uc)

		r := chi.NewRouter()
		r.Use(middleware.Logger)
		r.Use(openapiMiddleware.OapiRequestValidatorWithOptions(
			swagger,
			&openapiMiddleware.Options{
				Options: openapi3filter.Options{
					ExcludeRequestBody:    true,
					ExcludeResponseBody:   true,
					IncludeResponseStatus: true,
					MultiError:            false,
				},
			},
		))

		coffee.HandlerFromMux(httpH, r)

		log.Printf("http server runs on %s \n", cfg.Port)
		err = http.ListenAndServe(cfg.Port, r)
		if err != nil {
			log.Fatalf("error on start http server %v", err)
		}
	}
}
