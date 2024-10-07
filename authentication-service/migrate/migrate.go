package main

import (
	"authentication/initializers"
	"authentication/models"
)

func init() {
	initializers.LoadEnvs()
	initializers.ConnectDB()
}

func main() {

	initializers.DB.AutoMigrate(&models.User{})
}
