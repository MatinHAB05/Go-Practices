package service

import (
	"fmt"
	"movieapp/db"
	"movieapp/models"
	utils "movieapp/util"
	"sort"
	"strconv"
	"strings"
)

func ShowMovie(key int) {
	m, i := GetMovieById(key, db.MovieDB)
	if i < 0 {
		fmt.Println("invalid movie id")
		return
	}
	fmt.Println(m)
}

func String2VideoQuality(q string) (models.VideoQuality, error) {
	switch q {
	case string(models.HD):
		return models.HD, nil
	case string(models.FHD):
		return models.FHD, nil
	case string(models.FOUR_K):
		return models.FOUR_K, nil
	default:
		return "", fmt.Errorf("Invalid q : %s", q)
	}
}

func ShowActor(key int) {
	a, i := GetActorById(key, db.ActorDB)
	if i < 0 {
		fmt.Println("invalid cast id")
		return
	}
	fmt.Println(a)
}

func GetMovieById(key int, MovieList []*models.Movie) (*models.Movie, int) {
	for i := range MovieList {
		if MovieList[i].Id == key {
			return MovieList[i], i
		}
	}
	return nil, -1
}
func GetActorById(key int, ActorList []*models.Actor) (*models.Actor, int) {
	for i := range ActorList {
		if ActorList[i].Id == key {
			return ActorList[i], i
		}
	}
	return nil, -1
}

func RemoveMovie(key int) {
	m, i := GetMovieById(key, db.MovieDB)
	if i < 0 {
		fmt.Println("invalid movie id")
		return
	}
	db.MovieDB = append(db.MovieDB[:i], db.MovieDB[i+1:]...)

	for i := range db.ActorDB {
		_, ind := GetMovieById(key, db.ActorDB[i].Movies)
		if ind >= 0 {
			db.ActorDB[i].Movies = append(db.ActorDB[i].Movies[:ind], db.ActorDB[i].Movies[ind+1:]...)
		}
	}

	fmt.Printf("removed successfully %d\n", m.Id)

}

func RemoveActor(key int) {
	a, i := GetActorById(key, db.ActorDB)
	if i < 0 {
		fmt.Println("invalid cast id")
		return
	}
	db.ActorDB = append(db.ActorDB[:i], db.ActorDB[i+1:]...)

	for i := range db.MovieDB {
		_, ind := GetActorById(key, db.MovieDB[i].Actors)
		if ind >= 0 {
			db.MovieDB[i].Actors = append(db.MovieDB[i].Actors[:ind], db.MovieDB[i].Actors[ind+1:]...)
		}
	}

	fmt.Printf("removed successfully %d\n", a.Id)

}

func AddMovie(title, date, quality string) {
	if !utils.IsValidMovieTitle(title) {
		fmt.Println("invalid title")
		return
	}
	date_, _ := strconv.Atoi(date)
	if !utils.IsValidReleaseYear(date_) {
		fmt.Println("invalid date")
		return
	}
	if !utils.IsValidVideoQuality(quality) {
		fmt.Println("invalid quality")
		return
	}
	db.LastMovieID++
	q, _ := String2VideoQuality(quality)
	m := models.Movie{
		Title:       title,
		ReleaseYear: date_,
		Id:          db.LastMovieID,
		Quality:     q,
	}
	db.MovieDB = append(db.MovieDB, &m)
	fmt.Printf("added successfully %d\n", db.LastMovieID)
}

func AddActor(name string) {
	if !utils.IsValidActorName(name) {
		fmt.Println("invalid name")
		return
	}

	db.LastActorID++
	a := models.Actor{
		Name: name,
		Id:   db.LastActorID,
	}
	db.ActorDB = append(db.ActorDB, &a)
	fmt.Printf("added successfully %d\n", db.LastActorID)

}

func LinkActorToMovie(actor_id, movie_id int) {
	a, i := GetActorById(actor_id, db.ActorDB)
	if i < 0 {
		fmt.Println("invalid cast id")
		return
	}

	m, i := GetMovieById(movie_id, db.MovieDB)
	if i < 0 {
		fmt.Println("invalid movie id")
		return
	}

	_, ind := GetActorById(actor_id, m.Actors)
	if ind >= 0 {
		fmt.Println("already linked")
		return
	}

	m.Actors = append(m.Actors, a)
	a.Movies = append(a.Movies, m)
	fmt.Printf("successfully linked %d to %d\n", a.Id, m.Id)
}

func FilterMovies(Movies []*models.Movie, filter func(*models.Movie, ...string) bool, inputs ...string) {
	var ids []int
	for i := range Movies {
		if filter(Movies[i], inputs...) {
			ids = append(ids, Movies[i].Id)
		}
	}
	sort.Ints(ids)
	s := "["
	flag := true
	for i := range ids {
		if flag {
			s += strconv.Itoa(ids[i])
		} else {
			s += ", " + strconv.Itoa(ids[i])

		}
		flag = false
	}
	s += "]"
	fmt.Println(s)
}

func FilterMoviesByTitle(title string) {
	filter := func(m *models.Movie, inputs ...string) bool {
		return strings.HasPrefix(m.Title, title)
	}
	FilterMovies(db.MovieDB, filter)
}

func FilterMoviesByDate(sign string, n int) {
	filter := func(m *models.Movie, inputs ...string) bool {
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
	FilterMovies(db.MovieDB, filter)
}

func FilterMoviesByQuality(q string) {
	filter := func(m *models.Movie, inputs ...string) bool {
		qul, _ := String2VideoQuality(q)
		return m.Quality == qul
	}
	FilterMovies(db.MovieDB, filter)
}
