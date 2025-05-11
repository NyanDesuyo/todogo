package main

import (
	"example.com/todogo/config"
	"example.com/todogo/models"
)

func init() {
	config.LoadEnvironment()
	config.ConnectToPostgres()
}

func main() {
	config.Database.AutoMigrate(&models.Post{})
	config.Database.AutoMigrate(&models.Todo{})
}
