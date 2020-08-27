package main

import (
	"os"
	"user-identity-subdomain/controllers/users"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.POST("/v1/identity/login", users.Login)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)

}
