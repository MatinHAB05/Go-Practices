package db

import (
	"encoding/json"
	"fmt"
	"log"
	"movieapp/entity"
	"os"
)

var movieDB []entity.Movie
var castDB []entity.Cast

var lastMovieID = -1
var lastCastID = -1

type castJson struct {
	Casts      []entity.Cast
	LastCastID int
}

type movieJson struct {
	Movies      []entity.Movie
	LastMovieID int
}

func LoadCasts() error {
	js, err := os.Open("db/casts.json")
	if err != nil {
		log.Println(err)
		return err
	}
	defer js.Close()
	var temp castJson

	json.NewDecoder(js).Decode(&temp)
	lastCastID = temp.LastCastID
	castDB = temp.Casts
	return nil
}

func StoreCasts() error {
	js, err := os.Create("db/casts.json")
	if err != nil {
		log.Println(err)
		return err
	}
	defer js.Close()
	var temp castJson = castJson{
		Casts:      castDB,
		LastCastID: lastCastID,
	}

	enc := json.NewEncoder(js)
	enc.SetIndent("", "  ")
	enc.Encode(temp)

	return nil

}

func LoadMovies() error {
	js, err := os.Open("db/movies.json")
	if err != nil {
		log.Println(err)
		return err
	}
	defer js.Close()
	var temp movieJson

	json.NewDecoder(js).Decode(&temp)
	lastMovieID = temp.LastMovieID
	movieDB = temp.Movies

	return nil
}

func StoreMovies() error {
	js, err := os.Create("db/movies.json")
	if err != nil {
		log.Println(err)
		return err
	}
	defer js.Close()
	var temp movieJson = movieJson{
		Movies:      movieDB,
		LastMovieID: lastMovieID,
	}

	enc := json.NewEncoder(js)
	enc.SetIndent("", "  ")
	enc.Encode(temp)

	return nil

}

func String2VideoQuality(q string) (entity.VideoQuality, error) {
	switch q {
	case string(entity.HD):
		return entity.HD, nil
	case string(entity.FHD):
		return entity.FHD, nil
	case string(entity.FOUR_K):
		return entity.FOUR_K, nil
	default:
		return "", fmt.Errorf("Invalid q : %s", q)
	}
}

func GetMovieById(key int, MovieList []entity.Movie) (entity.Movie, int) {
	for i := range MovieList {
		if MovieList[i].Id == key {
			return MovieList[i], i
		}
	}
	return entity.Movie{}, -1
}
func GetCastById(key int, CastList []entity.Cast) (entity.Cast, int) {
	for i := range CastList {
		if CastList[i].Id == key {
			return CastList[i], i
		}
	}
	return entity.Cast{}, -1
}
