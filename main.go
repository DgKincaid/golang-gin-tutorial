package main

import (
	"net/http"

	"github.com/dgkincaid/golang-gin-tutorial/models"
	"github.com/dgkincaid/golang-gin-tutorial/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// models.ConnectDB()
	models.ConnectMongoDB()

	routes.Routes(router)

	router.LoadHTMLGlob("views/templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Home Page"})
	})

	router.Run()
}
