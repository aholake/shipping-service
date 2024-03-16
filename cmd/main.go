package main

import (
	"log"

	"github.com/aholake/shipping-service/config"
	"github.com/aholake/shipping-service/internal/adapters"
	"github.com/aholake/shipping-service/internal/application/core/api"
)

func main() {
	db, err := adapters.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("unable to create db adapter, error: %v", err)
	}
	apiPort := api.NewApplication(db)
	server := adapters.NewServer(int32(config.GetApplicationPort()), apiPort)

	server.Run()
}
