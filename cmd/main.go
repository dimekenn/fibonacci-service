package main

import (
	"context"
	"embed"
	"encoding/json"
	"fibonacciService/configs"
	"fibonacciService/internal/app"
	"log"
)

//go:embed configs.json
var fs embed.FS

const configName = "configs.json"

func main()  {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	//reading json file for configs
	data, readErr := fs.ReadFile(configName)
	if readErr != nil {
		log.Fatal(readErr)
	}

	//creating config entity to deserialize configs.json
	cfg := configs.NewConfig()
	if unmErr := json.Unmarshal(data, &cfg); unmErr != nil {
		log.Fatal(unmErr)
	}

	//channel for errors
	errCh := make(chan error, 1)

	//new goroutine for REST api server
	go app.StartHTTPServer(ctx,errCh, cfg)
	//new goroutine for GRPC server
	go app.StartGRPCServer(ctx,errCh, cfg)

	log.Fatalf("%v", <-errCh)
}