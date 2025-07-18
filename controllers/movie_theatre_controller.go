package controllers

import (
	"database/sql"
	"fmt"
	"movietheatres-go/db"
	"movietheatres-go/models"
	"movietheatres-go/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PostMovieTheatre(ctx *gin.Context) {
	var theatre models.MovieTheatre

	if err := ctx.ShouldBindJSON(&theatre); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
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

func GetMovieTheatres(ctx *gin.Context) {
	theatres, err := db.GetMovieTheatres()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"theatres": theatres,
		})
	}
}

func GetMovieTheatreById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
	} else {
		theatre, err := db.GetMovieTheatreById(id)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, gin.H{
					"error": "Theatre doesn't exist",
				})
			} else {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
			}
			return
		} else {
			ctx.JSON(http.StatusOK, theatre)
		}
	}
}

func UpdateMovieTheatre(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	} else {
		var input models.MovieTheatre

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":   "Invalid input",
				"details": err.Error(),
			})
			return
		} else {
			rowsUpdated, err := db.UpdateMovieTheatre(id, input)

			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error":   "Failed to update theatre",
					"details": err.Error(),
				})
				return
			} else if rowsUpdated == 0 {
				ctx.JSON(http.StatusNotFound, gin.H{
					"message": "No theatre found with the given ID",
				})
				return
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "Theatre was successfully updated",
					"rows":    rowsUpdated,
				})
			}
		}
	}
}

func DeleteMovieTheatre(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	} else {
		rowsDeleted, err := db.DeleteMovieTheatre(id)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":  "Failed to delete theatre",
				"detais": err.Error(),
			})
			return
		} else if rowsDeleted == 0 {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "No theatre found with the given ID",
			})
			return
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Theatre has been deleted",
				"rows":    rowsDeleted,
			})
		}
	}
}
