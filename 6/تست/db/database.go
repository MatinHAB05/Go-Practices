package db

import "movieapp/models"

var MovieDB []*models.Movie
var ActorDB []*models.Actor

var LastMovieID = -1
var LastActorID = -1
