package filtermovie

import (
	"movieapp/appconstant"
	"strings"
)

func ByQuality(movieQuality, filterQuality string) bool {
	return movieQuality == filterQuality
}

func ByTitle(movieTitle, filterTitle string) bool {
	return strings.HasPrefix(movieTitle, filterTitle)
}

func ByDate(movieYear, filterYear int, action string) bool {
	switch action {
	case appconstant.LessEqualThanSing:
		return movieYear <= filterYear

	case appconstant.LessThanSign:
		return movieYear < filterYear

	case appconstant.GreaterThanSign:
		return movieYear > filterYear

	case appconstant.GreaterEqualThanSign:
		return movieYear >= filterYear

	case appconstant.EqualitySign:
		return movieYear == filterYear

	case appconstant.NotEqualitySign:
		return movieYear != filterYear

	default:
		panic("WTF")

	}
}
