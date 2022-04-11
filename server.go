package main

import (
	"github.com/allanurbayramgeldiyev209/learngin/helpers"
	"github.com/allanurbayramgeldiyev209/learngin/routes"
)

func main() {

	api_routes := routes.ApiRoutes()
	err := api_routes.Run(":0704")
	helpers.CheckErr(err)
}
