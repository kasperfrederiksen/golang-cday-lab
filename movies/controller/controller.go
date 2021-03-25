package controller

import (
	"golang-cday-lab/movies/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type controller struct {
	repository repository.MovieRepository
}

func NewMovieController(repository repository.MovieRepository) controller {
	return controller{repository}
}

func (ctrl *controller) GetAll(c *gin.Context) {
	movies, err := ctrl.repository.GetAll()

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, movies)
}
