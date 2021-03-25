package main

import (
	"golang-cday-lab/movies/controller"
	"golang-cday-lab/movies/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	repo := repository.NewMovieRepository("db/top-rated-movies-01.json")
	movieCtrl := controller.NewMovieController(repo)

	v1 := r.Group("/v1")
	v1.GET("/movies/", movieCtrl.GetAll)

	r.Run(":8081")

}
