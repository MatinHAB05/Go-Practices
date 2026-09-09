package utils

import (
	"movieapp/models"
	"regexp"
)

func IsValidActorName(name string) bool {
	r := regexp.MustCompile(`^[a-zA-Z]{1,20}$`)
	return r.MatchString(name)
}

func IsValidMovieTitle(title string) bool {
	return len(title) <= 20
}

func IsValidReleaseYear(year int) bool {
	return 1888 <= year && year <= 2024
}

func IsValidVideoQuality(q string) bool {
	return q == string(models.FOUR_K) || q == string(models.HD) || q == string(models.FHD)
}

func IsValudComparatorOperator(sign string) bool {
	return sign == "=" || sign == ">" || sign == "<" || sign == "<=" || sign == ">="
}
