package main

import (
	"github.com/GuilhermeJuli/Go.ApiRest.Gin/api_rest_gin_go/database"
	"github.com/GuilhermeJuli/Go.ApiRest.Gin/api_rest_gin_go/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
