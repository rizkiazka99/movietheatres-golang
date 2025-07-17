package router

import (
	"movietheatres-go/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/add_theatre", controllers.PostMovieTheatre)

	return router
}
