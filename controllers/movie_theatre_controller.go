package controllers

import (
	"fmt"
	"movietheatres-go/db"
	"movietheatres-go/models"
	"movietheatres-go/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostMovieTheatre(ctx *gin.Context) {
	var theatre models.MovieTheatre

	if err := ctx.ShouldBindJSON(&theatre); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	} else if theatre.Nama == "" || theatre.Lokasi == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Field nama dan/atau lokasi tidak boleh kosong",
		})
	} else {
		theatre.ID = utils.IDGenerator()
		fmt.Println(theatre)

		db.CreateMovieTheatre(theatre)
		ctx.JSON(http.StatusCreated, gin.H{
			"theatre": theatre,
		})
	}
}
