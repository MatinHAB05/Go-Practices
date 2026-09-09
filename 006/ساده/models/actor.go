package models

import (
	"fmt"
	"sort"
)

type Actor struct {
	Name   string
	Id     int
	Movies []*Movie
}

func (a *Actor) String() string {

	var ids []int

	for _, m := range a.Movies {
		ids = append(ids, m.Id)
	}

	sort.Ints(ids)

	movies := "["
	for i, id := range ids {
		if i > 0 {
			movies += ", "
		}
		movies += fmt.Sprintf("%d", id)
	}

	movies += "]"

	return fmt.Sprintf(`{name:"%s", movies:%s}`, a.Name, movies)
}
