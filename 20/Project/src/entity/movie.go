package entity

import (
	"fmt"
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
}

func (m *Movie) String() string {

	return fmt.Sprintf(
		`{id : "%d" , title:"%s" , date:"%d" , quality:"%s"}`,
		m.Id,
		m.Title,
		m.ReleaseYear,
		m.Quality,
	)
}
