package db

import (
	"encoding/json"
	"fmt"
	"log"
	"movieapp/models"
	"os"
)

var movieDB []models.Movie
var actorDB []models.Actor

var lastMovieID = -1
var lastActorID = -1

type actorJson struct {
	Actors      []models.Actor
	LastActorID int
}

type movieJson struct {
	Movies      []models.Movie
	LastMovieID int
}

func LoadActors() error {
	js, err := os.Open("db/actors.json")
	if err != nil {
		log.Println(err)
		return err
	}
	defer js.Close()
	var temp actorJson

	json.NewDecoder(js).Decode(&temp)
	lastActorID = temp.LastActorID
	actorDB = temp.Actors
	return nil
}

func StoreActors() error {
	js, err := os.Create("db/actors.json")
	if err != nil {
		log.Println(err)
		return err
	}
	defer js.Close()
	var temp actorJson = actorJson{
		Actors:      actorDB,
		LastActorID: lastActorID,
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

func GetMovieById(key int, MovieList []models.Movie) (models.Movie, int) {
	for i := range MovieList {
		if MovieList[i].Id == key {
			return MovieList[i], i
		}
	}
	return models.Movie{}, -1
}
func GetActorById(key int, ActorList []models.Actor) (models.Actor, int) {
	for i := range ActorList {
		if ActorList[i].Id == key {
			return ActorList[i], i
		}
	}
	return models.Actor{}, -1
}
