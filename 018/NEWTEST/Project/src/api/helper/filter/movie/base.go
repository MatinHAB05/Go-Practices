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
	case appconstant.IsLessEqualThanSing:
		return movieYear <= filterYear

	case appconstant.IsLessThanSign:
		return movieYear < filterYear

	case appconstant.IsGreaterThanSign:
		return movieYear > filterYear

	case appconstant.IsGreaterEqualThanSign:
		return movieYear >= filterYear

	case appconstant.IsEqualSign:
		return movieYear == filterYear
	default:
		panic("WTF")

	}
}
