package validation

import (
	"github.com/go-playground/validator/v10"
)

func MovieTitleValid(fl validator.FieldLevel) bool {
	return IsValidMovieTitle(fl.Field().String())
}

func ComparatorOperatorValid(fl validator.FieldLevel) bool {
	return IsValudComparatorOperator(fl.Field().String())

}

func MovieQualityValid(fl validator.FieldLevel) bool {
	return IsValidVideoQuality(fl.Field().String())

}

func CastNameValid(fl validator.FieldLevel) bool {
	return IsValidCastName(fl.Field().String())

}

func ReleaseYearValid(fl validator.FieldLevel) bool {
	return IsValidReleaseYear(int(fl.Field().Int()))

}
