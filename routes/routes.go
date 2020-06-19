package routes

import (
	"github.com/dgkincaid/golang-gin-tutorial/controllers"
	"github.com/gin-gonic/gin"
)

// Routes register application routes
func Routes(route *gin.Engine) {

	books := route.Group("/books")
	{
		books.GET("", controllers.FindBooks)
		books.GET("/:id", controllers.FindBook)
		books.PATCH("/:id", controllers.UpdateBook)
		books.POST("", controllers.CreateBook)
		books.DELETE("/:id", controllers.DeleteBook)
	}
}
