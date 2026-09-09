package entity

import (
	"fmt"
	"sort"
)

type Cast struct {
	Name   string `json:"name"`
	Id     int    `json:"id"`
	Movies []int  `json:"movies"`
}

func (a *Cast) String() string {

	var ids []int

	ids = a.Movies
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
