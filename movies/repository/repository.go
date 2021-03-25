package repository

import (
	"encoding/json"
	"golang-cday-lab/movies/entity"
	"io/ioutil"
)

type repository struct {
	Filename string
}

func NewMovieRepository(Filename string) MovieRepository {
	return &repository{Filename}
}

func (repo *repository) GetAll() ([]entity.Movie, error) {
	data, err := ioutil.ReadFile(repo.Filename)
	if err != nil {
		return nil, err
	}

	var movies []entity.Movie
	err = json.Unmarshal(data, &movies)
	if err != nil {
		return nil, err
	}

	return movies, nil
}
