package repository

import "golang-cday-lab/movies/entity"

type MovieRepository interface {
	GetAll() ([]entity.Movie, error)
}
