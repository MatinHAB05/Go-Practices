package sorted

import (
	"movieapp/entity"
	"slices"
)

func SortMoviesById(movies []entity.Movie) {
	slices.SortFunc(movies, func(a, b entity.Movie) int {
		return a.Id - b.Id
	})
}
