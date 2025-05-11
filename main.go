package main

import (
	"github.com/gin-gonic/gin"

	"example.com/todogo/config"
	"example.com/todogo/routes"
)

func init() {
	config.LoadEnvironment()
	config.ConnectToPostgres()
}

func main() {
	r := gin.Default()

	routes.MainRoutes(r)

	r.Run()
}
