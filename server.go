package main

import (
	"github.com/allanurbayramgeldiyev209/learngin/helpers"
	"github.com/allanurbayramgeldiyev209/learngin/routes"
)

func main() {
	api_routes := routes.ApiRoutes()
	err := api_routes.Run()
	helpers.CheckErr(err)
}
