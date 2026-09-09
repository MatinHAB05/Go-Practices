package entity

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

func (vidq VideoQuality) ToString() string {
	return string(vidq)
}

type Movie struct {
	Title       string       `json:"title"`
	ReleaseYear int          `json:"releaseYear"`
	Quality     VideoQuality `json:"quality"`
	Id          int          `json:"id"`
	Casts       []int        `json:"casts"`
}

func (m *Movie) String() string {

	var ids []int

	ids = m.Casts

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
