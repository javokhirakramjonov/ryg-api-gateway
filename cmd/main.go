package main

import (
	"log"
	"ryg-api-gateway/api"
	"ryg-api-gateway/api/handlers"
	"ryg-api-gateway/conf"
)

func main() {
	cnf := conf.LoadConfig()

	cm := handlers.NewRpcClientManager(cnf)

	r := api.NewGinRouter(cm)

	err := r.Run(cnf.ApiGatewayPort)
	if err != nil {
		log.Fatal(err)
	}
}
