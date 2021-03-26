package main

import (
	"golang-cday-lab/movies"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	repo := movies.NewMovieRepository("db/top-rated-movies-01.json")
	movieCtrl := movies.NewMovieController(repo)

	v1 := r.Group("/v1")
	v1.GET("/movies/", movieCtrl.GetAll)

	r.Run(":8081")

}
