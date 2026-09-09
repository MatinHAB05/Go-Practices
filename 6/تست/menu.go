package main

import (
	service "movieapp/services"
	"strconv"
	"strings"
)

func Menu() {

	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)

	switch parts[0] {

	case "ADD-MOVIE":
		service.AddMovie(parts[1], parts[2], parts[3])

	case "REM-MOVIE":
		id, _ := strconv.Atoi(parts[1])
		service.RemoveMovie(id)

	case "ADD-CAST":
		service.AddActor(parts[1])

	case "REM-CAST":
		id, _ := strconv.Atoi(parts[1])
		service.RemoveActor(id)

	case "SHOW-MOVIE":
		id, _ := strconv.Atoi(parts[1])
		service.ShowMovie(id)

	case "SHOW-CAST":
		id, _ := strconv.Atoi(parts[1])
		service.ShowActor(id)

	case "LINK-CAST-TO-MOVIE":
		a, _ := strconv.Atoi(parts[1])
		m, _ := strconv.Atoi(parts[2])
		service.LinkActorToMovie(a, m)

	case "FILTER-MOVIES-BY-TITLE":
		service.FilterMoviesByTitle(parts[1])

	case "FILTER-MOVIES-BY-DATE":
		n, _ := strconv.Atoi(parts[2])
		service.FilterMoviesByDate(parts[1], n)

	case "FILTER-MOVIES-BY-QUALITY":
		service.FilterMoviesByQuality(parts[1])
	}
}
