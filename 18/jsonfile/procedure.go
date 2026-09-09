package db

import (
	"log"
	"movieapp/entity"
	"slices"
)

// func ShowAllMovie() []entity.Movie {
// 	LoadMovies()
// 	return movieDB
// }
// func ShowAllCast() []entity.Cast {
// 	LoadCasts()
// 	return castDB
// }

// func RemoveMovie(key int) *entity.Movie {
// 	LoadMovies()
// 	LoadCasts()
// 	defer StoreMovies()
// 	defer StoreCasts()

// 	m, i := GetMovieById(key, movieDB)
// 	if i < 0 {
// 		log.Println("invalid movie id")
// 		return nil
// 	}
// 	movieDB = append(movieDB[:i], movieDB[i+1:]...)

// 	for i := range castDB {
// 		ind := slices.Index(castDB[i].Movies, key)
// 		if ind >= 0 {
// 			castDB[i].Movies = append(castDB[i].Movies[:ind], castDB[i].Movies[ind+1:]...)
// 		}
// 	}

// 	log.Printf("removed successfully %d\n", m.Id)
// 	return &m

// }

// func RemoveCast(key int) *entity.Cast {
// 	LoadMovies()
// 	LoadCasts()
// 	defer StoreMovies()
// 	defer StoreCasts()

// 	a, i := GetCastById(key, castDB)
// 	if i < 0 {
// 		log.Println("invalid cast id")
// 		return nil
// 	}
// 	castDB = append(castDB[:i], castDB[i+1:]...)

// 	for i := range movieDB {
// 		ind := slices.Index(movieDB[i].Casts, key)
// 		if ind >= 0 {
// 			movieDB[i].Casts = append(movieDB[i].Casts[:ind], movieDB[i].Casts[ind+1:]...)
// 		}
// 	}

// 	log.Printf("removed successfully %d\n", a.Id)
// 	return &a
// }

// func AddMovie(title, date, quality string) *entity.Movie {
// 	LoadMovies()
// 	defer StoreMovies()

// 	if !utils.IsValidMovieTitle(title) {
// 		log.Println("invalid title")
// 		return nil
// 	}
// 	date_, _ := strconv.Atoi(date)
// 	if !utils.IsValidReleaseYear(date_) {
// 		log.Println("invalid date")
// 		return nil
// 	}
// 	if !utils.IsValidVideoQuality(quality) {
// 		log.Println("invalid quality")
// 		return nil
// 	}
// 	lastMovieID++
// 	q, _ := String2VideoQuality(quality)
// 	m := entity.Movie{
// 		Title:       title,
// 		ReleaseYear: date_,
// 		Id:          lastMovieID,
// 		Quality:     q,
// 	}
// 	movieDB = append(movieDB, m)
// 	log.Printf("added successfully %d\n", lastMovieID)
// 	return &m
// }

// func AddCast(name string) *entity.Cast {
// 	LoadCasts()
// 	defer StoreCasts()

// 	if !utils.IsValidCastName(name) {
// 		log.Println("invalid name")
// 		return nil
// 	}

// 	lastCastID++
// 	a := entity.Cast{
// 		Name: name,
// 		Id:   lastCastID,
// 	}
// 	castDB = append(castDB, a)
// 	log.Printf("added successfully %d\n", lastCastID)
// 	return &a
// }

// func ShowMovie(key int) *string {
// 	m, i := GetMovieById(key, movieDB)
// 	if i < 0 {
// 		log.Println("invalid movie id")
// 		return nil
// 	}
// 	out := fmt.Sprintln(m)
// 	return &out
// }

// func ShowCast(key int) *string {
// 	a, i := GetCastById(key, castDB)
// 	if i < 0 {
// 		log.Println("invalid cast id")
// 		return nil
// 	}
// 	out := fmt.Sprintln(a)
// 	return &out
// }

// func filter_movies(Movies []entity.Movie, filter func(entity.Movie) bool) []entity.Movie {
// 	LoadMovies()
// 	defer StoreMovies()

// 	var ids []entity.Movie = []entity.Movie{}
// 	for i := range Movies {
// 		if filter(Movies[i]) {
// 			ids = append(ids, Movies[i])
// 		}
// 	}

// 	slices.SortFunc(ids, func(a, b entity.Movie) int {
// 		return a.Id - b.Id
// 	})
// 	return ids
// }

// func FilterMoviesByTitle(title string, movies []entity.Movie) []entity.Movie {
// 	filter := func(m entity.Movie) bool {
// 		return strings.HasPrefix(m.Title, title)
// 	}
// 	return filter_movies(movies, filter)
// }

// func FilterMoviesByDate(sign string, n int, movies []entity.Movie) []entity.Movie {
// 	filter := func(m entity.Movie) bool {
// 		switch sign {
// 		case appconstant.IsLessEqualThanSing:
// 			return m.ReleaseYear <= n

// 		case appconstant.IsLessThanSign:
// 			return m.ReleaseYear < n

// 		case appconstant.IsGreaterThanSign:
// 			return m.ReleaseYear > n

// 		case appconstant.IsGreaterEqualThanSign:
// 			return m.ReleaseYear >= n

// 		case appconstant.IsEqualSign:
// 			return m.ReleaseYear == n
// 		default:
// 			panic("WTF")

// 		}
// 	}
// 	return filter_movies(movies, filter)
// }

// func FilterMoviesByQuality(q string, movies []entity.Movie) []entity.Movie {
// 	filter := func(m entity.Movie) bool {
// 		qul, _ := String2VideoQuality(q)
// 		return m.Quality == qul
// 	}
// 	return filter_movies(movies, filter)
// }

// func filter_casts(Casts []entity.Cast, filter func(entity.Cast) bool) []entity.Cast {
// 	LoadCasts()
// 	defer StoreCasts()

// 	var ids []entity.Cast = []entity.Cast{}
// 	for i := range Casts {
// 		if filter(Casts[i]) {
// 			ids = append(ids, Casts[i])
// 		}
// 	}

// 	slices.SortFunc(ids, func(a, b entity.Cast) int {
// 		return a.Id - b.Id
// 	})
// 	return ids
// }

// func FilterCastsByName(Name string, casts []entity.Cast) []entity.Cast {
// 	filter := func(a entity.Cast) bool {
// 	}
// 	return filter_casts(casts, filter)
// }

func LinkCastToMovie(cast_id, movie_id int) (*entity.Movie, *entity.Cast) {
	LoadMovies()
	LoadCasts()
	defer StoreMovies()
	defer StoreCasts()

	a, i := GetCastById(cast_id, castDB)
	if i < 0 {
		log.Println("invalid cast id")
		return &entity.Movie{}, nil
	}

	m, i := GetMovieById(movie_id, movieDB)
	if i < 0 {
		log.Println("invalid movie id")
		return nil, &entity.Cast{}
	}

	ind := slices.Index(m.Casts, cast_id)
	if ind >= 0 {
		log.Println("already linked")
		return nil, nil
	}

	m.Casts = append(m.Casts, cast_id)
	a.Movies = append(a.Movies, movie_id)
	log.Printf("successfully linked %d to %d\n", a.Id, m.Id)
	return &m, &a
}
