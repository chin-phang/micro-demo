package main

import (
	"authentication/controllers"
	"authentication/initializers"
	"authentication/middlewares"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvs()
	initializers.ConnectDB()
}

func main() {
	router := gin.Default()

	router.POST("/auth/signup", controllers.CreateUser)
	router.POST("/auth/login", controllers.Login)
	router.GET("/user/profile", middlewares.CheckAuth, controllers.GetUserProfile)
	router.Run()
}
