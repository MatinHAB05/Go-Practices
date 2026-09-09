package db

import (
	"fmt"
	"log"
	"movieapp/models"
	utils "movieapp/util"
	"slices"
	"strconv"
	"strings"
)

func ShowAllMovie() []models.Movie {
	LoadMovies()
	return movieDB
}
func ShowAllActor() []models.Actor {
	LoadActors()
	return actorDB
}

func RemoveMovie(key int) *models.Movie {
	LoadMovies()
	LoadActors()
	defer StoreMovies()
	defer StoreActors()

	m, i := GetMovieById(key, movieDB)
	if i < 0 {
		log.Println("invalid movie id")
		return nil
	}
	movieDB = append(movieDB[:i], movieDB[i+1:]...)

	for i := range actorDB {
		ind := slices.Index(actorDB[i].Movies, key)
		if ind >= 0 {
			actorDB[i].Movies = append(actorDB[i].Movies[:ind], actorDB[i].Movies[ind+1:]...)
		}
	}

	log.Printf("removed successfully %d\n", m.Id)
	return &m

}

func RemoveActor(key int) *models.Actor {
	LoadMovies()
	LoadActors()
	defer StoreMovies()
	defer StoreActors()

	a, i := GetActorById(key, actorDB)
	if i < 0 {
		log.Println("invalid cast id")
		return nil
	}
	actorDB = append(actorDB[:i], actorDB[i+1:]...)

	for i := range movieDB {
		ind := slices.Index(movieDB[i].Actors, key)
		if ind >= 0 {
			movieDB[i].Actors = append(movieDB[i].Actors[:ind], movieDB[i].Actors[ind+1:]...)
		}
	}

	log.Printf("removed successfully %d\n", a.Id)
	return &a
}

func AddMovie(title, date, quality string) *models.Movie {
	LoadMovies()
	defer StoreMovies()

	if !utils.IsValidMovieTitle(title) {
		log.Println("invalid title")
		return nil
	}
	date_, _ := strconv.Atoi(date)
	if !utils.IsValidReleaseYear(date_) {
		log.Println("invalid date")
		return nil
	}
	if !utils.IsValidVideoQuality(quality) {
		log.Println("invalid quality")
		return nil
	}
	lastMovieID++
	q, _ := String2VideoQuality(quality)
	m := models.Movie{
		Title:       title,
		ReleaseYear: date_,
		Id:          lastMovieID,
		Quality:     q,
	}
	movieDB = append(movieDB, m)
	log.Printf("added successfully %d\n", lastMovieID)
	return &m
}

func AddActor(name string) *models.Actor {
	LoadActors()
	defer StoreActors()

	if !utils.IsValidActorName(name) {
		log.Println("invalid name")
		return nil
	}

	lastActorID++
	a := models.Actor{
		Name: name,
		Id:   lastActorID,
	}
	actorDB = append(actorDB, a)
	log.Printf("added successfully %d\n", lastActorID)
	return &a
}

func LinkActorToMovie(actor_id, movie_id int) (*models.Movie, *models.Actor) {
	LoadMovies()
	LoadActors()
	defer StoreMovies()
	defer StoreActors()

	a, i := GetActorById(actor_id, actorDB)
	if i < 0 {
		log.Println("invalid cast id")
		return &models.Movie{}, nil
	}

	m, i := GetMovieById(movie_id, movieDB)
	if i < 0 {
		log.Println("invalid movie id")
		return nil, &models.Actor{}
	}

	ind := slices.Index(m.Actors, actor_id)
	if ind >= 0 {
		log.Println("already linked")
		return nil, nil
	}

	m.Actors = append(m.Actors, actor_id)
	a.Movies = append(a.Movies, movie_id)
	log.Printf("successfully linked %d to %d\n", a.Id, m.Id)
	return &m, &a
}

func ShowMovie(key int) *string {
	m, i := GetMovieById(key, movieDB)
	if i < 0 {
		log.Println("invalid movie id")
		return nil
	}
	out := fmt.Sprintln(m)
	return &out
}

func ShowActor(key int) *string {
	a, i := GetActorById(key, actorDB)
	if i < 0 {
		log.Println("invalid cast id")
		return nil
	}
	out := fmt.Sprintln(a)
	return &out
}

func filter_movies(Movies []models.Movie, filter func(models.Movie) bool) []models.Movie {
	LoadMovies()
	defer StoreMovies()

	var ids []models.Movie = []models.Movie{}
	for i := range Movies {
		if filter(Movies[i]) {
			ids = append(ids, Movies[i])
		}
	}

	slices.SortFunc(ids, func(a, b models.Movie) int {
		return a.Id - b.Id
	})
	return ids
}

func FilterMoviesByTitle(title string, movies []models.Movie) []models.Movie {
	filter := func(m models.Movie) bool {
		return strings.HasPrefix(m.Title, title)
	}
	return filter_movies(movies, filter)
}

func FilterMoviesByDate(sign string, n int, movies []models.Movie) []models.Movie {
	filter := func(m models.Movie) bool {
		switch sign {
		case "<=":
			return m.ReleaseYear <= n

		case "<":
			return m.ReleaseYear < n

		case ">":
			return m.ReleaseYear > n

		case ">=":
			return m.ReleaseYear >= n

		case "=":
			return m.ReleaseYear == n
		default:
			panic("WTF")

		}
	}
	return filter_movies(movies, filter)
}

func FilterMoviesByQuality(q string, movies []models.Movie) []models.Movie {
	filter := func(m models.Movie) bool {
		qul, _ := String2VideoQuality(q)
		return m.Quality == qul
	}
	return filter_movies(movies, filter)
}

func filter_actors(Actors []models.Actor, filter func(models.Actor) bool) []models.Actor {
	LoadActors()
	defer StoreActors()

	var ids []models.Actor = []models.Actor{}
	for i := range Actors {
		if filter(Actors[i]) {
			ids = append(ids, Actors[i])
		}
	}

	slices.SortFunc(ids, func(a, b models.Actor) int {
		return a.Id - b.Id
	})
	return ids
}

func FilterActorsByName(Name string, actors []models.Actor) []models.Actor {
	filter := func(a models.Actor) bool {
		return strings.HasPrefix(a.Name, Name)
	}
	return filter_actors(actors, filter)
}
