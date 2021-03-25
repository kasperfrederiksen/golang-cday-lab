package entity

type Movie struct {
	Title  string   `json:"title"`
	Year   string   `json:"year"`
	Genres []string `json:"genres"`
	Rating float32  `json:"imdbRating"`
}
