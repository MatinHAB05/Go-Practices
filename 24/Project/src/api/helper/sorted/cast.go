package sorted

import (
	"movieapp/entity"
	"slices"
)

func SortCastsById(casts []entity.Cast) {
	slices.SortFunc(casts, func(a, b entity.Cast) int {
		return a.Id - b.Id
	})
}
