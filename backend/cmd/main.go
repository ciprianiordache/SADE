package main

import (
	_ "embed"
	"sade-backend/api/routes"
	"sade-backend/api/server"
)

func main() {
	routes.AuthRoute(router)
	routes.MediaRoute(router)
	routes.PaymentRoute(router)
	err := server.RunServer(srv)
	if err != nil {
		log.Fatal(err)
	}
}
