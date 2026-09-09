package validation

import (
	"movieapp/appconstant"
	"movieapp/entity"
	"regexp"
	"slices"
)

func IsValidCastName(name string) bool {
	r := regexp.MustCompile(appconstant.CastNameRegex)
	return r.MatchString(name)
}

func IsValidMovieTitle(title string) bool {
	return len(title) <= appconstant.MaxMovieTitleLength
}

func IsValidReleaseYear(year int) bool {
	return appconstant.MinReleaseMovieYear <= year && year <= appconstant.MaxReleaseMovieYear
}

func IsValidVideoQuality(q string) bool {
	return q == string(entity.FOUR_K) || q == string(entity.HD) || q == string(entity.FHD)
}

func IsValudComparatorOperator(sign string) bool {
	return slices.Contains(appconstant.ComparatorOperatorSigns[:], sign)
}
