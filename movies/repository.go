package movies

import (
	"encoding/json"
	"io/ioutil"
)

type MovieRepository interface {
	GetAll() ([]Movie, error)
}

type repository struct {
	Filename string
}

func NewMovieRepository(Filename string) MovieRepository {
	return &repository{Filename}
}

func (repo *repository) GetAll() ([]Movie, error) {
	data, err := ioutil.ReadFile(repo.Filename)
	if err != nil {
		return nil, err
	}

	var movies []Movie
	err = json.Unmarshal(data, &movies)
	if err != nil {
		return nil, err
	}

	return movies, nil
}
