package sorted

import (
	"movieapp/entity"
	"slices"
)

func SortMovieCastById(movies []entity.MovieCast) {
	slices.SortFunc(movies, func(a, b entity.MovieCast) int {
		return a.Id - b.Id
	})
}
