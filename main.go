package main

import (
	"log"

	"github.com/NajmiddinAbdulhakim/iman/api-gateway/api"
	"github.com/NajmiddinAbdulhakim/iman/api-gateway/config"
	"github.com/NajmiddinAbdulhakim/iman/api-gateway/service"
)

func main() {
	cfg := config.Load()

	serviceManager, err := service.NewServiceManager(cfg)
	if err != nil {
		log.Fatal(`gRPC dial error: `, err)
	}

	server := api.New(&api.Option{
		Conf: cfg,
		ServiceManager: serviceManager,
	})

	if err := server.Run(cfg.HTTPPort); err != nil {
		log.Fatal(`failed to run http server`, err)
		panic(err)
	}


}
