package models

import (
	"fmt"
	"sort"
)

type VideoQuality string

const (
	HD     VideoQuality = "720p"
	FHD    VideoQuality = "1080p"
	FOUR_K VideoQuality = "4K"
)

type Movie struct {
	Title       string
	ReleaseYear int
	Quality     VideoQuality
	Id          int
	Actors      []*Actor
}

func (m *Movie) String() string {

	var ids []int

	for _, a := range m.Actors {
		ids = append(ids, a.Id)
	}

	sort.Ints(ids)

	cast := "["
	for i, id := range ids {
		if i > 0 {
			cast += ", "
		}
		cast += fmt.Sprintf("%d", id)
	}

	cast += "]"

	return fmt.Sprintf(
		`{title:"%s", date:"%d", quality:"%s", casts:%s}`,
		m.Title,
		m.ReleaseYear,
		m.Quality,
		cast,
	)
}
