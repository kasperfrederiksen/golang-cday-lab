package movies

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type controller struct {
	repository MovieRepository
}

func NewMovieController(repository MovieRepository) controller {
	return controller{repository}
}

func (ctrl *controller) GetAll(c *gin.Context) {
	movies, err := ctrl.repository.GetAll()

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, movies)
}
