package repositoryparam

import "movieapp/entity"

type Movie struct {
	Title       string              `json:"title" db:"title"`
	ReleaseYear int                 `json:"releaseYear" db:"release_year"`
	Quality     entity.VideoQuality `json:"quality" db:"quality"`
	Id          int                 `json:"id" db:"id"`
}
