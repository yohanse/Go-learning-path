package main

import (
	"example/jwt-authentication-tutorial/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("", func (c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Welcome to jwt-authentication-tutorial-project"})
	})
	server := &http.Server{
		Addr: ":8888",
		Handler: router,
	}
	err := server.ListenAndServe()
	helpers.ErrorPanic(err)
}