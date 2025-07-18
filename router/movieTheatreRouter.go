package router

import (
	"movietheatres-go/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/bioskop", controllers.PostMovieTheatre)
	router.GET("/bioskop", controllers.GetMovieTheatres)
	router.GET("/bioskop/:id", controllers.GetMovieTheatreById)
	router.PUT("/bioskop/:id", controllers.UpdateMovieTheatre)
	router.DELETE("/bioskop/:id", controllers.DeleteMovieTheatre)

	return router
}
